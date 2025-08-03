
gen:
	@protoc \
		--proto_path=protobuf "app/protobuf/users.proto" \
		--go_out=app/internal/services/common/genproto/users --go_opt=paths=source_relative \
  	--go-grpc_out=app/internal/services/common/genproto/users --go-grpc_opt=paths=source_relative