#!/bin/bash
#set -xe

tmp_file="/tmp/contracts_json"
rm -rf "${tmp_file}" && mkdir "${tmp_file}"

# L1
l1=("L1ScrollMessenger")

# L1/gateway
l1_gateway=("L1ETHGateway" "L1ERC20Gateway" "L1DAIGateway" "L1WETHGateway" "L1StandardERC20Gateway" "L1CustomERC20Gateway" "L1ERC721Gateway" "L1ERC1155Gateway")

# L1/rollup
l1_rollup=("ScrollChain" "L1MessageQueue")

# L2
l2=("L2ScrollMessenger")

# L2/gateway
l2_gateway=("L2ETHGateway" "L2ERC20Gateway" "L2DAIGateway" "L2WETHGateway" "L2StandardERC20Gateway" "L2CustomERC20Gateway" "L2ERC721Gateway" "L2ERC1155Gateway")

# L2/predeploys
l2_predeploys=("L1BlockContainer" "L1GasPriceOracle" "L2MessageQueue")

# token list
token_list=("IERC20" "IERC721" "IERC1155")

# The first parameter is scroll path(github.com:scroll-tech/scroll).
scroll_path=$1

extract_abi() {
  local services=("$@")
  for i in "${!services[@]}"; do
    abi="${tmp_file}"/${services[$i]}.json
    # jq '[ .metadata.output.abi | .[] | select(.name != "OwnershipTransferred" and .name != "UpdateWhitelist") ]'
    cat "${scroll_path}"/contracts/artifacts/src/"${services[$i]}".sol/"${services[$i]}".json | jq '.metadata.output.abi' >"${abi}"
    # shellcheck disable=SC2001
    contract=$(echo "${abi}" | sed 's#.*/##; s/\..*//')
    # shellcheck disable=SC2001
    pkg=$(echo "$dest" | sed 's#.*/##; s/\..*//')
    go run github.com/scroll-tech/go-ethereum/cmd/abigen --tmpl "./metrics.tmpl" --abi "${tmp_file}/${contract}.json" --pkg "${pkg}" --type "${contract}" --out scroll/"${dest}/${contract}.go"
  done
}

dest=$2
# shellcheck disable=SC2006
while [ -n "$dest" ]; do
  case "$dest" in
  L1)
    extract_abi "${l1[@]}"
    break
    ;;
  L1/gateway)
    extract_abi "${l1_gateway[@]}"
    break
    ;;
  L1/rollup)
    extract_abi "${l1_rollup[@]}"
    break
    ;;
  L2)
    extract_abi "${l2[@]}"
    break
    ;;
  L2/gateway)
    extract_abi "${l2_gateway[@]}"
    break
    ;;
  L2/predeploys)
    extract_abi "${l2_predeploys[@]}"
    break
    ;;
  token)
    extract_abi "${token_list[@]}"
    break
    ;;
  *)
    echo "$1 is not supported"
    exit 1
    ;;
  esac
  shift
done
