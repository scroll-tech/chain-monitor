.PHONY: abi update docker update format lint chain-monitor

IMAGE_NAME=chain-monitor
IMAGE_VERSION=latest

L2GETH_TAG=scroll-v5.1.5

lint:
	go mod tidy
	goimports -w .
	gofumpt -l -w .
	GOBIN=$(PWD)/build/bin go run ./build/lint.go

abi:
	bash ./build_abi.sh $(CPATH)

chain-monitor:
	go build -o build/bin/chain-monitor ./cmd/main.go

test:
	go test -v -race -gcflags="-l" -ldflags="-s=false" -coverprofile=coverage.txt -covermode=atomic ./...

docker:
	docker build --platform linux/amd64 -t scrolltech/${IMAGE_NAME}:${IMAGE_VERSION} ./

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@${L2GETH_TAG}
	go mod tidy
