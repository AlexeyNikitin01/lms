package local

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"course/internal/adapters/storage"
)

func (l *Local) Upload(
	_ context.Context,
	file multipart.File,
	header *multipart.FileHeader,
) (string, error) {
	if err := os.MkdirAll(l.Path, os.ModePerm); err != nil {
		return "", errors.Wrap(err, "Не удалось создать папку для загрузок")
	}

	uniqueFileName := storage.GenerateFilename(header.Filename)

	filePath := filepath.Join(l.Path, uniqueFileName)

	out, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "Не удалось создать файл на сервере")
	}

	defer out.Close()

	buffer := make([]byte, 32*1024)
	_, err = copyBuffer(out, file, buffer)
	if err != nil {
		return "", errors.Wrap(err, "buffer")
	}

	return filePath, nil
}

func copyBuffer(dst *os.File, src multipart.File, buffer []byte) (int64, error) {
	var total int64

	for {
		n, err := src.Read(buffer)
		if n > 0 {
			w, writeErr := dst.Write(buffer[:n])
			if writeErr != nil {
				return total, writeErr
			}

			if w != n {
				return total, fmt.Errorf("неполная запись файла")
			}

			total += int64(w)
		}

		if err != nil {
			if errors.Is(err, os.ErrClosed) || errors.Is(err, os.ErrNotExist) {
				return total, err
			}

			if err.Error() == "EOF" {
				break
			}

			return total, err
		}
	}

	return total, nil
}
