package util

import (
	"image"
	"image/png"
	"mime/multipart"
	"os"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func SavePhotoAndOptimze(photo *multipart.FileHeader) (string, error) {
	inputPhoto, err := photo.Open()
	if err != nil {
		return "", err
	}
	defer inputPhoto.Close()

	//Decode the input image
	img, _, err := image.Decode(inputPhoto)
	if err != nil {
		return "", err
	}
	//TODO: Remove EXIF from photo
	//Resize the image if needed
	img = resize.Resize(300, 300, img, resize.Lanczos3)

	//Optimize the image for the web
	img = imaging.AdjustBrightness(img, 10)

	// Save the image in PNG format
	outputFileName := uuid.NewString() + ".png"
	outdir := "public/profile/" + outputFileName
	outputFile, err := os.Create(outdir)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	//Save the PNG Image
	err = png.Encode(outputFile, img)
	if err != nil {
		return "", err
	}

	return outputFileName, nil
}
