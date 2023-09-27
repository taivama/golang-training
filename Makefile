.PHONY: build
build:
	go build

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html
	rm -f coverage.out
