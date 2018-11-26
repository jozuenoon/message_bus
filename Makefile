gen:
	protoc --go_out=plugins=grpc:. api/api.proto