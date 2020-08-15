package api

import (
	"API-example/block"
	"API-example/blockchain"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

var Blockchain *blockchain.Blockchain

type Message struct {
	Data string
}

func GetBlockchain(echoContext echo.Context) error {

	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, "No blockchain found on api.")
	}
	return echoContext.String(http.StatusOK, string(bytes))
}

//
func AddBlock(echoContext echo.Context) error {
	var message Message
	if err := echoContext.Bind(&message); err != nil {
		return echoContext.NoContent(http.StatusBadGateway)
	}
	Blockchain.AddBlock(Blockchain.GenerateNewBlock(message.Data))
	bytes, _ := json.MarshalIndent(&Blockchain, "", "  ")
	return echoContext.String(http.StatusOK, string(bytes))
}

func CreateBlockchain(echoContext echo.Context) error {
	var message Message
	if err := echoContext.Bind(&message); err != nil {
		return echoContext.NoContent(http.StatusBadGateway)
	}
	Blockchain = blockchain.NewBlockChain()
	Genesis := block.GenesisBlock(message.Data)
	Blockchain.AddBlock(Genesis)
	bytes, _ := json.MarshalIndent(&Blockchain, "", "  ")
	return echoContext.String(http.StatusOK, string(bytes))
}
