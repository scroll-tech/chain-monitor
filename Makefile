.PHONY: submodule abi update docker

IMAGE_NAME=chain-monitor
IMAGE_VERSION=latest

# scroll path is used when update abi.
SCROLL_PATH=-1

chain-monitor:
	go build -o build/bin/chain-monitor ./cmd/main.go

abi:
	cd $(SCROLL_PATH)/contracts && yarn install && forge build
	make -C bytecode scroll $(SCROLL_PATH)

docker:
	docker build --platform linux/amd64 -t scrolltech/${IMAGE_NAME}:${IMAGE_VERSION} ./

update: ## Let's keep it and docker version in consistent.
	#git submodule update --init --recursive
	go get -u github.com/scroll-tech/go-ethereum@scroll-v4.3.52
	go mod tidy
