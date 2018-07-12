package main

import (
	"math/big"
)

func HexToBigInt(hex string) (*big.Int, error) {
	i := new(big.Int)
	n, _ := i.SetString(hex, 0)
	return n, nil
}
