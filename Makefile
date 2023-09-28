.PHONY: build
build:
	go build

.PHONY: run
run:
	go run server.go

.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html
	rm -f coverage.out

.PHONY: proto
proto:
	rm -f proto/*.pb.go
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: docker
docker:
	docker build -t golang-training:latest .

.PHONY: certs
certs:
	rm -fr certs
	mkdir certs
	cd certs; minica -ca-cert ca-cert.pem -ca-key ca-key.pem -domains localhost -ip-addresses 127.0.0.1; cd ..
