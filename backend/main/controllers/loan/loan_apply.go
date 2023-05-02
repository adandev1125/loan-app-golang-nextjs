package loan

import (
	"encoding/json"
	"fmt"
	"main/main/controllers/thirdparty/accountingprovider"
	"main/main/controllers/thirdparty/decisionengine"
	"net/http"
)

type ApplyPostData struct {
	Amount             int    `json:"amount"`
	BusinessName       string `json:"businessName"`
	EstablishedYear    int    `json:"businessYear"`
	AccountingProvider string `json:"AccountingProvider"`
}

type ApplyResponse struct {
	Error    bool    `json:"error"`
	ErrorMsg string  `json:"errorMsg"`
	Amount   float32 `json:"amount"`
}

func LoanApplyHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var applyPostData ApplyPostData
	var applyResponse ApplyResponse

	// Parse the post json data.
	err = json.NewDecoder(r.Body).Decode(&applyPostData)
	fmt.Printf("%#v", applyPostData)

	if err != nil {
		applyResponse.Error = true
		applyResponse.ErrorMsg = err.Error()
		responseJson, _ := json.Marshal(applyResponse)
		w.Write(responseJson)
		return
	}

	// Retrieve balance sheet of the business from accounting provider.
	var balance = make([]accountingprovider.Balance, 0)
	err = accountingprovider.GetBalanceSheet(applyPostData.BusinessName, applyPostData.AccountingProvider, &balance)

	if err != nil {
		applyResponse.Error = true
		applyResponse.ErrorMsg = err.Error()
		responseJson, _ := json.Marshal(applyResponse)
		w.Write(responseJson)
		return
	}

	// Calculate preAssessment value and profit by the year.
	var avgAssetsValue = 0.0
	var sumProfit = 0.0
	var balanceLength = len(balance)
	var profitByYear = make(map[int]int, 10)

	for _, v := range balance {
		sumProfit += float64(v.ProfitOrLoss)
		avgAssetsValue += float64(v.AssetsValue) / float64(balanceLength)

		saved, ok := profitByYear[v.Year]
		if !ok {
			profitByYear[v.Year] = v.ProfitOrLoss
		} else {
			profitByYear[v.Year] = saved + v.ProfitOrLoss
		}
	}

	var preAssessment = 20
	if sumProfit > 0 {
		preAssessment = 60
	}
	if avgAssetsValue > float64(applyPostData.Amount) {
		preAssessment = 100
	}

	// Retrieves decision from decision engine.
	var decided = false
	decided = decisionengine.GetDecision(applyPostData.BusinessName, applyPostData.EstablishedYear, preAssessment, profitByYear)

	if !decided {
		applyResponse.Error = true
		applyResponse.ErrorMsg = "Sorry, but our decision engine rejected to loan to your business."
		responseJson, _ := json.Marshal(applyResponse)
		w.Write(responseJson)
		return
	}

	applyResponse.Error = false
	applyResponse.Amount = float32(preAssessment) * float32(applyPostData.Amount) / 100

	responseJson, _ := json.Marshal(applyResponse)
	w.Write(responseJson)
}
