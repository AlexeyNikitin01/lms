package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"lms-user/cmd/config"
	"lms-user/internal/adapters/storage"
)

const URL = "https://storage.yandexcloud.net/lms-user/"

type AWS struct {
	S3     *session.Session
	bucket string
	URL    string
}

func NewAWS(cnf *config.AWS) (storage.IFace, error) {
	s3Session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			cnf.PublicKey,
			cnf.SecretKey,
			"",
		),
		Endpoint: aws.String(cnf.Endpoint),
		Region:   aws.String(cnf.Region),
	})
	if err != nil {
		return nil, err
	}

	return &AWS{
		S3:     s3Session,
		bucket: cnf.Bucket,
		URL:    URL,
	}, nil
}
