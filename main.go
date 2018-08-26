package main

import (
	"github.com/gorilla/mux"
	"goWoemaBackend/receiptFunctions"
	"log"
	"net/http"
)

type Recipt struct {

	transactionID string

}

func main() {

	r := receiptFunctions.NewRESTReceiveImageService()

	router := mux.NewRouter()

	router.HandleFunc("/", r.Receive).Methods("GET")
	router.HandleFunc("/uploadReceipt", r.Receive).Methods("POST")
	//router.HandleFunc("/getReceipt/{id}", r.GetReceipt).Methods("GET")

	log.Fatal(http.ListenAndServe(":9002",router))
}
