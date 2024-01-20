package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/nfnt/resize"
)

// ResizeImage resizes an image to the specified width and height.
func ResizeImage(inputData []byte, width, height uint) (string, error) {
	// Create a buffer from the input data
	buf := bytes.NewBuffer(inputData)

	// Decode the image
	img, _, err := image.Decode(buf)
	if err != nil {
		return "Invalid Image", err
	}

	// Resize the image
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Create a buffer for the output image
	var outputBuffer bytes.Buffer

	// Save the resized image to the output buffer
	err = jpeg.Encode(&outputBuffer, resizedImg, nil)
	if err != nil {
		return "Invalid Image", err
	}

	base64String := base64.StdEncoding.EncodeToString(outputBuffer.Bytes())
	fmt.Println("base64String", base64String)
	return base64String, nil
}
