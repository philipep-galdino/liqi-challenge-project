package ethereum

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
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
