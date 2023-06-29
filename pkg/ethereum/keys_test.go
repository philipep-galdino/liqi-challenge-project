package ethereum

import (
	"strings"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	privateKey, publicKey, err := GenerateKeys()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
		return
	}

	if len(privateKey) == 0 {
		t.Errorf("Expected private key, got an empty string")
	}

	if len(publicKey) == 0 {
		t.Errorf("Expected public key, got an empty string")
	}

	if !strings.HasPrefix(publicKey, "04") {
		t.Errorf("Public key should start with '04'")
	}
}

func TestGetAddress(t *testing.T) {
	_, publicKey, _ := GenerateKeys()
	address, err := GetAddress(publicKey)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(address) == 0 {
		t.Errorf("Expected Ethereum address, got an empty string")
	}
}
