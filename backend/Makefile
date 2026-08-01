PHONY: run test test-verbose audit

run:
	go run .
	
test:
	go test -cover -coverprofile=coverage.out -race ./src/...

test-verbose:
	go test -cover -coverprofile=coverage.out -race -v ./src/...

audit:
	govulncheck ./...