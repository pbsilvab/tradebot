package tradebot

import (
	"fmt"
)

func PlaceTestOrder() {
	params := map[string][]string{
		"symbol":      {"LTCUSDT"},
		"side":        {"BUY"},
		"type":        {"LIMIT"},
		"timeInForce": {"GTC"},
		"quantity":    {"1"},
		"price":       {"150"},
		"recvWindow":  {"5000"},
	}

	response := CallApi("/api/v3/order/test", "POST", params)
	fmt.Println(response)
}

func PlaceOrder() {
	params := map[string][]string{
		"symbol":      {"ZILUSDT"},
		"side":        {"BUY"},
		"type":        {"LIMIT"},
		"timeInForce": {"GTC"},
		"quantity":    {"240"},
		"price":       {"0.07"},
		"recvWindow":  {"5000"},
	}

	response := CallApi("/api/v3/order", "POST", params)
	fmt.Println(response)
}

func CancelOrder() {
	params := map[string][]string{
		"symbol":  {"ZILUSDT"},
		"orderId": {"590032945"},
	}

	response := CallApi("/api/v3/order", "DELETE", params)
	fmt.Println(response)
}

func CancellAllOpenOrders() {
	params := map[string][]string{
		"symbol": {"ZILUSDT"},
	}

	response := CallApi("/api/v3/openOrders", "DELETE", params)
	fmt.Println(response)

}

// Get Orders by Symbol
func CheckOrderStatus() {
	params := map[string][]string{
		"symbol":  {"ZILUSDT"},
		"orderId": {"590032945"},
	}

	response := CallApi("/api/v3/order", "GET", params)
	fmt.Println(response)
}

// Get open orders
func GetOpenOrders() {
	params := map[string][]string{
		"symbol": {"ZILUSDT"},
	}

	response := CallApi("/api/v3/openOrders", "GET", params)
	fmt.Println(response)
}

// get all orders
func GetAllOrders() {
	params := map[string][]string{
		"symbol": {"ZILUSDT"},
	}

	response := CallApi("/api/v3/openOrders", "GET", params)
	fmt.Println(response)
}

// account information
func GetAccountInfo() {
	params := map[string][]string{}

	response := CallApi("/api/v3/account", "GET", params)
	fmt.Println(response)
}

// Get trades for a specific account and symbol.
func GetAccountTradeList() {
	params := map[string][]string{
		"symbol": {"ZILUSDT"},
	}

	response := CallApi("/api/v3/myTrades", "GET", params)
	fmt.Println(response)
}

//Displays the user's current order count usage for all intervals.
func GetCurrentOrderCountUsage() {
	params := map[string][]string{}

	response := CallApi("/api/v3/rateLimit/order", "GET", params)
	fmt.Println(response)
}
