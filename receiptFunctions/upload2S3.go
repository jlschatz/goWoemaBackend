package receiptFunctions

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Service interface {
	Upload(bucketName, fileToUpload string)
}

type service struct {
}

func NewRESTUploadS3Service() Service{
	s := &service{}
	return s
}

func (s * service)Upload(bucketName, fileToUpload string) {

	bucket := bucketName
	filename := fileToUpload

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	//select Region to use.
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}