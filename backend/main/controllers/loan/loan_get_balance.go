package loan

import (
	"encoding/json"
	"main/main/controllers/thirdparty/accountingprovider"
	"net/http"
)

type BalanceResponse struct {
	Error    bool                         `json:"error"`
	ErrorMsg string                       `json:"errorMsg"`
	Balance  []accountingprovider.Balance `json:"balance"`
}

func LoanGetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var response BalanceResponse

	var query = r.URL.Query()
	var businessName = query.Get("business_name")
	var accountingProvider = query.Get("accounting_provider")

	// Retrieve balance sheet from accounting provider.
	err = accountingprovider.GetBalanceSheet(businessName, accountingProvider, &response.Balance)

	if err != nil {
		response.Error = true
		response.ErrorMsg = err.Error()
	}

	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)
}
