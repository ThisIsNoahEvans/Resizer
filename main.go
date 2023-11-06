package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
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

	// Decode the image to find out the format.
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		os.Exit(1)
	}

	// Seek to the beginning of the file for future reads.
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Error seeking file:", err)
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
		outputFileName := fmt.Sprintf("%s_%dx%d%s", strings.TrimSuffix(imagePath, filepath.Ext(imagePath)), width, height, filepath.Ext(imagePath))
		out, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer out.Close()

		// Encode the image in the original format.
		switch format {
		case "jpeg":
			err = jpeg.Encode(out, m, nil)
		case "png":
			err = png.Encode(out, m)
		case "gif":
			err = gif.Encode(out, m, nil)
		default:
			fmt.Printf("Unsupported image format: %s\n", format)
			err = fmt.Errorf("unsupported image format: %s", format)
		}

		if err != nil {
			fmt.Printf("Error encoding image: %s\n", err)
			continue
		}

		fmt.Printf("Resized image to %dx%d and saved as %s\n", width, height, outputFileName)
	}
}
