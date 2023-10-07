package helpers

import (
	"context"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUpload(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cloudName string = os.Getenv("CLOUDINARY_CLOUD_NAME")
	var cloudAPIKey string = os.Getenv("CLOUDINARY_API_KEY")
	var cloudAPISecret string = os.Getenv("CLOUDINARY_API_SECRET")
	var cloudUploadFolder string = os.Getenv("CLOUDINARY_UPLOAD_FOLDER")

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(cloudName, cloudAPIKey, cloudAPISecret)
	if err != nil {
		log.Error("No se pudo crear la instancia de cloudinary: ", err)
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: cloudUploadFolder})
	if err != nil {
		log.Error("No se pudo subir la imagen: ", err)
		return "", err
	}
	return uploadParam.SecureURL, nil
}
