package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func JPGToByteArray(filePath string) ([]byte, error) {
	// Read the entire JPG file into a byte slice
	imgBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}

// ImageToByteArray reads an image file and returns its byte array representation.
func ImageToByteArray(filePath string) ([]byte, error) {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Create a buffer for the image data
	var buffer bytes.Buffer

	// Encode the image into the buffer
	err = jpeg.Encode(&buffer, img, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func SaveImageToFile(imageData []byte, outputPath string) error {
	// Create the output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the image data to the output file
	_, err = outFile.Write(imageData)
	if err != nil {
		return err
	}

	return nil
}
