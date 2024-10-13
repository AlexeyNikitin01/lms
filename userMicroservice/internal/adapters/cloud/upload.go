package cloud

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

func (a AWS) Upload(ctx context.Context, file *os.File, name string, info os.FileInfo, mime string) error {
	_, err := s3manager.NewUploader(a.S3).UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:             aws.String(a.bucket),
		Key:                aws.String(name),
		Body:               file,
		ContentType:        aws.String(mime),
		ContentDisposition: aws.String(fmt.Sprintf(`filename="%s"`, url.QueryEscape(info.Name()))),
	})
	if err != nil {
		return errors.Wrap(err, "cloud upload err")
	}

	return nil
}
