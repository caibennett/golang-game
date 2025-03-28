sudo apt install -y protobuf-compiler nodejs
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/ktr0731/evans@latest

bun add protoc-gen-ts -g