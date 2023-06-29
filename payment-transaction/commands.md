## to list running services and their ports
ss -nltp

## protoc to generate stubs for gRPC
### message data
protoc --proto_path=./adapter/grpc/proto ./adapter/grpc/proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./adapter/grpc/pb

### grpc data
protoc --proto_path=./adapter/grpc/proto ./adapter/grpc/proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=./adapter/grpc/pb


## Evans (gRPC Client)
<!-- 0.0.0.0:3160->50051/tcp, :::3160->50051/tcp -->
### from inside the container
evans -r repl --port=50051 

### from outside the container
evans -r repl --port=3160

## calling the service named Process (CLI Evans)
call Process
