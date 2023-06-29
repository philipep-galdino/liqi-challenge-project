package ethereum

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
