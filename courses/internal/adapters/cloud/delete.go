package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func (a AWS) DeleteFile(ctx context.Context, key string) error {
	if key == "" {
		return nil
	}

	svc := s3.New(a.S3)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(a.bucket),
		Key:    aws.String(key),
	}

	_, err := svc.DeleteObjectWithContext(ctx, input)
	if err != nil {
		return errors.Wrapf(err, "failed to delete object with key %q from bucket %q", key, a.bucket)
	}

	return nil
}
