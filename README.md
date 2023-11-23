# chain-monitor

It's a internal chain assets monitor service, call by api.

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
make chain-monitor
```