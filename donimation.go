package ethutils

import (
	"math/big"

	gparams "github.com/ethereum/go-ethereum/params"
)

// These are the multipliers for ether denominations.
var (
	Wei   = big.NewInt(gparams.Wei)
	GWei  = big.NewInt(gparams.GWei)
	Ether = big.NewInt(gparams.Ether)
)

// Normal used numbers in ethereum dapp.
var (
	One = big.NewInt(1)
	Two = big.NewInt(2)
)
