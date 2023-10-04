package upload

import (
	"context"
	"mime/multipart"
)

type Upload interface {
	//this method will return url of image as string
	UploadImage(ctx context.Context, file multipart.File) (url, publicID string, err error)
	//remove image using public id
	RemoveImage(ctx context.Context, publicID string) error
}
