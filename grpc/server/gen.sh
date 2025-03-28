export PATH="/home/vscode/.bun/bin:$PATH"

cd proto

protoc --ts_out=../../gen --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   math.proto
