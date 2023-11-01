#!/bin/bash

# generates go bindings from contracts, to paste into abi/bridge_abi.go
# compile artifacts in /contracts folder with `forge build`` first 

if [ $# != 1 ]; then
  echo "Usage: ./build_abi.sh <contracts_path>"
  exit
fi

contracts_path=$1;
echo "contracts path: $contracts_path"

if [ ! -d "$contracts_path" ]; then
  echo "Directory $contracts don't exists"
  exit
fi

# cd contract path
cd $contracts_path && forge build

# switch to current workdir
cd -

if [ -d contracts_tmp ]; then
  rm -rf contracts_tmp
fi
mkdir contracts_tmp

mv $contracts_path/artifacts ./contracts_tmp

# l1 messenger
l1_messenger=("IL1ScrollMessenger")

# l1 gateway
l1_gateway=("IL1ETHGateway" "IL1ERC20Gateway" "L1ERC721Gateway" "IL1ERC1155Gateway")

# l2 messenger
l2_messenger=("IL2ScrollMessenger")

# l2 gateway
l2_gateway=("IL2ETHGateway" "IL2ERC20Gateway" "L2ERC721Gateway" "L2ERC1155Gateway")

# token list
token_list=("IScrollERC20" "IScrollERC721" "IScrollERC1155")

if [ -d "internal/logic/contracts/abi/" ]; then
  rm -rf internal/logic/contracts/abi/
fi
mkdir internal/logic/contracts/abi/

mkdir -p tmp
contracts=("${l1_messenger[@]}" "${l1_gateway[@]}" "${l2_messenger[@]}" "${l2_gateway[@]}" "${token_list[@]}")
for abi_name in "${contracts[@]}"; do
  echo "Generating code for: $abi_name"

  abi="tmp/$abi_name.json"
  cat ./contracts_tmp/artifacts/src/$abi_name.sol/$abi_name.json | jq '.abi' > $abi
  pkg=`echo $abi_name | tr 'A-Z' 'a-z'`
  out="internal/logic/contracts/abi/${pkg}/${pkg}.go"
  echo $out
  echo "generating ${out} from ${abi}"
  mkdir -p internal/logic/contracts/abi/$pkg
  abigen --abi=$abi --pkg=$pkg --out=$out
  awk '{sub("github.com/ethereum","github.com/scroll-tech")}1' $out > temp && mv temp $out
done

rm -rf tmp
rm -rf contracts_tmp