package local

import (
	"context"
	"os"

	"github.com/pkg/errors"
)

func (l *Local) DeleteFile(ctx context.Context, avatar string) error {
	if avatar == "" {
		return nil
	}

	err := os.Remove(avatar)
	if err != nil {
		return errors.Wrap(err, "Не удалось удалить файл")
	}

	return nil
}
