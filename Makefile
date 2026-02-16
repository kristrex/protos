install-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

install-deps-lat:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen-sso:
	protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-profile:
	protoc -I proto proto/profile/profile.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-referral:
	protoc -I proto proto/referral/referral.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

gen-balance:
	protoc -I proto proto/balance/balance.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative