package ethutils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
)

// MustGenrateKey generates an random key using ethereum/crypto.
func MustGenrateKey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	log.Info("generate success", "privkey", PrivateKeyToHex(key), "pub addr", crypto.PubkeyToAddress(key.PublicKey))

	return key
}

// PrivateKeyToHex change to private key to hex format.
func PrivateKeyToHex(key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(key))
}

// MustGetEthAccount returns an ethereum operation account by hex private key.
func MustGetEthAccount(keyHex string) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(key)
}

// MustGetETHAccountByEncryptFile returns an ethereum account by keystore file path path.
func MustGetETHAccountByEncryptFile(keypath string, pw string) *bind.TransactOpts {
	b, err := ioutil.ReadFile(keypath)
	if err != nil {
		panic(err)
	}
	k, err := keystore.DecryptKey(b, pw)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(k.PrivateKey)
}

// MustGetDefaultOptByEncryptFile returns an ethereum account with nonce set by keystore file path path.
func MustGetDefaultOptByEncryptFile(keypath string, pw string, client *ethclient.Client) *bind.TransactOpts {
	opt := MustGetETHAccountByEncryptFile(keypath, pw)
	MustSetNonce(client, opt)

	return opt
}

// MustGetDefaultOptByHexKey returns an ethereum account with nonce set by hex key.
func MustGetDefaultOptByHexKey(keyHex string, client *ethclient.Client) *bind.TransactOpts {
	opt := MustGetEthAccount(keyHex)
	MustSetNonce(client, opt)

	return opt
}
