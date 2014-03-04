package main

import "math/big"

func NewBig(val int) *big.Int {
	return big.NewInt(int64(val))
}

func bigExp(base, exp int) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(int64(base)), big.NewInt(int64(exp)), big.NewInt(0))
}
