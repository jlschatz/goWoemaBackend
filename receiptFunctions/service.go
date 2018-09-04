package receiptFunctions

import (
	"fmt"
	"image"
	"io"
	"net/http"
	"os"
	"github.com/jlschatz/goWoemaBackend/s3bucket"
	"github.com/gorilla/mux"
	"log"
)

type transactions struct {
	trans []transaction
}

type transaction struct {
	id  string
	img image.Image
}

type Service interface {
	Upload(w http.ResponseWriter, r *http.Request)
	GetReceipt(w http.ResponseWriter, r *http.Request)
}

type service struct {
	u2S3 s3bucket.Service
}

func NewRESTReceiveImageService(s3 s3bucket.Service) Service {
	s := &service{u2S3: s3}
	return s
}

func (s *service) Upload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionID := params["id"]

	log.Println(transactionID)

	fmt.Println("method:", r.Method)

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("chuck.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer os.Remove(f.Name())
	io.Copy(f, file)
	os.Rename(f.Name(), transactionID + ".jpg")
	s.u2S3.Upload2S3(transactionID + ".jpg")
}


func (s *service)GetReceipt(w http.ResponseWriter, r *http.Request) {

}
