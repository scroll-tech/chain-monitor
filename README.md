# chain-monitor

Detect real-time threats and events on Scroll blockchains. chain monitor can
detect ETH, ERC20, ERC721, ERC1155. As for erc20 can support WETH, StandardERC20,
CustomERC20, USDC, DAI, LIDO.

Detect features:
1. L2 withdraw root message hash check.
2. ERC20, ERC721, ERC1155's token id and amount check.
3. ETH balance check.
4. Event which happened on L1/L2 can match.

# Dependencies

* solc

> The detail tutorial please follow [here](https://docs.soliditylang.org/en/latest/installing-solidity.html).

* abigen

```
go install -v github.com/scroll-tech/go-ethereum/cmd/abigen@develop
```

* foundry

```bash
curl -L https://foundry.paradigm.xyz | bash
```

# Build chain-monitor

* Update dependence

```
# compile contracts with foundry and translate abi to go source files 
make abi CPATH=xxx

# updaye golang depedence
make update
```

* Make chain-monitor

```
make 
```