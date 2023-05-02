package loan

import (
	"encoding/json"
	"log"
	"net/http"
)

type InitResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

func LoanInitHandler(w http.ResponseWriter, r *http.Request) {
	// Checks connectivity with accounting providers and decision engine.

	response := InitResponse{Error: false, Msg: "No error"}

	responseJson, err := json.Marshal(response)

	if err != nil {
		log.Fatal(err)
	}

	w.Write(responseJson)
}
