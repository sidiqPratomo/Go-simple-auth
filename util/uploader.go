package util

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
)

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}
	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}

func DeleteInCloudinary(filePath string) error {
	ctx := context.Background()
	cld, err := config.SetupCloudinary()
	if err != nil {
		return err
	}
	publicId := getPublicIdFromUrl(filePath)
	deleteParams := uploader.DestroyParams{
		PublicID:     publicId,
		ResourceType: "image",
	}
	_, err = cld.Upload.Destroy(ctx, deleteParams)
	if err != nil {
		return err
	}
	return nil
}

func getPublicIdFromUrl(filePath string) string {
	splitUrl := strings.Split(filePath, "/")
	fileSplit := strings.Split(splitUrl[len(splitUrl)-1], ".")
	publicId := splitUrl[len(splitUrl)-2] + "/" + fileSplit[0]
	return publicId
}
