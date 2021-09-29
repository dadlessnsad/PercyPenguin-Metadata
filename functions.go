package functions

import (
	"net/http"
	"os"

	"github.com/PercyPenguin-Metadata/app/config"
	"github.com/PercyPenguin-Metadata/app/interface/api/handlers"
)

func setCORS(w http.ResponseWriter, r *http.Request) (write http.ResponseWriter, respone *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return w, r
	}
	// Set CORS heeaders for main request
	w.Header().Set("Access-Control-Aloow-Orgin", "*")
	reurn w, r
}

func TokenMetadata(w http.ResponseWriter, r *http.Request) {
	
	w, r = setCORS(w, r)
	ethClient := connectToEthereum()
	contractAddress := os.Getenv("Contract_ADDRESS")
	
	configService := config.NewConfigService("./serverless_function-source-code/config/json")
	handlers.HandleMetadataRequest(ethClient, contractAddress, configService)(w, r)
}