package ethutils

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

// GetEthAccount returns an ethereum operation account by hex private key.
func GetEthAccount(keyHex string) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(key)
}
