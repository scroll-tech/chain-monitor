.PHONY: submodule abi update docker

IMAGE_NAME=chain-monitor
IMAGE_VERSION=latest

submodule:
	git submodule update --init --recursive

chain-monitor:
	go build -o build/bin/chain-monitor ./cmd/main.go

abi: submodule
	cd scroll/contracts && yarn install && forge build
	make -C bytecode scroll

docker:
	docker build --platform linux/amd64 -t scrolltech/${IMAGE_NAME}:${IMAGE_VERSION} ./

update: ## Let's keep it and docker version in consistent.
	#git submodule update --init --recursive
	go get -u github.com/scroll-tech/go-ethereum@scroll-v4.3.52
	go mod tidy
