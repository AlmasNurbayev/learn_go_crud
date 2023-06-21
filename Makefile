.PHONY: build
build:
	go build -o apiserverP -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL :=build
