DOCKER_REGISTRY ?= jozuenoon

GIT_BRANCH := $(shell git branch | sed -n '/\* /s///p' 2>/dev/null)
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null)


# This docker build produce intermittent tags tied to branch and commit. Useful while developing and
# making deployments to test environment.
build_docker: cqserver_docker

cqserver_docker:
	docker build -f cmd/cq/Dockerfile -t $(DOCKER_REGISTRY)/cqserver:$(GIT_BRANCH)_$(GIT_COMMIT) -t $(DOCKER_REGISTRY)/cqserver:latest .

test:
	go test -race ./...

.PHONY: bin
bin:
	go build -o bin/cqserver cmd/cq/main.go
	go build -o bin/mbcli mbcli/main.go

install:
	go build -o bin/mbcli mbcli/main.go
	chmod +x bin/mbcli
	cp bin/mbcli $$HOME/bin/mbcli

# Generators
gen: collector query

## Collector
collector: collector/collector.pb.go collector/mocks/repository.go

collector/collector.pb.go: proto/collector.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./collector

collector/mocks/repository.go: collector/repository.go
	mockery -name=Repository -dir ./collector -output ./collector/mocks -case snake

## Query
query: query/query.pb.go query/mocks/repository.go

query/query.pb.go: proto/query.proto
	protoc -I proto/ $< --go_out=plugins=grpc:./query

query/mocks/repository.go: query/repository.go
	mockery -name=Repository -dir ./query -output ./query/mocks -case snake

deployment-check: deployment-cqserver-check

deployment-cqserver-check:
	helm install deployment/cqserver/ --debug --dry-run