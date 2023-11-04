ifndef ${IDL_PATH}
	IDL_PATH = ./idl
endif

proto:
	protoc --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative --proto_path=${IDL_PATH} user.proto
