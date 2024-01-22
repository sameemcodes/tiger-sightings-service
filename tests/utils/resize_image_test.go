package tests

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"tigerhall-kittens/cmd/utils"
)

func TestResizeImage(t *testing.T) {
	rootDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}
	inputDir := rootDir + "/testimages/testImage.jpeg"
	outputDir := rootDir + "/testimages/resizedImage.jpeg"
	fmt.Println("inputDir", inputDir)
	imageByteArray, err := utils.JPGToByteArray(inputDir)
	if err != nil {
		t.Errorf("Unable to Convert to ByteArray %v but got %v", nil, err)
	}
	resized, err := utils.ResizeImage(imageByteArray, 250, 250)
	if err != nil {
		t.Errorf("Unable to Resize Image %v but got %v", nil, err)
	}
	fmt.Println("resized base64 image", resized)
	if resized == "" {
		t.Errorf("Unable to Resize Image %v but got %v", nil, err)

	}

	//base64 string to test output image
	// Decode the base64 string
	decodedData, err := base64.StdEncoding.DecodeString(resized)
	if err != nil {
		t.Errorf("Unable to Decode Base64 Image %v but got %v", nil, err)

	}

	// Write the decoded data to the output file
	err = os.WriteFile(outputDir, decodedData, 0644)
	if err != nil {
		t.Errorf("Unable to Write Decoded Image %v but got %v", nil, err)
	}

}
