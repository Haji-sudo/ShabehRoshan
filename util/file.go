package util

import (
	"image"
	"image/png"
	"mime/multipart"
	"os"

	_ "github.com/disintegration/imaging"
	"github.com/google/uuid"
	_ "github.com/nfnt/resize"
)

func SaveProfileImageAndOptimize(photo *multipart.FileHeader) (string, error) {
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

func SaveImageAndOptimize(photo *multipart.FileHeader) (string, error) {
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

	// Save the image in PNG format
	outputFileName := uuid.NewString() + ".png"
	outdir := "public/images/" + outputFileName
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
