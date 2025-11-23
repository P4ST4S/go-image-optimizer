package main

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/disintegration/imaging"
)

// ResizeImage resizes an image to the specified width and encodes it as JPEG
func ResizeImage(img image.Image, width int, quality int) (*bytes.Buffer, error) {
	// Resize the image
	dstImage := imaging.Resize(img, width, 0, imaging.Lanczos)

	// Encode as JPEG
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, dstImage, &jpeg.Options{Quality: quality}); err != nil {
		return nil, err
	}

	return buf, nil
}
