package app

import (
	"context"
	"mime/multipart"

	"github.com/pkg/errors"
)

// UploadPhoto TODO: transaction.
func (a appUser) UploadPhoto(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	userID string,
) error {
	u, err := a.repo.GetUserDB(ctx, userID)
	if err != nil {
		return errors.Wrap(err, "get user")
	}

	if err = a.stg.DeleteFile(ctx, u.Avatar); err != nil {
		return errors.Wrap(err, "Delete file")
	}

	fileName, err := a.stg.Upload(ctx, fileForm, header)
	if err != nil {
		return errors.Wrap(err, "UploadPhoto")
	}

	if err = a.repo.SaveAvatarFileName(ctx, fileName, userID); err != nil {
		return errors.Wrap(err, "SaveAvatarFileName")
	}

	return nil
}
