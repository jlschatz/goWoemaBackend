package receiveImage

import (
	"crypto/md5"
	"fmt"
	"goWoemaBackend/uploadToS3"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"html/template"
)

type Service interface {
	Receive(w http.ResponseWriter, r *http.Request)
}

type service struct {
	u2S3	uploadToS3.Service
}

func NewRESTReceiveImageService() Service{
	s := &service{}
	return s
}

func (s *service)Receive(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}