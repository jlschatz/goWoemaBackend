package uploadToS3

import "log"

type Service interface {
	testPrint()
}

type service struct {
}

func NewUploadS3Service() Service{
	s := &service{}
	return s
}

func (s *service)testPrint() {
	log.Println("HTTP request received")
}