.PHONY: submodule abi

submodule:
	git submodule update --init --recursive

chain-monitor:
	go build -o build/bin/chain-monitor ./cmd/main.go

abi:
	make -C bytecode scroll