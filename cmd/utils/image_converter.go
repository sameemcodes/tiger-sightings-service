package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
)

func JPGToByteArray(filePath string) ([]byte, error) {
	imgBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return imgBytes, nil
}

func ImageToByteArray(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	err = jpeg.Encode(&buffer, img, nil)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func SaveImageToFile(imageData []byte, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.Write(imageData)
	if err != nil {
		return err
	}

	return nil
}
