package tests

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"tigerhall-kittens/cmd/utils"
)

func TestResizeImage(t *testing.T) {
	imageByteArray, err := utils.JPGToByteArray("/Users/msameem/Dev/tigerhall-kittens/tests/testimages/testImage.jpeg")
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
	err = os.WriteFile("/Users/msameem/Dev/tigerhall-kittens/tests/testimages/decoded_image.jpg", decodedData, 0644)
	if err != nil {
		t.Errorf("Unable to Write Decoded Image %v but got %v", nil, err)
	}

}
