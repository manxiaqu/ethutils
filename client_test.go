package ethutils

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func testDeployFuncSuccess() (common.Address, *types.Transaction, interface{}, error) {
	return common.Address{}, &types.Transaction{}, nil, nil
}

func testDeployFuncFailed() (common.Address, *types.Transaction, interface{}, error) {
	return common.Address{}, nil, nil, errors.New("test deploy failed")
}

func TestAutoDeploy(t *testing.T) {
	auth := MustGetEthAccount(testHexKey)
	auth.Nonce = big.NewInt(1)

	// test for deploy success
	before := auth.Nonce.Uint64()
	AutoDeploy("", testDeployFuncSuccess, auth)
	assert.Equal(t, auth.Nonce.Uint64()-before, uint64(1))

	// test for deploy failed
	before = auth.Nonce.Uint64()
	assert.Panics(t, func() { AutoDeploy("", testDeployFuncFailed, auth) })
	assert.Equal(t, auth.Nonce.Uint64()-before, uint64(0))
}

func testAutoSendTxToContractSuccess() (*types.Transaction, error) {
	return &types.Transaction{}, nil
}
func testAutoSendTxToContractFailed() (*types.Transaction, error) {
	return &types.Transaction{}, errors.New("test AutoSendTxToContract Failed")
}

func TestAutoSendTxToContract(t *testing.T) {
	auth := MustGetEthAccount(testHexKey)
	auth.Nonce = big.NewInt(1)

	// test for send tx success
	before := auth.Nonce.Uint64()
	AutoSendTx("", testAutoSendTxToContractSuccess, auth)
	assert.Equal(t, auth.Nonce.Uint64()-before, uint64(1))

	// test for send tx failed
	before = auth.Nonce.Uint64()
	assert.Panics(t, func() { AutoSendTx("", testAutoSendTxToContractFailed, auth) })
	assert.Equal(t, auth.Nonce.Uint64()-before, uint64(0))
}
