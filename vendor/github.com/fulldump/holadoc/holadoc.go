package holadoc

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/url"
	"os"
	"path"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/chroma"
	html2 "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	goldmarkHtml "github.com/yuin/goldmark/renderer/html"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Config struct {
	Src       string `json:"src"`
	Www       string `json:"www"`
	Versions  string `json:"versions" usage:"default version is the first one"`
	Languages string `json:"languages" usage:"default language is the first one"`
	Serve     string `json:"serve" usage:"Address to serve files locally, example ':8080'"`
	BaseURL   string `json:"baseurl" usage:"Base URL for sitemap generation, e.g. https://holacloud.com"`
	Version   bool   `json:"version" usage:"Display version and exit"`
}

// todo: avoid globals
var versions []string
var languages []string

func HolaDoc(c Config) {

	// clear output
	_ = os.RemoveAll(c.Www)
	err := os.MkdirAll(c.Www, 0777)
	if err != nil {
		panic(err.Error())
	}

	versions = strings.Split(c.Versions, ",")
	languages = strings.Split(c.Languages, ",")

	root := &Node{}

	readNodes(root, c.Src, c.Www)
	root.PrettyPrint(0)

	urls := []string{}

	traverseNodes(root, func(node *Node) {

		for _, version := range versions {
			for _, language := range languages {

				variation := getBestVariation(node.Variations, language, version)
				if variation == nil {
					fmt.Println("skip:", node.Path)
					continue
				}

				langMenu := ""
				{
					langMenu += `<div class="languages">`
					for _, l := range languages {
						class := ""
						if l == language {
							class += "selected"
						}
						langMenu += `<a class="` + class + `" href="` + getLink(node, l, version) + `">` + l + `</a>`
					}
					langMenu += `</div>`
				}

				versionMenu := ""
				{
					// Collect versions that have distinct content for this node tree
					availableVersions := []string{}
					if hasVersions(node) {
						for _, v := range versions {
							for _, variation := range node.Variations {
								if variation.Version == v {
									availableVersions = append(availableVersions, v)
									break
								}
							}
						}
					}
					if len(availableVersions) > 1 {
						versionMenu += `<div class="versions">`
						for _, v := range availableVersions {
							class := ""
							if v == version {
								class += "selected"
							}
							versionMenu += `<a class="` + class + `" href="` + getLink(node, language, v) + `">` + v + `</a>`
						}
						versionMenu += `</div>`
					}
				}

				onThisPage := ""

				content := ""

				{ // content

					var htmlReader io.Reader

					switch strings.ToLower(path.Ext(variation.Filename)) {
					case ".html":
						src, err := os.Open(variation.Filename)
						if err != nil {
							panic(err.Error())
						}
						htmlReader = src
					case ".md":
						src, err := os.ReadFile(variation.Filename)
						if err != nil {
							panic(err.Error())
						}
						htmlReader = md2html(src)
						// htmlReader = strings.NewReader("Hello, this is the new md!")
					}

					doc := &html.Node{
						Type:     html.ElementNode,
						Data:     "body",
						DataAtom: atom.Body,
					}
					nodes, err := html.ParseFragment(htmlReader, doc)
					if err != nil {
						panic(err.Error())
					}

					{ // index

						for _, n := range nodes {
							traverseHtml(n, func(node *html.Node) {
								tag := strings.ToLower(node.Data)
								if in([]string{"h2", "h3", "h4", "h5", "h6"}, tag) && node.FirstChild != nil {
									title := node.FirstChild.Data
									node.Attr = append(node.Attr, html.Attribute{
										Key: "id",
										Val: url.PathEscape(title), // todo: slug?
									})
									onThisPage += `<div class="index-` + node.Data + `">` + "\n"
									onThisPage += `<a href="#` + url.PathEscape(title) + `">` + title + `</a>` + "\n"
									onThisPage += `</div>` + "\n"
								}
								if tag == "a" {
									href := getAttribute(node, "href")
									if href != "" {
										target := getNode(root, href)
										if target != nil {
											setAttribute(node, "href", getLink(target, variation.Language, variation.Version))
											if node.FirstChild != nil && node.FirstChild.FirstChild == nil && node.FirstChild.Type == html.TextNode {
												node.FirstChild.Data = target.Name
											} else if node.FirstChild == nil {
												node.AppendChild(&html.Node{
													Type:     html.TextNode,
													DataAtom: 0,
													Data:     target.Name,
												})
											}
										}
									}
								}
								if tag == "code" {
									if node.Parent == nil || node.Parent.Data != "pre" {
										return // skip inline code
									}

									code := node.FirstChild.Data
									code = strings.TrimPrefix(code, "\n")

									if request, ok := parseRequestBlock(code); ok {
										component := renderHttpRequestComponent(request, code)
										parts, err := html.ParseFragment(strings.NewReader(component), doc)
										if err != nil {
											panic(err)
										}
										if len(parts) > 0 {
											replaceNode(node.Parent, parts[0])
										}
										return
									}

									lexer := lexers.Get(getAttribute(node, "lang"))
									if lexer == nil {
										lexer = lexers.Get(strings.TrimPrefix(getAttribute(node, "class"), "language-"))
									}
									if lexer == nil {
										lexer = lexers.Analyse(code)
									}
									if lexer == nil {
										lexer = lexers.Fallback
									}
									lexer = chroma.Coalesce(lexer)

									style := styles.Get("solarized-dark") // monokai github-dark
									if style == nil {
										style = styles.Fallback
									}
									formatter := html2.New(html2.WithLineNumbers(true), html2.LinkableLineNumbers(true, "L"))

									iterator, err := lexer.Tokenise(nil, code)

									codeOutput := &bytes.Buffer{}
									err = formatter.Format(codeOutput, style, iterator)
									if err != nil {
										panic(err.Error())
									}

									fragmentDoc := &html.Node{
										Type:     html.ElementNode,
										Data:     "body",
										DataAtom: atom.Body,
									}

									parts, err := html.ParseFragment(codeOutput, fragmentDoc)
									if err != nil {
										panic(err)
									}
									if len(parts) > 0 {
										replaceNode(node.Parent, parts[0])
									}

								}
							})
						}
					}

					{ // print content
						b := &bytes.Buffer{}

						// if variation.Version != "" && version > variation.Version { // todo: make this comparison better (taking into account numbers, not only strings)
						// 	fmt.Fprintln(b, `<div class="alert">This has been unchanged since version `+variation.Version+`</div>`)
						// }

						for _, n := range nodes {
							html.Render(b, n)
						}
						content = b.String()
					}

				}

				outputPath := getOutputPath(node, variation, language, version)

				data := map[string]any{
					"lang":        variation.Language,
					"langs":       languages,
					"langMenu":    template.HTML(langMenu),
					"title":       variation.Title,
					"url":         variation.Url,
					"canonicalUrl": absoluteURL(c.BaseURL, outputPath),
					"hreflangLinks": buildHreflangLinks(c.BaseURL, node, version),
					"filename":    variation.Filename,
					"version":     variation.Version,
					"versions":    versions,
					"versionMenu": template.HTML(versionMenu),
					"tree":        template.HTML(getIndex(root, node, language, version)),
					"breadcrumb":  template.HTML(getBreadcrumb(node, language, version)),
					"index":       template.HTML(onThisPage),
					"content":     template.HTML(content),
				}

				urls = append(urls, outputPath)

				newFilename := path.Join(c.Www, outputPath)
				os.MkdirAll(path.Dir(newFilename), 0777) // todo: handle err

				f, err := os.Create(newFilename)
				if err != nil {
					panic(err.Error())
				}

				temp := getTemplate(node, map[string]any{
					"link": func(p string) template.HTML {

						target := getNode(root, p)
						if target == nil {
							panic("link for '" + p + "' does not exist")
						}

						class := "link"
						if target == node {
							class += " selected"
						}

						variation := getBestVariation(target.Variations, language, version)

						return template.HTML(`<a class="` + class + `" href="` + getLink(target, language, version) + `">` + variation.Title + `</a>`)
					},

					"tree": func(p string) template.HTML {

						target := getNode(root, p)
						if target == nil {
							panic("link for '" + p + "' does not exist")
						}

						return template.HTML(getIndex(target, node, language, version))
					},

					"isUnder": func(p string) bool {

						target := getNode(root, p)
						if target == nil {
							panic("link for '" + p + "' does not exist")
						}

						n := node
						for n != nil {
							if n == target {
								return true
							}
							n = n.Parent
						}

						return false
					},
				})
				temp.Execute(f, data)

				err = f.Close()
				if err != nil {
					panic(err.Error())
				}

			}
		}

	})

	if c.BaseURL != "" {
		generateSitemap(c.Www, c.BaseURL, urls)
	}

}

func generateSitemap(www, baseURL string, paths []string) {
	if len(paths) == 0 {
		return
	}

	now := time.Now().Format("2006-01-02")

	var b bytes.Buffer
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n")

	seen := map[string]bool{}
	for _, p := range paths {
		// clean path: strip index.html, ensure leading /
		u := "/" + strings.TrimSuffix(p, "index.html")
		if seen[u] {
			continue
		}
		seen[u] = true

		loc := strings.TrimRight(baseURL, "/") + u
		b.WriteString("  <url>\n")
		b.WriteString(fmt.Sprintf("    <loc>%s</loc>\n", loc))
		b.WriteString(fmt.Sprintf("    <lastmod>%s</lastmod>\n", now))
		b.WriteString("  </url>\n")
	}

	b.WriteString("</urlset>\n")

	err := os.WriteFile(path.Join(www, "sitemap.xml"), b.Bytes(), 0644)
	if err != nil {
		fmt.Println("WARNING: failed to write sitemap.xml:", err)
	}
}

type hreflangLink struct {
	Lang string
	URL  string
}

func buildHreflangLinks(baseURL string, node *Node, version string) []hreflangLink {
	if baseURL == "" {
		return nil
	}

	links := []hreflangLink{}

	for _, l := range languages {
		variation := getBestVariation(node.Variations, l, version)
		if variation == nil {
			continue
		}

		links = append(links, hreflangLink{
			Lang: l,
			URL:  absoluteURL(baseURL, getOutputPath(node, variation, l, version)),
		})
	}

	if len(languages) > 0 {
		variation := getBestVariation(node.Variations, languages[0], version)
		if variation != nil {
			links = append(links, hreflangLink{
				Lang: "x-default",
				URL:  absoluteURL(baseURL, getOutputPath(node, variation, languages[0], version)),
			})
		}
	}

	return links
}

func absoluteURL(baseURL, outputPath string) string {
	if baseURL == "" {
		return ""
	}

	u := "/" + strings.TrimSuffix(outputPath, "index.html")
	return strings.TrimRight(baseURL, "/") + u
}

func getTemplate(node *Node, funcs template.FuncMap) *template.Template {

	for node != nil {
		if node.Template == "" {
			node = node.Parent
			continue
		}

		gohtml, err := os.ReadFile(node.Template)
		if err != nil {
			panic(err.Error())
		}

		temp, err := template.New("").Funcs(funcs).Parse(string(gohtml))
		if err != nil {
			panic(err.Error())
		}

		return temp
	}

	panic("No template found!!!")

	return nil
}

func getNode(root *Node, path string) *Node {
	if path == "" {
		return root
	}

	n := root
out:
	for _, p := range strings.Split(path, "/") {
		for _, child := range n.Children {
			if child.Name == p {
				n = child
				continue out
			}
		}
		return nil
	}

	return n
}

func getAttribute(node *html.Node, key string) string {
	for _, a := range node.Attr {
		if strings.EqualFold(a.Key, key) {
			return a.Val
		}
	}
	return ""
}

func setAttribute(node *html.Node, key, value string) {
	for i, a := range node.Attr {
		if strings.EqualFold(a.Key, key) {
			node.Attr[i].Val = value
			return
		}
	}
	node.Attr = append(node.Attr, html.Attribute{
		Key: key,
		Val: value,
	})
}

func replaceNode(oldNode, newNode *html.Node) {
	parent := oldNode.Parent
	if parent == nil {
		*oldNode = *newNode
		return
	}
	parent.InsertBefore(newNode, oldNode)
	parent.RemoveChild(oldNode)
}

type httpRequest struct {
	Method  string
	URL     string
	Headers []httpHeader
	Body    string
}

type httpHeader struct {
	Name  string
	Value string
}

func parseRequestBlock(code string) (httpRequest, bool) {
	if req, ok := parseCurlRequest(code); ok {
		return req, true
	}
	return parseRawHTTPRequest(code)
}

func parseCurlRequest(code string) (httpRequest, bool) {
	tokens := shellFields(code)
	if len(tokens) == 0 || tokens[0] != "curl" {
		return httpRequest{}, false
	}
	for _, token := range tokens[1:] {
		if token == "curl" {
			return httpRequest{}, false
		}
	}

	req := httpRequest{Method: "GET"}
	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		switch token {
		case "-X", "--request":
			if i+1 < len(tokens) {
				req.Method = strings.ToUpper(tokens[i+1])
				i++
			}
		case "-H", "--header":
			if i+1 < len(tokens) {
				req.Headers = append(req.Headers, splitHeader(tokens[i+1]))
				i++
			}
		case "-d", "--data", "--data-raw", "--data-binary", "--data-ascii":
			if i+1 < len(tokens) {
				if req.Body != "" {
					req.Body += "&"
				}
				req.Body += tokens[i+1]
				i++
			}
		case "--url":
			if i+1 < len(tokens) {
				req.URL = tokens[i+1]
				i++
			}
		default:
			if strings.HasPrefix(token, "http://") || strings.HasPrefix(token, "https://") {
				req.URL = token
			}
		}
	}

	if req.URL == "" {
		return httpRequest{}, false
	}
	if req.Method == "GET" && req.Body != "" {
		req.Method = "POST"
	}

	return req, true
}

func parseRawHTTPRequest(code string) (httpRequest, bool) {
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(code), "\r\n", "\n"), "\n")
	if len(lines) == 0 {
		return httpRequest{}, false
	}

	first := strings.Fields(lines[0])
	if len(first) < 2 || !isHTTPMethod(first[0]) {
		return httpRequest{}, false
	}

	req := httpRequest{Method: strings.ToUpper(first[0])}
	pathOrURL := first[1]
	bodyStart := 0
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if strings.TrimSpace(line) == "" {
			bodyStart = i + 1
			break
		}
		name, value, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		header := httpHeader{Name: strings.TrimSpace(name), Value: strings.TrimSpace(value)}
		if strings.EqualFold(header.Name, "Host") {
			req.URL = "https://" + header.Value + pathOrURL
			continue
		}
		req.Headers = append(req.Headers, header)
	}

	if req.URL == "" {
		if strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://") {
			req.URL = pathOrURL
		} else {
			req.URL = "https://api.hola.cloud" + pathOrURL
		}
	}
	if bodyStart > 0 && bodyStart < len(lines) {
		req.Body = strings.TrimSpace(strings.Join(lines[bodyStart:], "\n"))
	}

	return req, true
}

func isHTTPMethod(method string) bool {
	switch strings.ToUpper(method) {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS":
		return true
	default:
		return false
	}
}

func splitHeader(header string) httpHeader {
	name, value, found := strings.Cut(header, ":")
	if !found {
		return httpHeader{Name: strings.TrimSpace(header)}
	}
	return httpHeader{Name: strings.TrimSpace(name), Value: strings.TrimSpace(value)}
}

func shellFields(s string) []string {
	fields := []string{}
	var b strings.Builder
	quote := rune(0)
	inField := false
	runes := []rune(s)

	flush := func() {
		if inField {
			fields = append(fields, b.String())
			b.Reset()
			inField = false
		}
	}

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if quote == 0 {
			if r == '\\' {
				if i+1 < len(runes) && runes[i+1] == '\n' {
					i++
					for i+1 < len(runes) && (runes[i+1] == ' ' || runes[i+1] == '\t') {
						i++
					}
					continue
				}
				if i+1 < len(runes) {
					i++
					b.WriteRune(runes[i])
					inField = true
				}
				continue
			}
			if r == '\'' || r == '"' {
				quote = r
				inField = true
				continue
			}
			if r == ' ' || r == '\t' || r == '\r' || r == '\n' {
				flush()
				continue
			}
			b.WriteRune(r)
			inField = true
			continue
		}

		if r == quote {
			quote = 0
			continue
		}
		if quote == '"' && r == '\\' && i+1 < len(runes) {
			i++
			b.WriteRune(runes[i])
			continue
		}
		b.WriteRune(r)
	}
	flush()

	return fields
}

func renderHttpRequestComponent(req httpRequest, source string) string {
	curl := strings.TrimSpace(source)
	if !strings.HasPrefix(curl, "curl") {
		curl = renderCurlRequest(req)
	}
	snippets := []struct {
		Lang  string
		Label string
		Code  string
	}{
		{Lang: "curl", Label: "curl", Code: curl},
		{Lang: "http", Label: "HTTP", Code: renderRawHTTPRequest(req)},
		{Lang: "go", Label: "Go", Code: renderGoRequest(req)},
		{Lang: "php", Label: "PHP", Code: renderPHPRequest(req)},
		{Lang: "python", Label: "Python", Code: renderPythonRequest(req)},
		{Lang: "nodejs", Label: "Node.js", Code: renderNodeRequest(req)},
		{Lang: "javascript", Label: "JavaScript", Code: renderJavaScriptRequest(req)},
		{Lang: "java", Label: "Java", Code: renderJavaRequest(req)},
	}

	var b strings.Builder
	b.WriteString(`<div class="http-request" data-http-request>`)
	b.WriteString(`<div class="http-request-header"><div class="http-request-tabs" role="tablist" aria-label="HTTP request examples">`)
	for i, snippet := range snippets {
		selected := "false"
		class := "http-request-tab"
		if i == 0 {
			selected = "true"
			class += " active"
		}
		b.WriteString(`<button type="button" class="` + class + `" data-http-lang="` + snippet.Lang + `" role="tab" aria-selected="` + selected + `">` + snippet.Label + `</button>`)
	}
	b.WriteString(`</div><button type="button" class="http-request-copy" data-http-copy title="Copy request" aria-label="Copy request"><svg aria-hidden="true" viewBox="0 0 24 24"><path d="M8 7a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2V7Zm2 0v11h8V7h-8ZM4 3a2 2 0 0 1 2-2h9v2H6v11H4V3Z" fill="currentColor"/></svg><span>Copy</span></button></div>`)
	for i, snippet := range snippets {
		class := "http-request-panel"
		if i == 0 {
			class += " active"
		}
		b.WriteString(`<div class="` + class + `" data-http-panel="` + snippet.Lang + `" role="tabpanel"><pre><code>`)
		b.WriteString(template.HTMLEscapeString(snippet.Code))
		b.WriteString(`</code></pre></div>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

func renderCurlRequest(req httpRequest) string {
	var b strings.Builder
	b.WriteString("curl -X " + req.Method + " " + strconv.Quote(req.URL))
	for _, header := range req.Headers {
		b.WriteString(" \\\n  -H " + strconv.Quote(header.Name+": "+header.Value))
	}
	if req.Body != "" {
		b.WriteString(" \\\n  -d " + shellQuote(req.Body))
	}
	return b.String()
}

func renderRawHTTPRequest(req httpRequest) string {
	u, err := url.Parse(req.URL)
	path := req.URL
	host := ""
	if err == nil {
		path = u.RequestURI()
		host = u.Host
	}
	var b strings.Builder
	b.WriteString(req.Method + " " + path + " HTTP/1.1\n")
	if host != "" {
		b.WriteString("Host: " + host + "\n")
	}
	for _, header := range req.Headers {
		b.WriteString(header.Name + ": " + header.Value + "\n")
	}
	if req.Body != "" {
		b.WriteString("\n" + req.Body)
	}
	return b.String()
}

func renderGoRequest(req httpRequest) string {
	body := "nil"
	jsonBody, isJSON := parseJSONBody(req.Body)
	if req.Body != "" {
		body = "strings.NewReader(body)"
	}

	var b strings.Builder
	b.WriteString("package main\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"io\"\n")
	b.WriteString("\t\"net/http\"\n")
	if isJSON {
		b.WriteString("\t\"encoding/json\"\n")
	}
	if req.Body != "" {
		b.WriteString("\t\"strings\"\n")
	}
	b.WriteString(")\n\n")
	b.WriteString("func main() {\n")
	if isJSON {
		b.WriteString("\tpayload := " + goValue(jsonBody) + "\n")
		b.WriteString("\tbodyBytes, err := json.Marshal(payload)\n")
		b.WriteString("\tif err != nil {\n\t\tpanic(err)\n\t}\n")
		b.WriteString("\tbody := string(bodyBytes)\n\n")
	} else if req.Body != "" {
		b.WriteString("\tbody := " + strconv.Quote(req.Body) + "\n\n")
	}
	b.WriteString("\treq, err := http.NewRequest(" + strconv.Quote(req.Method) + ", " + strconv.Quote(req.URL) + ", " + body + ")\n")
	b.WriteString("\tif err != nil {\n\t\tpanic(err)\n\t}\n")
	for _, header := range req.Headers {
		b.WriteString("\treq.Header.Set(" + strconv.Quote(header.Name) + ", " + strconv.Quote(header.Value) + ")\n")
	}
	b.WriteString("\n\tresp, err := http.DefaultClient.Do(req)\n")
	b.WriteString("\tif err != nil {\n\t\tpanic(err)\n\t}\n")
	b.WriteString("\tdefer resp.Body.Close()\n\n")
	b.WriteString("\tresponseBody, err := io.ReadAll(resp.Body)\n")
	b.WriteString("\tif err != nil {\n\t\tpanic(err)\n\t}\n")
	b.WriteString("\tfmt.Println(string(responseBody))\n")
	b.WriteString("}\n")
	return b.String()
}

func renderPHPRequest(req httpRequest) string {
	jsonBody, isJSON := parseJSONBody(req.Body)
	var b strings.Builder
	b.WriteString("<?php\n")
	if isJSON {
		b.WriteString("$payload = " + phpValue(jsonBody) + ";\n")
		b.WriteString("$body = json_encode($payload);\n\n")
	} else if req.Body != "" {
		b.WriteString("$body = " + phpQuote(req.Body) + ";\n\n")
	}
	b.WriteString("$ch = curl_init();\n\n")
	b.WriteString("curl_setopt_array($ch, [\n")
	b.WriteString("    CURLOPT_URL => " + phpQuote(req.URL) + ",\n")
	b.WriteString("    CURLOPT_RETURNTRANSFER => true,\n")
	b.WriteString("    CURLOPT_CUSTOMREQUEST => " + phpQuote(req.Method) + ",\n")
	if req.Body != "" {
		b.WriteString("    CURLOPT_POSTFIELDS => $body,\n")
	}
	if len(req.Headers) > 0 {
		b.WriteString("    CURLOPT_HTTPHEADER => [\n")
		for _, header := range req.Headers {
			b.WriteString("        " + phpQuote(header.Name+": "+header.Value) + ",\n")
		}
		b.WriteString("    ],\n")
	}
	b.WriteString("]);\n\n")
	b.WriteString("$response = curl_exec($ch);\n")
	b.WriteString("if ($response === false) {\n    throw new Exception(curl_error($ch));\n}\n")
	b.WriteString("curl_close($ch);\n\necho $response;\n")
	return b.String()
}

func renderPythonRequest(req httpRequest) string {
	jsonBody, isJSON := parseJSONBody(req.Body)
	var b strings.Builder
	b.WriteString("import requests\n\n")
	if isJSON {
		b.WriteString("import json\n\n")
	}
	if len(req.Headers) > 0 {
		b.WriteString("headers = {\n")
		for _, header := range req.Headers {
			b.WriteString("    " + strconv.Quote(header.Name) + ": " + strconv.Quote(header.Value) + ",\n")
		}
		b.WriteString("}\n\n")
	}
	if isJSON {
		b.WriteString("payload = " + pythonValue(jsonBody, 0) + "\n")
		b.WriteString("body = json.dumps(payload)\n\n")
	} else if req.Body != "" {
		b.WriteString("body = " + strconv.Quote(req.Body) + "\n\n")
	}
	b.WriteString("response = requests.request(\n")
	b.WriteString("    " + strconv.Quote(req.Method) + ",\n")
	b.WriteString("    " + strconv.Quote(req.URL))
	if len(req.Headers) > 0 {
		b.WriteString(",\n    headers=headers")
	}
	if req.Body != "" {
		b.WriteString(",\n    data=body")
	}
	b.WriteString("\n)\n\n")
	b.WriteString("print(response.text)\n")
	return b.String()
}

func renderNodeRequest(req httpRequest) string {
	return renderFetchRequest(req, true)
}

func renderJavaScriptRequest(req httpRequest) string {
	return renderFetchRequest(req, false)
}

func renderFetchRequest(req httpRequest, node bool) string {
	jsonBody, isJSON := parseJSONBody(req.Body)
	var b strings.Builder
	if isJSON {
		b.WriteString("const payload = " + javascriptValue(jsonBody, 0) + ";\n\n")
	} else if req.Body != "" {
		b.WriteString("const body = " + strconv.Quote(req.Body) + ";\n\n")
	}
	b.WriteString("const response = await fetch(" + strconv.Quote(req.URL) + ", {\n")
	b.WriteString("  method: " + strconv.Quote(req.Method))
	if len(req.Headers) > 0 || req.Body != "" {
		b.WriteString(",")
	}
	b.WriteString("\n")
	if len(req.Headers) > 0 {
		b.WriteString("  headers: {\n")
		for i, header := range req.Headers {
			comma := ","
			if i == len(req.Headers)-1 {
				comma = ""
			}
			b.WriteString("    " + strconv.Quote(header.Name) + ": " + strconv.Quote(header.Value) + comma + "\n")
		}
		b.WriteString("  }")
		if req.Body != "" {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	if req.Body != "" {
		if isJSON {
			b.WriteString("  body: JSON.stringify(payload)\n")
		} else {
			b.WriteString("  body\n")
		}
	}
	b.WriteString("});\n\n")
	if node {
		b.WriteString("console.log(await response.text());\n")
	} else {
		b.WriteString("const text = await response.text();\nconsole.log(text);\n")
	}
	return b.String()
}

func renderJavaRequest(req httpRequest) string {
	jsonBody, isJSON := parseJSONBody(req.Body)
	var b strings.Builder
	b.WriteString("import java.net.URI;\n")
	b.WriteString("import java.net.http.HttpClient;\n")
	b.WriteString("import java.net.http.HttpRequest;\n")
	b.WriteString("import java.net.http.HttpResponse;\n")
	if isJSON {
		b.WriteString("import com.fasterxml.jackson.databind.ObjectMapper;\n")
		b.WriteString("import java.util.Map;\n")
		b.WriteString("import java.util.List;\n")
	}
	b.WriteString("\npublic class Main {\n")
	b.WriteString("    public static void main(String[] args) throws Exception {\n")
	if isJSON {
		b.WriteString("        var payload = " + javaValue(jsonBody) + ";\n")
		b.WriteString("        var body = new ObjectMapper().writeValueAsString(payload);\n\n")
	} else if req.Body != "" {
		b.WriteString("        var body = " + javaTextBlock(req.Body) + ";\n\n")
	}
	bodyPublisher := "HttpRequest.BodyPublishers.noBody()"
	if req.Body != "" {
		bodyPublisher = "HttpRequest.BodyPublishers.ofString(body)"
	}
	b.WriteString("        var request = HttpRequest.newBuilder()\n")
	b.WriteString("            .uri(URI.create(" + strconv.Quote(req.URL) + "))\n")
	b.WriteString("            .method(" + strconv.Quote(req.Method) + ", " + bodyPublisher + ")\n")
	for _, header := range req.Headers {
		b.WriteString("            .header(" + strconv.Quote(header.Name) + ", " + strconv.Quote(header.Value) + ")\n")
	}
	b.WriteString("            .build();\n\n")
	b.WriteString("        var response = HttpClient.newHttpClient().send(request, HttpResponse.BodyHandlers.ofString());\n")
	b.WriteString("        System.out.println(response.body());\n")
	b.WriteString("    }\n}")
	return b.String()
}

func phpQuote(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "\\'")
	s = strings.ReplaceAll(s, "\n", "\\n")
	return "'" + s + "'"
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'"
}

func parseJSONBody(body string) (any, bool) {
	if strings.TrimSpace(body) == "" {
		return nil, false
	}
	var v any
	decoder := json.NewDecoder(strings.NewReader(body))
	decoder.UseNumber()
	if err := decoder.Decode(&v); err != nil {
		return nil, false
	}
	return v, true
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func goValue(v any) string {
	switch x := v.(type) {
	case map[string]any:
		parts := []string{}
		for _, key := range sortedKeys(x) {
			parts = append(parts, strconv.Quote(key)+": "+goValue(x[key]))
		}
		return "map[string]any{" + strings.Join(parts, ", ") + "}"
	case []any:
		parts := []string{}
		for _, item := range x {
			parts = append(parts, goValue(item))
		}
		return "[]any{" + strings.Join(parts, ", ") + "}"
	case string:
		return strconv.Quote(x)
	case json.Number:
		return x.String()
	case bool:
		if x {
			return "true"
		}
		return "false"
	case nil:
		return "nil"
	default:
		return fmt.Sprint(x)
	}
}

func phpValue(v any) string {
	switch x := v.(type) {
	case map[string]any:
		parts := []string{}
		for _, key := range sortedKeys(x) {
			parts = append(parts, phpQuote(key)+" => "+phpValue(x[key]))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case []any:
		parts := []string{}
		for _, item := range x {
			parts = append(parts, phpValue(item))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case string:
		return phpQuote(x)
	case json.Number:
		return x.String()
	case bool:
		if x {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	default:
		return phpQuote(fmt.Sprint(x))
	}
}

func pythonValue(v any, level int) string {
	switch x := v.(type) {
	case map[string]any:
		parts := []string{}
		for _, key := range sortedKeys(x) {
			parts = append(parts, strconv.Quote(key)+": "+pythonValue(x[key], level+1))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case []any:
		parts := []string{}
		for _, item := range x {
			parts = append(parts, pythonValue(item, level+1))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case string:
		return strconv.Quote(x)
	case json.Number:
		return x.String()
	case bool:
		if x {
			return "True"
		}
		return "False"
	case nil:
		return "None"
	default:
		return strconv.Quote(fmt.Sprint(x))
	}
}

func javascriptValue(v any, level int) string {
	switch x := v.(type) {
	case map[string]any:
		parts := []string{}
		for _, key := range sortedKeys(x) {
			parts = append(parts, strconv.Quote(key)+": "+javascriptValue(x[key], level+1))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case []any:
		parts := []string{}
		for _, item := range x {
			parts = append(parts, javascriptValue(item, level+1))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case string:
		return strconv.Quote(x)
	case json.Number:
		return x.String()
	case bool:
		if x {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	default:
		return strconv.Quote(fmt.Sprint(x))
	}
}

func javaValue(v any) string {
	switch x := v.(type) {
	case map[string]any:
		parts := []string{}
		for _, key := range sortedKeys(x) {
			parts = append(parts, strconv.Quote(key)+", "+javaValue(x[key]))
		}
		return "Map.of(" + strings.Join(parts, ", ") + ")"
	case []any:
		parts := []string{}
		for _, item := range x {
			parts = append(parts, javaValue(item))
		}
		return "List.of(" + strings.Join(parts, ", ") + ")"
	case string:
		return strconv.Quote(x)
	case json.Number:
		return x.String()
	case bool:
		if x {
			return "true"
		}
		return "false"
	case nil:
		return "null"
	default:
		return strconv.Quote(fmt.Sprint(x))
	}
}

func javaTextBlock(s string) string {
	return `"""` + "\n" + strings.ReplaceAll(s, `"""`, `\"\"\"`) + "\n" + `"""`
}

func hasVersions(node *Node) bool {
	if node == nil {
		return false
	}

	if node.Name == "{version}" {
		return true
	}

	return hasVersions(node.Parent)
}

// return the best variation for a given language and version
func getBestVariation(variations []*Variation, language, version string) *Variation {
	var variation *Variation

	// choose best possible variation

	for _, v := range variations {
		if v.Language == language && v.Version == version {
			variation = v
			break
		}
	}

	// fallback by version
	if variation == nil {
		// todo: sort and filter by version (should be less or eq than "version")
		for _, v := range variations {
			if v.Version == version {
				variation = v
				break
			}
		}
	}

	// fallback by language
	if variation == nil {
		for _, v := range variations {
			if v.Language == language {
				variation = v
				break
			}
		}
	}

	// fallback by any variation (version must be less or equal!)
	if variation == nil {
		for _, v := range variations {
			variation = v
			break
		}
	}

	return variation
}

var basepath = "/"

func getLink(n *Node, lang, version string) string {
	variation := getBestVariation(n.Variations, lang, version)
	return path.Join(basepath, getOutputPath(n, variation, lang, version))
}

func getBreadcrumb(n *Node, lang, version string) string {
	breadcrumb := []*Node{}

	for n != nil && len(n.Variations) > 0 {
		if n.Parent == nil {
			break
		}
		breadcrumb = append(breadcrumb, n)
		n = n.Parent
	}

	if len(breadcrumb) < 2 {
		return ""
	}

	slices.Reverse(breadcrumb)

	result := ""

	result += `<div class="breadcrumb">`
	for i, node := range breadcrumb {
		if node.Name == "{version}" {
			continue
		}
		if i > 0 {
			result += `<span class="arrow">→</span>`
		}
		v := getBestVariation(node.Variations, lang, version)
		class := "item"
		if i == len(breadcrumb)-1 {
			class += " selected"
		}
		result += `<a class="` + class + `" href="` + getLink(node, lang, version) + `">` + v.Title + `</a>`
	}
	result += `</div>`

	return result
}

func getIndex(root, target *Node, lang, version string) string {

	nodesToParent := []*Node{}
	n := target
	for n != nil {
		nodesToParent = append(nodesToParent, n)
		n = n.Parent
	}

	result := ""

	for _, child := range root.Children {

		if child.Name == "{version}" {
			result += getIndex(child, target, lang, version)
			continue
		}

		link := getLink(child, lang, version)

		variation := getBestVariation(child.Variations, lang, version)

		class := "item"
		if nodeIn(nodesToParent, child) {
			class += " active"
		}
		if child == target {
			class += " selected"
		}

		result += `<div class="` + class + `"><a href="` + link + `">` + variation.Title + `</a></div>` + "\n"

		if len(child.Children) == 0 {
			continue
		}

		result += `<div class="children">` + "\n"
		result += getIndex(child, target, lang, version)
		result += `</div>` + "\n"
	}

	return result
}

func traverseNodes(root *Node, callback func(*Node)) {

	callback(root)
	for _, child := range root.Children {
		// callback(child)
		traverseNodes(child, callback)
	}
}

type Node struct {
	Order      int
	Name       string
	Path       string
	Children   []*Node
	Variations []*Variation
	Parent     *Node
	Template   string
}

type Variation struct {
	Url      string
	Language string
	Version  string
	Filename string
	Title    string
}

func (n *Node) PrettyPrint(indent int) {
	for _, child := range n.Children {
		fmt.Printf("%s (%d)\n", strings.Repeat("    ", indent)+child.Name, len(child.Variations))
		child.PrettyPrint(indent + 1)
	}
}

func getOutputPath(node *Node, variation *Variation, lang, version string) string {

	if variation == nil {
		return "NIL VARIATION!!! panic(?)"
	}

	result := []string{}

	for node != nil && node.Parent != nil {

		p := ""

		for _, v := range node.Variations {
			// todo: take version into account :S
			if v.Language == variation.Language {
				p = v.Url
				break
			}
		}

		if p == "" {
			for _, v := range node.Variations {
				if v.Version == variation.Version {
					p = v.Url
					break
				}
			}
		}

		if p == "" {
			p = node.Name // fallback
		}

		if p == "{version}" {
			p = version
		}

		result = append([]string{p}, result...)

		node = node.Parent
	}

	defaultLanguage := languages[0]
	if lang != defaultLanguage {
		result = append([]string{lang}, result...)
	}

	result = append(result, "index.html")

	return path.Join(result...)
}

func copyFile(src, dst string) {
	info, err := os.Stat(src)
	if err != nil {
		panic(err.Error())
	}

	if info.IsDir() {
		err = os.MkdirAll(dst, 0777)
		if err != nil {
			panic(err.Error())
		}

		entries, err := os.ReadDir(src)
		if err != nil {
			panic(err.Error())
		}

		for _, entry := range entries {
			copyFile(path.Join(src, entry.Name()), path.Join(dst, entry.Name()))
		}
		return
	}

	d, err := os.Create(dst)
	if err != nil {
		panic(err.Error())
	}
	s, err := os.Open(src)
	if err != nil {
		panic(err.Error())
	}
	_, err = io.Copy(d, s)
	if err != nil {
		panic(err.Error())
	}
	d.Close()
	s.Close()
}

func readNodes(root *Node, src, www string) { // todo: return errors instead of miserably panic
	entries, err := os.ReadDir(src)
	if err != nil {
		panic(err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			var order int
			var name string
			if entry.Name() == "{version}" {
				name = entry.Name()
			} else {
				parts := strings.Split(entry.Name(), "_")
				if len(parts) != 2 {
					copyFile(path.Join(src, entry.Name()), path.Join(www, entry.Name()))
					continue
				}
				order, err = strconv.Atoi(parts[0])
				if err != nil {
					continue
				}
				name = parts[1]
			}

			newNode := &Node{
				Order: order,
				Name:  name,
				Path:  src,
			}
			readNodes(newNode, path.Join(src, entry.Name()), www)

			newNode.Parent = root
			root.Children = append(root.Children, newNode)

		} else {
			ext := strings.ToLower(path.Ext(entry.Name()))
			if ext == ".gohtml" {
				root.Template = path.Join(src, entry.Name())
				continue
			}
			if !in([]string{".html", ".md"}, ext) {
				copyFile(path.Join(src, entry.Name()), path.Join(www, entry.Name()))
				continue
			}

			base := strings.ToLower(strings.TrimSuffix(path.Base(entry.Name()), path.Ext(entry.Name())))
			parts := strings.Split(base, "_")

			if len(parts) == 1 {
				copyFile(path.Join(src, entry.Name()), path.Join(www, entry.Name()))
				continue
			}

			friendlyUrl := parts[0]
			lang := ""
			version := ""

			for _, p := range parts[1:] {
				p = strings.ToLower(p)
				if in(languages, p) {
					lang = p
				}
				if in(versions, p) {
					version = p
				}
			}
			parts = parts[1:]

			filename := path.Join(src, entry.Name())

			title := getTitle(filename)
			if title == "" {
				title = friendlyUrl // fallback
				base, _ := os.Getwd()
				fmt.Printf("WARNING: %s:1 needs a title <h1>\n", path.Join(base, filename))
			}

			root.Variations = append(root.Variations, &Variation{
				Url:      friendlyUrl,
				Language: lang,
				Version:  version,
				Filename: filename,
				Title:    title,
			})

		}

	}

	sort.Slice(root.Children, func(i, j int) bool {
		return root.Children[i].Order < root.Children[j].Order
	})

}

func getTitle(filename string) string {

	var htmlReader io.Reader
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	switch strings.ToLower(path.Ext(filename)) {
	case ".html":

		htmlReader = f

	case ".md":
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err.Error())
		}

		htmlReader = md2html(b)
	}

	title := ""
	doc, err := html.Parse(htmlReader)
	if err != nil {
		panic(err.Error())
	}

	traverseHtml(doc, func(node *html.Node) {
		if node.Data == "h1" && node.FirstChild != nil {
			title = node.FirstChild.Data
		}
	})

	return title
}

func traverseHtml(n *html.Node, callback func(node *html.Node)) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHtml(c, callback)
	}
	if n.Type == html.ElementNode {
		callback(n)
	}
}

func in[T cmp.Ordered](a []T, v T) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

func nodeIn(a []*Node, v *Node) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}

func md2html(md []byte) io.Reader {

	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
		// parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkHtml.WithXHTML(),
			goldmarkHtml.WithUnsafe(),
		),
	)

	buf := &bytes.Buffer{}
	err := gm.Convert(md, buf)
	if err != nil {
		panic(err.Error())
	}
	return buf
}
