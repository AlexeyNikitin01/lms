package storage

import (
	"context"
	"mime/multipart"
)

type IFace interface {
	Upload(ctx context.Context, fileForm multipart.File, header *multipart.FileHeader) (string, error)
	DeleteFile(ctx context.Context, avatar string) error
}
