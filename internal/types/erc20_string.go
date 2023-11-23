// Code generated by "stringer -type ERC20"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERC20Unknown-0]
	_ = x[WETH-1]
	_ = x[StandardERC20-2]
	_ = x[CustomERC20-3]
	_ = x[USDC-4]
	_ = x[DAI-5]
	_ = x[LIDO-6]
}

const _ERC20_name = "ERC20UnknownWETHStandardERC20CustomERC20USDCDAILIDO"

var _ERC20_index = [...]uint8{0, 12, 16, 29, 40, 44, 47, 51}

func (i ERC20) String() string {
	if i < 0 || i >= ERC20(len(_ERC20_index)-1) {
		return "ERC20(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ERC20_name[_ERC20_index[i]:_ERC20_index[i+1]]
}
