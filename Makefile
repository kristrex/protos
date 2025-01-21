install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/pmkdrotoc-gen-go-grpc@1.5.1
generate-protoc:
	protoc -I protos/proto protos/proto/sso/sso.proto --go_out=./protos/gen/go --go_opt=paths=source_relative --go-grpc_out=./protos/gen/go/ --go-grpc_opt=paths=source_relative