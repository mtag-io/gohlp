package cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
)

type Asset struct {
	Path           string
	PublicID       string
	UniqueFilename bool
	Overwrite      bool
}

type Class struct {
	ctx context.Context
	cld *cloudinary.Cloudinary
}
