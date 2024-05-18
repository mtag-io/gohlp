package cloudinary

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/mtag-io/gohlp/term"
	"os"
)

// This method presumes you have already loaded the
// .env variables as described in the README.md file

func (that *Class) credentials() {
	var err error
	that.cld, err = cloudinary.New()
	if err != nil {
		term.Abort("Unable to initialize the Cloudinary class.")
	}
	that.cld.Config.URL.Secure = true
	that.ctx = context.Background()
}

func (that *Class) Upload(asset Asset) (string, error) {

	if _, err := os.Stat(asset.Path); err != nil {
		msg := fmt.Sprintf("Path not found %s", asset.Path)
		term.Abort(msg)
	}

	resp, uErr := that.cld.Upload.Upload(that.ctx, asset.Path, uploader.UploadParams{
		PublicID:       asset.PublicID,
		UniqueFilename: &asset.UniqueFilename,
		Overwrite:      &asset.Overwrite,
	})

	if uErr != nil {
		term.Abort(uErr.Error())
		return "", uErr
	}

	if resp == nil {
		msg := "cloudinary upload returned an invalid response"
		term.Abort(msg)
		return "", errors.New(msg)
	}

	if resp.SecureURL == "" {
		msg := fmt.Sprintf("Upload failed for %s.", asset.PublicID)
		term.Error(msg)
		return "", errors.New(msg)
	}

	msg := fmt.Sprintf("Succesfully uploaded %s.", asset.PublicID)
	term.Ok(msg)
	return resp.SecureURL, nil
}
