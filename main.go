package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//uploadS3 :=

	//mux := http.NewServeMux()

	//http.Handle("/", accessControl(mux))
	//mux.HandleFunc("/upload",)

	errs := make(chan error, 2)
	go func() {
		log.Println("transport", "http", "address", ":8002", "msg", "listening")
		errs <- http.ListenAndServe(":8002", nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Println("terminated", <-errs)
}

//func accessControl(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
//
//		if r.Method == "OPTIONS" {
//			return
//		}
//
//		h.ServeHTTP(w, r)
//	})
//}