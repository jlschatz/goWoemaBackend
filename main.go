package main

import (
	"github.com/gorilla/mux"
	"github.com/jlschatz/goWoemaBackend/receiptFunctions"
	"log"
	"net/http"
)

func main() {

	r := receiptFunctions.NewRESTReceiveImageService()

	router := mux.NewRouter()

	router.HandleFunc("/", r.Receive).Methods("GET")
	router.HandleFunc("/uploadReceipt", r.Receive).Methods("POST")
	//router.HandleFunc("/getReceipt/{id}", r.GetReceipt).Methods("GET")

	log.Fatal(http.ListenAndServe(":9002",router))
}
