package types

import (
	"fmt"
	"testing"

	"github.com/andreistefanciprian/cryptoblockutxo/crypto"
	"github.com/andreistefanciprian/cryptoblockutxo/proto"
	"github.com/andreistefanciprian/cryptoblockutxo/util"
	"github.com/stretchr/testify/assert"
)

// TestNewTransaction tests the creation of a new transaction
// Send 5 units from one address to another, with 100 units available
// Expect two outputs: 5 to recipient, 95 back to sender as change
func TestNewTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivKey.Public().Address().Bytes()
	toPrivKey := crypto.GeneratePrivateKey()
	toAddress := toPrivKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.Public().Bytes(),
	}
	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}
	tx := &proto.Transaction{
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(fromPrivKey, tx)
	input.Signature = sig.Bytes()

	fmt.Printf("Transaction: %+v\n", tx)

	assert.True(t, VerifyTransaction(tx), "Transaction should be valid")
}

func TestSignTransaction(t *testing.T) {

}
