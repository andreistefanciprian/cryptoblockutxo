package types

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/andreistefanciprian/cryptoblockutxo/crypto"
	"github.com/andreistefanciprian/cryptoblockutxo/util"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	fmt.Println("Head of the Block: ", hex.EncodeToString(hash))

	assert.Equal(t, 32, len(hash), "Hash length should be 32 bytes")
}

func TestSignBlock(t *testing.T) {
	var (
		block     = util.RandomBlock()
		privKey   = crypto.GeneratePrivateKey()
		signature = SignBlock(block, privKey)
		pubKey    = privKey.Public()
	)

	assert.Equal(t, 64, len(signature.Bytes()), "Signature length should be 64 bytes")
	assert.True(t, signature.Verify(pubKey, HashBlock(block)), "Signature should be valid")
}
