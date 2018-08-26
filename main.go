package main

import (
	"github.com/gorilla/mux"
	"goWoemaBackend/receiveImage"
	"log"
	"net/http"
)

type Recipt struct {

	transactionID string

}

func main() {

	r := receiveImage.NewRESTReceiveImageService()

	router := mux.NewRouter()

	router.HandleFunc("/uploadReceipt", r.Receive).Methods("POST")
	//router.HandleFunc("/getReceipt/{id}", u.).Methods("GET")

	log.Fatal(http.ListenAndServe(":9002",router))
}
