package main

import (
	"github.com/jlschatz/goWoemaBackend/receiptFunctions"
	"log"
	"net/http"
	"time"
	"flag"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"context"
	"github.com/jlschatz/goWoemaBackend/s3bucket"
)

func main() {

	s3 := s3bucket.NewRESTUploadS3Service()
	r := receiptFunctions.NewRESTReceiveImageService(s3)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/getreceipt/{id}", r.GetReceipt).Methods("GET")
	router.HandleFunc("/upload/{id}", r.Upload).Methods("POST")


	srv := &http.Server{
		Addr:         ":9004",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	log.Println("HTTP server up. Listening on port 9004")
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)


}
