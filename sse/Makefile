build-buf:
#	protoc -I=app/user/protobuf --go_out=.  --proto_path=app/user/protobuf  --go-grpc_out=require_unimplemented_servers=false:. app/user/protobuf/*.proto
	protoc -I=./third_party/googleapis --go_out=. --grpc-gateway_out=.  --proto_path=proto  --go-grpc_out=require_unimplemented_servers=false:. proto/person.proto
#	protoc -I=app/task/protobuf --go_out=.  --proto_path=app/task/protobuf  --go-grpc_out=require_unimplemented_servers=false:. app/task/protobuf/*.proto
