Before running the code,please follow https://grpc.io/docs/languages/go/quickstart/ 

To run code-

0. Run : go mod tidy 

1. Go to path sample_gRPC/gRPCHandler/user/protos
2. Run : protoc --go_out=. --go-grpc_out=. user.proto 

3. Go to path sample_gRPC/gRPCHandler/user/Server
4. Run : go run main.go

5. Go to path sample_gRPC/gRPCHandler/user/Client
6. Run : go run main.go
