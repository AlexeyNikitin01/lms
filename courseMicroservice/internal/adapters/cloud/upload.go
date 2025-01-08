package cloud

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

func (a AWS) Upload(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
) (string, error) {
	uniqueFileName := generateFilename(header.Filename)

	_, err := s3manager.NewUploader(a.S3).UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:             aws.String(a.bucket),
		Key:                aws.String(uniqueFileName),
		Body:               fileForm,
		ContentType:        aws.String("image/jpeg"),
		ContentDisposition: aws.String(fmt.Sprintf(`filename="%s"`, url.QueryEscape(header.Filename))),
	})
	if err != nil {
		return "", errors.Wrap(err, "cloud upload err")
	}

	return uniqueFileName, nil
}

func generateFilename(original string) string {
	ext := filepath.Ext(original)
	name := filepath.Base(original[:len(original)-len(ext)])
	timestamp := time.Now().UnixNano()

	return fmt.Sprintf("%s_%d%s", name, timestamp, ext)
}
