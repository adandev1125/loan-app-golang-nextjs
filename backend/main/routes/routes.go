package routes

import (
	"main/main/controllers"
	"main/main/controllers/loan"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var router *mux.Router
var once sync.Once

func GetRouter() *mux.Router {

	once.Do(func() {

		router = mux.NewRouter()
		initRouter()
	})

	return router
}

func initRouter() {

	// Add your routes here.

	router.HandleFunc("/", controllers.WelcomeHandler).Methods("GET")

	router.HandleFunc("/loan/init", loan.LoanInitHandler).Methods("GET")

	router.HandleFunc("/loan/apply_loan", loan.LoanApplyHandler).Methods("POST")

	router.HandleFunc("/loan/get_balance", loan.LoanGetBalanceHandler).Methods("GET")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
}
