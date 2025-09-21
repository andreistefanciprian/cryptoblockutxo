package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	// Generate a new private key
	privKey := GeneratePrivateKey()

	// Ensure the private key is of the expected length
	assert.Equal(t, privKeyLen, len(privKey.Bytes()), "Private key length should be 64 bytes")

	pubKey := privKey.Public()
	// Ensure the public key is of the expected length
	assert.Equal(t, pubKeyLen, len(pubKey.Bytes()), "Public key length should be 32 bytes")
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "0f1ee7359069ea5c7ecbf23415289ac698ceee0f7779f0df564ef345cee4f0f0"
		addressStr = "9d04904b58abd5934683298541463b4f866edd20"
	)

	privKey := NewPrivateKeyFromString(seed)
	address := privKey.Public().Address()

	assert.Equal(t, privKeyLen, len(privKey.Bytes()), "Private key length should be 64 bytes")
	assert.Equal(t, addressStr, address.String(), "Addresses should match")
}

// func TestNewPrivateKeyFromSeed(t *testing.T) {
// 	seed := "0f1ee7359069ea5c7ecbf23415289ac698ceee0f7779f0df564ef345cee4f0f0"
// 	privKey := NewPrivateKeyFromSeed([]byte(seed))
// 	assert.Equal(t, seed, fmt.Sprintf("%x", privKey.Bytes()), "Private keys should match")
// }

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("test message")

	sig := privKey.Sign(msg)
	// // Ensure the signature is of the expected length
	// assert.Equal(t, 64, len(sig.value), "Signature length should be 64 bytes")
	// Verify the signature
	valid := sig.Verify(pubKey, msg)
	assert.True(t, valid, "Signature should be valid")

	// Test with a modified message
	invalid := sig.Verify(pubKey, []byte("modified message"))
	assert.False(t, invalid, "Signature should be invalid for modified message")

	// Test with different pub key
	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.Public()
	assert.False(t, sig.Verify(otherPubKey, msg), "Signature should be invalid for different public key")

}

func TestPublickKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	// Ensure the address is of the expected length
	fmt.Println("Address string:", address.String())
	assert.Equal(t, addressLen, len(address.Bytes()), "Address length should be 20 bytes")
}
