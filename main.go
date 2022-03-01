package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/quiz_pbp_1/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/wallet/{wallet_id}", controllers.GetAllWalletTransaction).Methods("GET")
	router.HandleFunc("/transaction", controllers.InsertTransaction).Methods("POST")
	router.HandleFunc("/wallet/{wallet_id}", controllers.UpdateWallet).Methods("PUT")
	router.HandleFunc("/wallet/{wallet_id}", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/login/{login_id}", controllers.Login).Methods("PUT")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
