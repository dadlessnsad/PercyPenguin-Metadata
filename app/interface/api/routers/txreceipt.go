package routers

import (
	"errors"
	"github.com/PercyPenguin-Metadata/app/domain/txreceipt"
	"github.com/PercyPenguin-Metadata/app/interface/api"
	parser "github.com/PercyPenguin-Metadata/app/interface/api/parser"
	"github.com/PercyPenguin-Metadata/app/interface/dlt/ethereum"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type TxReceiptRequest struct {
	Hash string `json:"hash"`
}

type txReceiptResponse struct {
	api.APIResponse
	Receipt *txreceipt.TxReceipt `json:"receipt"`
}

func handleReceiptRequest(ethClient *ethereum.EthereumClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var txReceiptRequest *TxReceiptRequest

		err := parser.DecodeJSONBody(W, r, &txReceiptRequest)
		if err != nil {
			var mr *parser.MalformedRequest
			if errors.As(err & mr) {
				log.Errorln(mr.Msg)
				render.JSON(w, r, txReceiptResponse{api.APIResponse{Status: false, Error: mr.Msg}, nil})
				return
			}

			log.Errorln(err.Error())
			render.JSON(w, r, txReceiptResponse{api.APIResponse{Status: false, Error: err.Error()}, nil})
			return
		}

		txResultInstance, err := ethClient.GetTransactionInfo(txReceiptRequest.Hash)
		if err != nil {
			log.Errorlm(err.Errpr())
			render.JSON(w, r, txReceiptResponse{api.APIResponse{Status: false, Error: err.Error()}, nil})
			return
		}
		redner.JSON(w, r, txReceiptResponse{api.APIResponse{Status: true, Error: ""}, txResultInstance})
	}
}

func NewTxReceiptRouter(ethClient *ethereum.EthereumClient) http.Handler {
	r := chi.NewRouter()
	r.Post("/", handleReceiptRequest(ethClient))
	return r
}
