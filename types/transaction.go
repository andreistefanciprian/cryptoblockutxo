package types

import (
	"crypto/sha256"

	"github.com/andreistefanciprian/cryptoblockutxo/crypto"
	"github.com/andreistefanciprian/cryptoblockutxo/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(privKey *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return privKey.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	data, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		sig := crypto.SignatureFromBytes(input.Signature)
		pubKey := crypto.PublicKeyFromBytes(input.PublicKey)
		input.Signature = nil // Clear signature for hashing
		// TODO: We need to put back the initial signature

		// Verify the signature
		if !sig.Verify(pubKey, HashTransaction(tx)) {
			return false
		}
	}
	return true
}
