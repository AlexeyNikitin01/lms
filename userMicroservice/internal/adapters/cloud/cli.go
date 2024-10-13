package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"lms-user/cmd/config"
)

type AWS struct {
	S3     *session.Session
	bucket string
}

func NewAWS(cnf *config.AWS) (*AWS, error) {
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
	}, nil
}
