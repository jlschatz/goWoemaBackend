package uploadToS3

import (
	"fmt"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

}