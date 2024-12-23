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

func (a appUser) UploadPhoto(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	userID string,
) (bool, string, error) {
	if a.s3.AWSIsActive() {
		url, err := a.uploadAvatarS3(ctx, fileForm, header, userID)
		return a.s3.AWSIsActive(), url, err
	}

	path, err := a.uploadAvatarLocal(ctx, fileForm, header, userID)

	return a.s3.AWSIsActive(), path, err
}

// UploadAvatarS3 TODO: удалить старое фото профиля.
func (a appUser) uploadAvatarS3(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	userID string,
) (string, error) {
	path, fileName, err := saveTmpFile(fileForm, header.Filename)
	if err != nil {
		return "", errors.Wrap(err, "tmp")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("не удалось получить информацию о файле: %v", err)
	}

	err = a.s3.Upload(ctx, file, fileName, fileInfo, "image/jpeg")
	if err != nil {
		return "", errors.Wrap(err, "s3")
	}

	if err := os.Remove(path); err != nil {
		log.Printf("Не удалось удалить временный файл: %v", err)
	}

	url := a.s3.Link(fileName)

	if err := a.repo.SaveAvatarUrl(ctx, url, userID); err != nil {
		log.Printf("Не удалось удалить временный файл: %v", err)
	}

	return url, nil
}

func saveTmpFile(file multipart.File, fileName string) (string, string, error) {
	uploadPath := "./tmp"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", "", errors.Wrap(err, "Не удалось создать папку для загрузок")
	}

	uniqueFileName := generateFilename(fileName)

	filePath := filepath.Join(uploadPath, uniqueFileName)

	out, err := os.Create(filePath)
	if err != nil {
		return "", "", errors.Wrap(err, "Не удалось создать файл на сервере")
	}

	defer out.Close()

	buffer := make([]byte, 32*1024)
	_, err = copyBuffer(out, file, buffer)
	if err != nil {
		return "", "", errors.Wrap(err, "buffer")
	}

	return filePath, uniqueFileName, nil
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

func generateFilename(original string) string {
	ext := filepath.Ext(original)
	name := filepath.Base(original[:len(original)-len(ext)])
	timestamp := time.Now().UnixNano()

	return fmt.Sprintf("%s_%d%s", name, timestamp, ext)
}

// UploadAvatarLocal загружает файл локально, если нет данных для AWS.
func (a appUser) uploadAvatarLocal(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	userID string,
) (string, error) {
	path, _, err := saveTmpFile(fileForm, header.Filename)
	if err != nil {
		return "", errors.Wrap(err, "tmp")
	}

	if err = a.repo.SaveAvatarLocalPath(ctx, path, userID); err != nil {
		log.Printf("Не удалось удалить временный файл: %v", err)
	}

	return path, nil
}
