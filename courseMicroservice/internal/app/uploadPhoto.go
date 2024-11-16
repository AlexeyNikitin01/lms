package app

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
)

func (c CourseApp) UploadPhoto(ctx context.Context, photo []byte, filename string, mime string) (string, error) {
	path, err := saveFileLocally(filename, photo)
	if err != nil {
		return "", fmt.Errorf("failed to save file locally: %v", err)
	}

	defer func() {
		if err := os.Remove(path); err != nil {
			log.Printf("Не удалось удалить временный файл: %v", err)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "не удалось открыть файл: %v")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", errors.Wrap(err, "не удалось получить информацию о файле: %v")
	}

	if err := c.S3.Upload(ctx, file, filename, fileInfo, mime); err != nil {
		return "", errors.Wrap(err, "error upload photo")
	}

	url := c.S3.Link(filename)

	return url, nil
}

func saveFileLocally(filename string, data []byte) (string, string, error) {
	uploadPath := "./tmp"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", "", errors.Wrap(err, "Не удалось создать папку для загрузок")
	}

	uniqueFileName := generateFilename(filename)

	filePath := filepath.Join(uploadPath, uniqueFileName)

	out, err := os.Create(filePath)
	if err != nil {
		return "", "", errors.Wrap(err, "Не удалось создать файл на сервере")
	}

	defer out.Close()

	return filePath, uniqueFileName, nil
}


func generateFilename(original string) string {
	ext := filepath.Ext(original)
	name := filepath.Base(original[:len(original)-len(ext)])
	timestamp := time.Now().UnixNano()

	return fmt.Sprintf("%s_%d%s", name, timestamp, ext)
}
