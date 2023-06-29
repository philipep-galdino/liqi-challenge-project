package ethereum

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestConnectToEthNetwork(t *testing.T) {
	client, err := ConnectToEthNetwork()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Error("Expected client to be non-nil, got nil")
	}
}

func TestSignTransaction(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()
	to := "0xAbC123..."
	value := "0x1bc16d674ec80000"
	nonce := uint64(0)

	signedTx, err := SignTransaction(to, value, nonce, privateKey)

	if err != nil {
		t.Errorf("Got error %v", err)
	}

	if signedTx == nil {
		t.Error("SignedTX is nil")
	}
}
