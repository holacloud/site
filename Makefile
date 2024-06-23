GIT_VERSION = $(shell git describe --tags --always)
FLAGS = -ldflags "\
  -X main.VERSION=$(GIT_VERSION) \
"
PROJECT="console"

.PHONY: run
run:
	STATICS=statics/www/ go run  $(FLAGS) ./cmd/site

.PHONY: build
build: site
	CGO_ENABLED=0 go build $(FLAGS) -o bin/ ./cmd/...

.PHONY: release
release: clean
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build $(FLAGS) -o bin/$(PROJECT).linux.arm64 ./cmd/...
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build $(FLAGS) -o bin/$(PROJECT).linux.amd64 ./cmd/...
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build $(FLAGS) -o bin/$(PROJECT).win.arm64.exe ./cmd/...
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(FLAGS) -o bin/$(PROJECT).win.amd64.exe ./cmd/...
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build $(FLAGS) -o bin/$(PROJECT).mac.arm64 ./cmd/...
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build $(FLAGS) -o bin/$(PROJECT).mac.amd64 ./cmd/...
	md5sum bin/* > bin/checksum

.PHONY: clean
clean:
	rm -f bin/*

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

.PHONY: version
version:
	@echo -n "$(GIT_VERSION)"

.PHONY: site
site:
	go run ./cmd/sitegen --src src/ --www statics/www/ --versions "v1,v2" --languages "en,es,zh"

.PHONY: watch
watch:
	inotifywait -r -m src/ -e close_write --format "%T" --timefmt "%S" | uniq

.PHONY: dev
dev:
	go build -o bin/sitegen ./cmd/sitegen
	./watch.sh

