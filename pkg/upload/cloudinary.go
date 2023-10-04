package upload

import (
	"context"
	"mime/multipart"
	"project-adhyaksa/pkg/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.uber.org/zap"
)

type cloudinaryUpload struct {
	Cloudinary *cloudinary.Cloudinary
	config     *config.Config
}

func NewCloudinaryUpload(Cloudinary *cloudinary.Cloudinary, config *config.Config) Upload {
	return &cloudinaryUpload{Cloudinary: Cloudinary, config: config}
}

func (c *cloudinaryUpload) UploadImage(ctx context.Context, file multipart.File) (url, publicID string, err error) {

	uploadParam, err := c.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{Folder: c.config.Cloudinary.Folder})
	if err != nil {
		zap.L().Error(err.Error())
		return "", "", err
	}

	return uploadParam.SecureURL, uploadParam.PublicID, nil
}

func (c *cloudinaryUpload) RemoveImage(ctx context.Context, publicID string) error {
	if _, err := c.Cloudinary.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID}); err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}
