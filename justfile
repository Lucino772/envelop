gen-proto:
    buf generate

tidy:
    go fmt ./...
    go mod tidy