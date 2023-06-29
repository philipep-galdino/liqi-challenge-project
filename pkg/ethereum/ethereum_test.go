package ethereum

import (
	"testing"
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
