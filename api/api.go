package api

import (
	"net/http"

	"github.com/fulldump/box"

	"github.com/holacloud/site/statics"
)

type JSON = map[string]any

func Build(staticsDir, version string) *box.B {

	b := box.NewBox()

	b.WithInterceptors(
		NaiveAccessLog,
		PrettyErrorInterceptor,
	)

	// Version
	b.Handle("GET", "/version", Write(version))

	// Static files (catch-all: must be registered last so specific routes match first)
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).
		WithInterceptors(
			IfModifiedSince(),
		)

	return b
}

func BuildWithComments(staticsDir, version, commentsDir string) *box.B {
	b := box.NewBox()

	b.WithInterceptors(
		NaiveAccessLog,
		PrettyErrorInterceptor,
	)

	// Version
	b.Handle("GET", "/version", Write(version))

	// Draw redirect
	b.Handle("GET", "/draw", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://draw.hola.cloud/")
		w.WriteHeader(http.StatusFound)
	})

	// Comments API
	commentsAPI, err := NewCommentsAPI(commentsDir)
	if err != nil {
		panic("comments api: " + err.Error())
	}

	b.Handle("GET", "/api/comments", commentsAPI.List)
	b.Handle("POST", "/api/comments", commentsAPI.Create)

	// Static files (catch-all: must be registered last so specific routes match first)
	b.Handle("GET", "/*", statics.ServeStatics(staticsDir)).
		WithInterceptors(
			IfModifiedSince(),
		)

	return b
}
