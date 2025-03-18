package cloud

import (
	"context"
	"mime/multipart"
)

type ICloud interface {
	Upload(ctx context.Context, fileForm multipart.File, header *multipart.FileHeader) (string, error)
	Link(fileName string) string
	DeleteFile(ctx context.Context, key string) error
}
