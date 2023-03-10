protoc:
	@echo "go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26"
	@echo "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1"
	protoc --go_out=./ --go-grpc_out=./ ./api/*.proto