package functions

import (
	"os"

	"github.com/PercyPenguin-Metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
)

func connectToEthereum() *ethereum.EthereumClient {

	nodeURL := os.Getenv("NODE_URL")

	client, err := ethereum.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Successfully connect to ethereum client")

	return client
}
