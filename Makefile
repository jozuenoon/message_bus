all: collector/collector.pb.go  query/query.pb.go


collector/collector.pb.go: proto/collector.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./collector

query/query.pb.go: proto/query.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./query
