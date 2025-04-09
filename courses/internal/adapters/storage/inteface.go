package storage

import (
	"context"
	"mime/multipart"
)

type ICloud interface {
	Upload(ctx context.Context, fileForm multipart.File, header *multipart.FileHeader) (string, error)
	DeleteFile(ctx context.Context, key string) error
}
