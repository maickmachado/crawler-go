package routes

import (
	"log"
	"net/http"

	"github.com.br/maickmachado/crawler-go/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/trends/dailytrends", controllers.GetDailyTrends).Methods("GET")
	myRouter.HandleFunc("/trends/realtimetrends", controllers.GetRealTimeTrends).Methods("GET")
	myRouter.HandleFunc("/trends/cryptocoins/{name}", controllers.GetAvailableCategories).Methods("GET")
	myRouter.HandleFunc("/healthcheck", controllers.GetCompareKeywords).Methods("GET")
	// myRouter.HandleFunc("/cryptocoins/vote/{text}", controllers.VoteCrypto).Methods("POST")
	// myRouter.NotFoundHandler = http.Handler(http.HandlerFunc(controllers.ErrorHandler404))

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
