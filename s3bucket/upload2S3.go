package s3bucket

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

type Service interface {
	Upload2S3(fileToUpload string)
}

type service struct {
}

func NewS3UploadService() Service{
	s := &service{}
	return s
}

func (s * service)Upload2S3(fileToUpload string) {

	aws_access_key_id := ""
	aws_secret_access_key := ""
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		// handle error
	}
	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(fileToUpload)
	if err != nil {
		// handle error
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/receiptImages/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String("receipt-storage-101"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		// handle error
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
}