package ethutils

import "testing"

func TestMustGenerateKey(t *testing.T) {
	key := MustGenrateKey()
	hexKey := PrivateKeyToHex(key)
	MustGetEthAccount(hexKey)
}
