#!/usr/bin/env python3
"""Checks the InceptionDB documentation.

  check_inceptiondb_docs.py lint     sources: forbidden hosts and operators,
                                     Content-Type, same code in every language
  check_inceptiondb_docs.py links    generated site: links, canonical URLs and
                                     the same navigation on every page
  check_inceptiondb_docs.py run LANG runs every curl of the docs and compares
                                     the answers with the documented ones

`run` needs INCEPTIONDB_DATABASE, INCEPTIONDB_API_KEY and INCEPTIONDB_API_SECRET
of an EMPTY test database (it refuses to start otherwise) and deletes what it
creates, even when a check fails. INCEPTIONDB_URL sends the requests to another
server, like a local one. Secrets are never printed.

The guides run in order, without resetting anything between them, so they must
work one after the other as a reader follows them. A code block preceded by
<!-- docs-check: skip --> is not run (for example, cleanup of a previous try).
"""
import glob
import html.parser
import json
import os
import re
import subprocess
import sys
import tempfile

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
SRC = os.path.join(ROOT, 'src/30_docs/10_inceptiondb')
PAGES_DIR = os.path.join(SRC, '{version}')
LANDING = os.path.join(ROOT, 'src/50_inceptiondb')
WWW = os.path.join(ROOT, 'statics/www')
LANGS = ['en', 'es', 'zh']
BASE = 'https://inceptiondb.hola.cloud'
SKIP = '<!-- docs-check: skip -->'

# The guides, in reading order. They share the `customers` collection.
TOUR = ['1_getting-started', '2_managing-collections', '3_working-with-data',
        '4_querying', '5_indexes', '6_api-keys']

SEED3 = '\n'.join([
    '{"id":"ada","name":"Ada Lovelace","email":"ada@example.com","plan":"pro","country":"GB","mrr":120,"tags":["vip","beta"],"address":{"city":"London","zip":"N1"}}',
    '{"id":"alan","name":"Alan Turing","email":"alan@example.com","plan":"free","country":"GB","mrr":0,"tags":["beta"],"address":{"city":"Manchester","zip":"M1"}}',
    '{"id":"grace","name":"Grace Hopper","email":"grace@example.com","plan":"team","country":"US","mrr":480,"tags":["vip"],"address":{"city":"New York","zip":"10001"}}'])
SHORT = '{"id":"ada","name":"Ada Lovelace","plan":"pro","mrr":120}\n{"id":"alan","name":"Alan Turing","plan":"free","mrr":0}'

# Reference pages and the `customers` they expect (None: no collection).
REFERENCE = [
    ('9_api-reference/5_create-collection', None),
    ('9_api-reference/10_list-collections', SEED3),
    ('9_api-reference/14_collection-management', SEED3),
    ('9_api-reference/17_insert-documents', ''),
    ('9_api-reference/18_find-documents', SHORT),
    ('9_api-reference/21_get-document', SHORT),
    ('9_api-reference/19_patch-documents', SHORT),
    ('9_api-reference/20_remove-documents', SHORT),
    ('9_api-reference/15_index-management', SEED3),
]

FENCE = re.compile(r'(<!-- docs-check: skip -->\s*)?```(\w*)\n(.*?)```', re.S)


def page_file(page, lang):
    found = glob.glob(os.path.join(PAGES_DIR, page, f'*_{lang}.md'))
    return found[0] if found else None


def all_pages():
    pages = []
    for f in glob.glob(os.path.join(PAGES_DIR, '**/*_en.md'), recursive=True):
        pages.append(os.path.relpath(os.path.dirname(f), PAGES_DIR))
    return sorted(pages)


def blocks(path):
    """(skip, lang, code) of every fenced block of a markdown file."""
    return [(bool(m.group(1)), m.group(2), m.group(3)) for m in FENCE.finditer(open(path, encoding='utf-8').read())]


# ---------------------------------------------------------------- lint

def lint():
    problems = []
    sources = glob.glob(os.path.join(SRC, '**/*.md'), recursive=True) + glob.glob(os.path.join(LANDING, '*.html'))
    for f in sources:
        text = open(f, encoding='utf-8').read()
        rel = os.path.relpath(f, ROOT)
        if 'api.hola.cloud' in text.replace('inceptiondb.hola.cloud', ''):
            problems.append(f'{rel}: uses the old host api.hola.cloud')
        for code in re.findall(r'```\w*\n(.*?)```|<pre><code>(.*?)</code></pre>', text, re.S):
            code = ''.join(code)
            if re.search(r'\$(gte|lte)\b', code):
                problems.append(f'{rel}: $gte or $lte in a code block; the operators are $ge and $le')
            for curl in re.findall(r'curl .*?(?:\n(?!\s)|$)', code, re.S):
                if re.search(r"\s(-d|--data-binary) ", curl) and 'Content-Type:' not in curl:
                    problems.append(f'{rel}: a curl with a body has no Content-Type: {curl.splitlines()[0][:80]}')

    for page in all_pages():
        reference = [(lang, code) for _, lang, code in blocks(page_file(page, 'en'))]
        for lang in LANGS[1:]:
            path = page_file(page, lang)
            if not path:
                problems.append(f'{page}: no {lang} version')
                continue
            other = [(l, code) for _, l, code in blocks(path)]
            if other != reference:
                problems.append(f'{page}: the code blocks of {lang} differ from en')

    for p in problems:
        print('FAIL', p)
    print(f'lint: {len(problems)} problems')
    return not problems


# ---------------------------------------------------------------- links

class Links(html.parser.HTMLParser):
    def __init__(self):
        super().__init__()
        self.hrefs, self.canonical, self.nav, self.item = [], None, [], None

    def handle_starttag(self, tag, attrs):
        a = dict(attrs)
        if tag == 'link' and a.get('rel') == 'canonical':
            self.canonical = a.get('href')
        if tag == 'div' and (a.get('class') or '').split()[:1] == ['item']:
            self.item = a.get('class')
        if tag == 'a' and a.get('href'):
            self.hrefs.append(a['href'])
            if self.item is not None:
                if '/docs/inceptiondb/' in a['href']:
                    self.nav.append((a['href'], 'selected' in self.item.split()))
                self.item = None


def links():
    problems = []
    pages = []
    for prefix in ('', 'es/', 'zh/'):
        pages += glob.glob(os.path.join(WWW, prefix + 'docs/inceptiondb/**/index.html'), recursive=True)
    navs = {}
    for f in pages:
        url = '/' + os.path.relpath(f, WWW)[:-len('index.html')]
        p = Links()
        p.feed(open(f, encoding='utf-8').read())
        if p.canonical != 'https://hola.cloud' + url:
            problems.append(f'{url}: canonical is {p.canonical}')
        for href in p.hrefs:
            if not href.startswith('/') or href.startswith('//'):
                continue
            target = href.split('#')[0].split('?')[0]
            if target.endswith('index.html') and '/docs/' in target:
                problems.append(f'{url}: links to {href} instead of its canonical URL')
            on_disk = os.path.join(WWW, target.lstrip('/'))
            if not (os.path.isfile(on_disk) or os.path.isfile(os.path.join(on_disk, 'index.html'))):
                problems.append(f'{url}: broken link {href}')
        selected = [h for h, sel in p.nav if sel]
        if selected != [url]:
            problems.append(f'{url}: the navigation marks {selected} as the current page')
        lang_version = (url.split('/docs/')[0], (re.findall(r'/(v\d+)/', url) or [''])[0])
        navs.setdefault(lang_version, {})[url] = [h for h, _ in p.nav]

    for key, by_page in navs.items():
        variants = {tuple(n) for n in by_page.values()}
        if len(variants) > 1:
            problems.append(f'{key}: pages show {len(variants)} different InceptionDB menus')
        listed = {h for n in by_page.values() for h in n}
        for url in by_page:
            if url not in listed:
                problems.append(f'{url}: not in the navigation')

    for p in problems:
        print('FAIL', p)
    print(f'links: {len(pages)} pages, {len(problems)} problems')
    return not problems


# ---------------------------------------------------------------- run

VOLATILE = re.compile(r'^[0-9a-f]{8}-[0-9a-f]{4}-|^\d{4}-\d\d-\d\dT')


def norm(v, key=''):
    if isinstance(v, dict):
        return {k: norm(x, k) for k, x in v.items() if not k.startswith('index.')}
    if isinstance(v, list):
        return [norm(x) for x in v]
    if isinstance(v, str) and VOLATILE.match(v):
        return '<volatile>'
    if key in ('created', 'disk', 'memory'):
        return '<number>'
    return v


def parse(text):
    try:
        return [norm(json.loads(line)) for line in text.strip().splitlines()]
    except ValueError:
        return None


def canon(docs):
    if len(docs) == 1 and isinstance(docs[0], list):
        docs = docs[0]
    return sorted(json.dumps(d, sort_keys=True) for d in docs)


class Runner:
    def __init__(self, lang):
        self.lang = lang
        self.url = os.environ.get('INCEPTIONDB_URL', BASE).rstrip('/')
        self.db = os.environ['INCEPTIONDB_DATABASE']
        self.headers = ['-H', 'Api-Key: ' + os.environ['INCEPTIONDB_API_KEY'],
                        '-H', 'Api-Secret: ' + os.environ['INCEPTIONDB_API_SECRET']]
        self.collections = f'{self.url}/v1/databases/{self.db}/collections'
        self.cwd = tempfile.mkdtemp()
        self.failures = 0

    def call(self, *args):
        out = subprocess.run(['curl', '-s', '-w', '\n%{http_code}', *args, *self.headers],
                             capture_output=True, text=True).stdout
        body, _, status = out.rpartition('\n')
        return status, body

    def names(self):
        status, body = self.call(self.collections)
        if status != '200':
            sys.exit(f'listing the collections answered {status}: check the credentials')
        return [c['name'] for c in json.loads(body)]

    def seed(self, docs):
        self.call('-X', 'POST', f'{self.collections}/customers:dropCollection')
        if docs is None:
            return
        self.call('-X', 'POST', self.collections, '-d', '{"name":"customers"}')
        if docs:
            self.call('-X', 'POST', f'{self.collections}/customers:insert', '--data-binary', docs)

    def check(self, ok, what):
        print(f'  {"OK " if ok else "BAD"} {what}')
        if not ok:
            self.failures += 1

    def page(self, page):
        path = page_file(page, self.lang)
        print(f'== {page}')
        if not path:
            self.check(False, f'no {self.lang} version')
            return
        found = blocks(path)
        for i, (skip, _, code) in enumerate(found):
            if skip or not code.lstrip().startswith('curl'):
                continue
            expected = None
            if i + 1 < len(found) and found[i + 1][1] in ('json', 'jsonl'):
                expected = found[i + 1][2]
            cmd = code.replace(BASE, self.url).strip() + " -s -w '\\n%{http_code}'"
            out = subprocess.run(['bash', '-c', cmd], capture_output=True, text=True, cwd=self.cwd).stdout
            body, _, status = out.rpartition('\n')
            first = code.strip().splitlines()[0][:100]
            diff = ''
            if status.startswith('2') and expected is not None:
                got, want = parse(body), parse(expected)
                if want is not None and (got is None or canon(got) != canon(want)):
                    diff = f'\n      expected: {expected.strip()[:300]}\n      got:      {body.strip()[:300]}'
            self.check(status.startswith('2') and not diff, f'{status} {first}{diff}')

    def run(self):
        existing = self.names()
        if existing:
            sys.exit(f'the database must be empty, it has {len(existing)} collections')
        try:
            for page in TOUR:
                self.page(page)
            print('== state after the guides')
            status, body = self.call(f'{self.collections}/customers')
            info = json.loads(body) if status == '200' else {}
            self.check(self.names() == ['customers'], 'only customers is left')
            self.check(info.get('total') == 2 and info.get('indexes') == 0, 'customers has Ada and Grace, without indexes')
            status, _ = self.call('-X', 'POST', f'{self.collections}/customers:find', '-d', '{"filter":{"mrr":{"$gte":100}}}')
            self.check(status == '204', 'an unknown operator matches nothing (204), as documented')
            for page, docs in REFERENCE:
                self.seed(docs)
                self.page(page)
        finally:
            for name in self.names():
                self.call('-X', 'POST', f'{self.collections}/{name}:dropCollection')
            self.check(self.names() == [], 'clean up: the database is empty again')
        print(f'run {self.lang}: {self.failures} failures')
        return self.failures == 0


def main():
    if len(sys.argv) < 2 or sys.argv[1] not in ('lint', 'links', 'run'):
        sys.exit(__doc__)
    if sys.argv[1] == 'lint':
        ok = lint()
    elif sys.argv[1] == 'links':
        ok = links()
    else:
        ok = Runner(sys.argv[2] if len(sys.argv) > 2 else 'en').run()
    sys.exit(0 if ok else 1)


if __name__ == '__main__':
    main()
