package app

import (
	"context"
	"mime/multipart"

	"github.com/pkg/errors"
)

// UploadPhoto TODO: delete old photo.
func (a appUser) UploadPhoto(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	userID string,
) error {
	fileName, err := a.stg.Upload(ctx, fileForm, header)
	if err != nil {
		return errors.Wrap(err, "UploadPhoto")
	}

	if err = a.repo.SaveAvatarFileName(ctx, fileName, userID); err != nil {
		return errors.Wrap(err, "SaveAvatarFileName")
	}

	return nil
}
