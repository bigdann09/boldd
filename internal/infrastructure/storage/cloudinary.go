package storage

import (
	"context"
	"mime/multipart"

	"github.com/boldd/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type ICloudinary interface {
	Upload(file multipart.File, folder string) (string, error)
}

type Cloudinary struct {
	ctx context.Context
	cld *cloudinary.Cloudinary
}

func NewCloudinary(cfg *config.CloudinaryConfig) (*Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(
		cfg.CloudName,
		cfg.Key,
		cfg.Secret,
	)
	if err != nil {
		return &Cloudinary{}, err
	}

	return &Cloudinary{
		cld: cld,
		ctx: context.Background(),
	}, nil
}

func (cloudinary *Cloudinary) Upload(file multipart.File, folder string) (string, error) {
	result, err := cloudinary.cld.Upload.Upload(
		cloudinary.ctx,
		file,
		uploader.UploadParams{
			Folder:         folder,
			PublicID:       uuid.NewString(),
			UseFilename:    api.Bool(false),
			UniqueFilename: api.Bool(true),
		},
	)
	return result.URL, err
}
