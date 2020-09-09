package ethutils

import "testing"

func TestMustGenerateKey(t *testing.T) {
	key := MustGenrateKey()
	hexKey := PrivateKeyToHex(key)
	GetEthAccount(hexKey)
}
