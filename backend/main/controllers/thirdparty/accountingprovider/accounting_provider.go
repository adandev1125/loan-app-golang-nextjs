package accountingprovider

import (
	"encoding/json"
)

type Balance struct {
	Year         int `json:"year"`
	Month        int `json:"month"`
	ProfitOrLoss int `json:"profitOrLoss"`
	AssetsValue  int `json:"assetsValue"`
}

const GET_BALANCE_API = "http://localhost:9090/get_balance"

func GetBalanceSheet(businessName string, accountingProvider string, balance *[]Balance) error {
	var err error

	// Sends a balance sheet request with those values above.
	// urlValues := url.Values{}
	// urlValues.Set("business_name", businessName)
	// urlValues.Set("accounting_provider", accountingProvider)

	// request, err := http.NewRequest("GET", config.GET_BALANCE_API+"?"+urlValues.Encode(), nil)

	// if err != nil {
	// 	return err
	// }

	// client := &http.Client{}
	// balanceRes, err := client.Do(request)

	// if err != nil {
	// 	return err
	// }

	// defer balanceRes.Body.Close()

	// balanceResBytes := make([]byte, 0)

	// _, err = balanceRes.Body.Read(balanceResBytes)

	// if err != nil {
	// 	return err
	// }

	jsonStr := `[
		{
			"year": 2023,
			"month": 3,
			"profitOrLoss": 250000,
			"assetsValue": 1234
		},
		{
			"year": 2023,
			"month": 2,
			"profitOrLoss": 1150,
			"assetsValue": 5789
		},
		{
			"year": 2023,
			"month": 1,
			"profitOrLoss": 2500,
			"assetsValue": 22345
		},
		{
			"year": 2022,
			"month": 12,
			"profitOrLoss": -187000,
			"assetsValue": 223452
		}
	]`

	balanceResBytes := []byte(jsonStr)

	err = json.Unmarshal(balanceResBytes, balance)

	return err
}
