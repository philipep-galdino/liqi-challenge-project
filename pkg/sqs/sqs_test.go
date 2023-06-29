package sqs

import (
	"encoding/json"
	"testing"
)

func TestTransactionJSON(t *testing.T) {
	jsonMessage := `{"to":"0xAbC123...", "value":"0x1bc16d674ec80000"}`
	var trans Transaction

	err := json.Unmarshal([]byte(jsonMessage), &trans)

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if trans.To != "0xAbC123..." {
		t.Errorf("Expected '0xAbC123...', got %v", trans.To)
	}

	if trans.Value != "0x1bc16d674ec80000" {
		t.Errorf("Expected '0x1bc16d674ec80000', got %v", trans.Value)
	}

}
