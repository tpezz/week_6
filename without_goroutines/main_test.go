package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"testing"
)

func TestGrayscaleSimpler(t *testing.T) {
	// set image url (all red image)
	img := imageprocessing.ReadImage("images/test.jpg")

	// Convert the image to grayscale
	grayImg := imageprocessing.Grayscale(img)

	// get pixel color
	c := grayImg.At(0, 0)
	r, g, b, _ := c.RGBA()

	//Test
	if r != g || r != b {
		t.Errorf("Expected grayscale pixel (R=G=B), got R=%d, G=%d, B=%d", r, g, b)
	}
}
func TestResize(t *testing.T) {
	// set image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	//resize image
	resized := imageprocessing.Resize(img)
	//test
	if resized.Bounds().Dx() == 0 || resized.Bounds().Dy() == 0 {
		t.Errorf("Expected non-zero dimensions, got %v", resized.Bounds())
	}
}
