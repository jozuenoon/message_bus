all: collector query


collector: collector/collector.pb.go collector/mocks/repository.go

collector/collector.pb.go: proto/collector.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./collector

collector/mocks/repository.go: collector/repository.go
	mockery -name=Repository -dir ./collector -output ./collector/mocks -case snake


query: query/query.pb.go query/mocks/repository.go

query/query.pb.go: proto/query.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./query

query/mocks/repository.go: query/repository.go
	mockery -name=Repository -dir ./query -output ./query/mocks -case snake