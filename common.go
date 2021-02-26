package ethutils

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

// common params
var (
	CallOpt = &bind.CallOpts{}
)

// PanicIfErr panics if error occurs.
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
