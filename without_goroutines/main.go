package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"strings"
	"time"
)

func main() {
	// Start the timer
	startTime := time.Now()
	//Define image paths
	imagePaths := []string{
		"images/image1.jpg",
		"images/image2.jpg",
		"images/image3.jpg",
		"images/image4.jpg",
	}

	// Process image one at a time (iterate over image paths with all steps)
	for _, path := range imagePaths {
		outPath := strings.Replace(path, "images/", "images/output/", 1)

		img := imageprocessing.ReadImage(path)
		img = imageprocessing.Resize(img)
		img = imageprocessing.Grayscale(img)
		imageprocessing.WriteImage(outPath, img)
		fmt.Println("Success!")
	}
	//stop time and prints
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total processing time: %s\n", elapsedTime)
}
