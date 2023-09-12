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
make abi SCROLL_PATH=xxx

# updaye golang depedence
make update
```

* Make chain-monitor

```
make chain-monitor
```

# Explain key fields in config file.

+ l1_config
    + l1chain_url: l1chain node url.
    + l1_gateways: All the deployed gateway contract addresses.
+ l2_config:
    + l2chain_url: l2chain node url.
    + l2_gateways: All the gateway contracts deployed on l2chain, monitor the events.
    + message_queue: Used to get withdraw root from l2chain node.
+ chain_monitor:
    + bridge_history_url: The url used connect to bridge-history API service.

# Cmds explain

```
# show cmd details
./build/bin/chain-monitor --help

# `--init` If first run, we should create tables.
./build/bin/chain-monitor --config config.json --init

# `--http` start http api service
./build/bin/chain-monitor --config config.json --http --http.addr localhost --http.port 8750

# `--pprof` it's used for performance analysis
./build/bin/chain-monitor --config config.json --pprof
```

# API call
```
request:
curl --location 'http://localhost:8750/v1/blocks_status?start_number=528638&end_number=529638'

response:
{
    "errcode": 0,
    "errmsg": "",
    "data": false
}
```