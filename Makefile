.PHONY: abi update docker update format lint chain-monitor

IMAGE_NAME=chain-monitor
IMAGE_VERSION=latest

L2GETH_TAG=scroll-v4.3.63

# scroll path is used when update abi.
SCROLL_PATH=-1

format:
	go mod tidy
	goimports -local . -w .

lint:
	GOBIN=$(PWD)/build/bin go run ./build/lint.go

chain-monitor:
	go build -o build/bin/chain-monitor ./cmd/main.go

abi:
	cd $(SCROLL_PATH)/contracts && yarn install && forge build
	make -C bytecode scroll $(SCROLL_PATH)

docker:
	docker build --platform linux/amd64 -t scrolltech/${IMAGE_NAME}:${IMAGE_VERSION} ./

update: ## Let's keep it and docker version in consistent.
	go get -u github.com/scroll-tech/go-ethereum@${L2GETH_TAG} && go mod tidy
	go mod tidy
