package main

import (
	"fmt"
	"image"
	_ "image/png"
	"image/jpeg"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: resizer <image_path> <size> ...")
		os.Exit(1)
	}

	imagePath := os.Args[1]
	sizes := os.Args[2:]

	// Open the file.
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the image.
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		os.Exit(1)
	}

	for _, size := range sizes {
		dimensions := strings.Split(size, "x")
		var width, height uint
		var err error

		// Parse width and height.
		width64, err := strconv.ParseUint(dimensions[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid width:", err)
			continue
		}
		width = uint(width64)

		if len(dimensions) == 2 {
			height64, err := strconv.ParseUint(dimensions[1], 10, 64)
			if err != nil {
				fmt.Println("Invalid height:", err)
				continue
			}
			height = uint(height64)
		} else {
			height = width
		}

		// Resize the image.
		m := resize.Resize(width, height, img, resize.Lanczos3)

		// Create the output file.
		out, err := os.Create(fmt.Sprintf("%s_%dx%d.jpg", strings.TrimSuffix(imagePath, ".png"), width, height))
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer out.Close()

		// Write the new image to file.
		jpeg.Encode(out, m, nil)
		fmt.Printf("Resized image to %dx%d and saved as %s_%dx%d.jpg\n", width, height, strings.TrimSuffix(imagePath, ".png"), width, height)
	}
}
