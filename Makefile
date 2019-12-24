.PHONY: test lint vet

test: vet lint
	go test ./... -v

vet:
	go vet ./...

lint:
	golangci-lint run