package config

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cldName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cldKey := os.Getenv("CLOUDINARY_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
