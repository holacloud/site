
.PHONY: test
test:
	go test -count=1 -cover -race ./...

.PHONY: deps
deps:
	go mod tidy
	go mod vendor

.PHONY: docker-%
docker-%:
	docker-compose run --use-aliases --rm app make $*