package types

import (
	"crypto/sha256"

	"github.com/andreistefanciprian/cryptoblockutxo/crypto"
	"github.com/andreistefanciprian/cryptoblockutxo/proto"
	pb "google.golang.org/protobuf/proto"
)

// HashBlock returns the SHA-256 hash of the header.
func HashBlock(block *proto.Block) []byte {
	data, err := pb.Marshal(block)
	if err != nil {
		return nil
	}
	hash := sha256.Sum256(data)
	return hash[:]
}

func SignBlock(block *proto.Block, privKey *crypto.PrivateKey) *crypto.Signature {
	return privKey.Sign(HashBlock(block))
}
