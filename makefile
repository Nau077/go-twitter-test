THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))


.PHONY: run/server
run/server:
	@cd subscribe_service; \
	exec go run ./cmd/main.go;

.PHONY: run/db
run/db:
	docker-compose up neo4j;

.PHONY: run/test
run/test:
	@cd subscribe_service; \
	exec go test -v ./...