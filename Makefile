
.PHONY: protobufs
protobufs:
	@protoc --go_out=. \
			--go_opt=module=github.com/Lucino772/envelop \
			--go-grpc_out=. \
			--go-grpc_opt=module=github.com/Lucino772/envelop \
			./resources/protobufs/*.proto


.PHONY: tidy
tidy:
	@go fmt ./...
	@go mod tidy -v

