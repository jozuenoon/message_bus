all: collector/collector.pb.go  query/query.pb.go collector/mocks/repository.go


collector/collector.pb.go: proto/collector.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./collector

query/query.pb.go: proto/query.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./query

collector/mocks/repository.go: collector/repository.go
	mockery -name=Repository -dir ./collector -output ./collector/mocks -case snake