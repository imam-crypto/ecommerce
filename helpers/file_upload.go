package helpers

import (
	"context"
	"ecommerce/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"log"
	"time"
)

func ImageUpload(upload interface{}, publicID string) (string, string, error) {

	config, configErr := utils.LoadConfig("..", false)
	if configErr != nil {
		log.Fatal("[CONFIGURATION] Can not load configuration file.")
	}

	var allowedType [4]string
	allowedType[0] = "jpg"
	allowedType[1] = "png"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(config.CloudName, config.CloudApiKey, config.CloudApiSecret)
	if err != nil {
		return "", "", err
	}
	up, err := cld.Upload.Upload(ctx, upload, uploader.UploadParams{PublicID: publicID, Folder: config.CloudUploadCategoryImage, AllowedFormats: config.CloudAllowTypeImage})
	//uploadParam, err := cld.Add.Add.(ctx, upload, uploader.UploadParams{Folder: cl.CloudUploadCategoryImage})
	if err != nil {
		return "", "", err
	}
	if up.SecureURL == "" {
		return "", "", err
	}
	return up.SecureURL, up.PublicID, nil
}
func FileDestroy(publicID string) (string, error) {

	cl, configErr := utils.LoadConfig("..", false)
	if configErr != nil {
		log.Fatal("[CONFIGURATION] Can not load configuration file.")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(cl.CloudName, cl.CloudApiKey, cl.CloudApiSecret)
	if err != nil {
		return "", err
	}

	resp, errDel := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})

	if errDel != nil {
		return "", errDel
	}
	return resp.Result, nil
}
