package receiptFunctions

import (
	"fmt"
	"github.com/jlschatz/goWoemaBackend/uploadToS3"
	"image"
	"io"
	"net/http"
	"os"
)

type transactions struct {
	trans []transaction
}

type transaction struct {
	id  string
	img image.Image
}

type Service interface {
	Receive(w http.ResponseWriter, r *http.Request)
	//	GetSpecificReceipt(w http.ResponseWriter, r *http.Request)
}

type service struct {
	u2S3 uploadToS3.Service
}

func NewRESTReceiveImageService() Service {
	s := &service{}
	return s
}

func (s *service) Receive(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
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
	os.Rename(f.Name(), "poop.jpg")

}

//
//func (s *service)GetSpecificReceipt(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//
//
//
//	for _, trans := range t {
//		if trans.
//	}
//}
