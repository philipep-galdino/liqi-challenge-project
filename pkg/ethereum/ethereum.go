package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func ConnectToEthNetwork() (*ethclient.Client, error) {
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")
	infuraProjectSecret := os.Getenv("INFURA_PROJECT_SECRET")

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/" + infuraProjectID + ":" + infuraProjectSecret)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
		return nil, err
	}

	return client, nil

}

func SignTransaction(to string, value string, nonce uint64, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	toAddress := common.HexToAddress(to)
	amount := new(big.Int)
	amount.SetString(value[2:], 16)

	gasLimit := uint64(21000)
	gasPrice := big.NewInt(20 * params.GWei)

	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)

	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
		return nil, err
	}

	return signedTx, nil
}

func SendTransaction(client *ethclient.Client, signedTx *types.Transaction) (string, error) {
	err := client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
		return "", err
	}

	hash := signedTx.Hash().Hex()

	return hash, nil
}

func GetNonce(client *ethclient.Client, address string) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(address))

	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return 0, err
	}

	return nonce, nil
}

func StringToPrivateKey(privateKeyString string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to convert private key to ECDSA pk, got error: %v", err)
		return nil, err
	}

	return privateKey, nil
}
