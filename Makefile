DOCKER_REGISTRY ?= jozuenoon

GIT_BRANCH := $(shell git branch | sed -n '/\* /s///p' 2>/dev/null)
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null)


# This docker build produce intermittent tags tied to branch and commit. Useful while developing and
# making deployments to test environment.
build_docker: cqserver_docker

cqserver_docker:
	# eval $(minikube docker-env)
	docker build -f cmd/cq/Dockerfile -t $(DOCKER_REGISTRY)/message_bus:$(GIT_BRANCH)_$(GIT_COMMIT) -t $(DOCKER_REGISTRY)/message_bus:latest .


test:
	go test -race ./...

integration:
	ETCD_ENDPOINTS=http://127.0.0.1:2379 go test -tags integration -race ./...

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

# NOTE
# With nginx ingress in minikube is now at version 0.23 and have some problems with GRPC request proxying.
# Upgrading to 0.25 should solve this problem. However to workaround that in minikube we set NodePort service
# to get everything working instantly, check the node ports:
# $ kubectl describe svc --namespace tdc message-bus
deploy-minikube:
	helm upgrade message-bus \
	--install \
	--kube-context minikube \
	--set ImageTag=$(GIT_BRANCH)_$(GIT_COMMIT) \
	--set DockerRegistry=$(DOCKER_REGISTRY) \
	--set ServiceType=NodePort
	--namespace=tdc deployment/cqserver

deploy-etcd-operator:
	helm upgrade etcd-operator --install stable/etcd-operator --namespace tdc

deploy-etcd:
	kubectl apply -f deployment/etcd/etcd.yaml --namespace tdc