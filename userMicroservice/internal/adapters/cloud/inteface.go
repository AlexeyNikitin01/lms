package cloud

import (
	"context"
	"os"
)

type ICloud interface {
	Upload(ctx context.Context, file *os.File, name string, info os.FileInfo, mime string) error
	Link(fileName string) string
	AWSIsActive() bool
}
