package cloudinary

import (
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"log"
	"path"
	"testing"
)

const root = "__fixtures__"

func TestCloudinary_upload(t *testing.T) {

	err := godotenv.Load(path.Join(root, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	asset := Asset{
		PublicID:       "test-icon",
		Path:           path.Join(root, "icon50.png"),
		UniqueFilename: false,
		Overwrite:      true,
	}

	tCld := New()
	actual, err := tCld.Upload(asset)

	if err != nil {
		t.Fail()
		log.Fatal(err)
	}

	if actual == "" {
		t.Fail()
	}

	_, err = tCld.cld.Upload.Destroy(tCld.ctx, uploader.DestroyParams{PublicID: asset.PublicID})

	if err != nil {
		msg := fmt.Sprintf("Error cleaning up %s", asset.PublicID)
		log.Fatal(msg)

	}
}
