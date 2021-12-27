package tradebot

import (
	"io/ioutil"
	"log"
	"net/http"
)

const endpoint string = "https://api.binance.com"

func getSymbolStatusEnums() []string {
	return []string{
		"PRE_TRADING",
		"TRADING",
		"POST_TRADING",
		"END_OF_DAY",
		"HALT",
		"AUCTION_MATCH",
		"BREAK",
	}
}

func getOrderStatusEnums() []string {
	return []string{
		"NEW",
		"PARTIALLY_FILLED",
		"FILLED",
		"CANCELED",
		"PENDING_CANCEL",
		"REJECTED",
		"EXPIRED",
	}
}

func getOrderTypesEnums() []string {
	return []string{
		"LIMIT",
		"MARKET",
		"STOP_LOSS",
		"STOP_LOSS_LIMIT",
		"TAKE_PROFIT",
		"TAKE_PROFIT_LIMIT",
		"LIMIT_MAKER",
	}
}

type requestStruct struct {
	endpoint string
	method   string
	params   []string
}

func PrepareRequest(endpoint string, method string, params []string) requestStruct {
	newReq := requestStruct{
		endpoint: endpoint,
		method:   method,
		params:   params,
	}

	return newReq
}

func call(request requestStruct) string {
	url := endpoint
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	b := string(body)

	return b
}
