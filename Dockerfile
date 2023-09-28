FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/taivama/golang-training
COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /golang-training

FROM scratch
COPY --from=builder /golang-training /golang-training
ENTRYPOINT ["/golang-training"]
