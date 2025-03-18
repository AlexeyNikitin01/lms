package app

import (
	"context"
	"mime/multipart"

	"github.com/pkg/errors"

	"course/internal/repository/pg/entity"
)

// UploadPhoto TODO: transaction.
func (c CourseApp) UploadPhoto(
	ctx context.Context,
	fileForm multipart.File,
	header *multipart.FileHeader,
	course *entity.Course,
) error {
	if err := c.S3.DeleteFile(ctx, course.PhotoURL); err != nil {
		return errors.Wrap(err, "Delete file")
	}

	fileName, err := c.S3.Upload(ctx, fileForm, header)
	if err != nil {
		return errors.Wrap(err, "UploadPhoto")
	}

	if err = c.DB.SaveAvatarCourse(ctx, fileName, course.ID); err != nil {
		return errors.Wrap(err, "SaveAvatarCourse")
	}

	return nil
}
