.PHONY: scroll

SCROLL_PATH=-1

scroll:
	# L1
	./build_abi.sh $(SCROLL_PATH) L1
	# L1/gateway
	./build_abi.sh $(SCROLL_PATH) L1/gateway
	# L1/rollup
	./build_abi.sh $(SCROLL_PATH) L1/rollup
	# L2
	./build_abi.sh $(SCROLL_PATH) L2
	# L2/gateway
	./build_abi.sh $(SCROLL_PATH) L2/gateway
	# L2/predeploys
	./build_abi.sh $(SCROLL_PATH) L2/predeploys
	# token
	./build_abi.sh $(SCROLL_PATH) token
