module gRPCHandler/user

go 1.16

require (
	Sample_gRPC v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)

replace Sample_gRPC => ../../../sample_gRPC
