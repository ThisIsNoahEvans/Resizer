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
		fmt.Println("Usage: resizer <image_path>... <size>...")
		os.Exit(1)
	}

	var images []string
	var sizes [][]uint

	// Separate image paths and sizes
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "x") || (len(arg) == 3 && strings.Contains("0123456789", string(arg[0]))) {
			// Parse size
			dimensions := strings.Split(arg, "x")
			if len(dimensions) != 2 {
				fmt.Println("Invalid size format:", arg)
				continue
			}
			width, err := strconv.ParseUint(dimensions[0], 10, 64)
			if err != nil {
				fmt.Println("Invalid width:", err)
				continue
			}
			height, err := strconv.ParseUint(dimensions[1], 10, 64)
			if err != nil {
				fmt.Println("Invalid height:", err)
				continue
			}
			sizes = append(sizes, []uint{uint(width), uint(height)})
		} else {
			// Assume it's a file path
			images = append(images, arg)
		}
	}

	if len(images) == 0 || len(sizes) == 0 {
		fmt.Println("No valid images or sizes provided.")
		os.Exit(1)
	}

	// Process each image with each size
	for _, imagePath := range images {
		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		img, format, err := image.Decode(file)
		if err != nil {
			fmt.Println("Error decoding image:", err)
			continue
		}

		for _, size := range sizes {
			width, height := size[0], size[1]
			m := resize.Resize(width, height, img, resize.Lanczos3)

			outputFileName := fmt.Sprintf("%s_%dx%d%s", strings.TrimSuffix(imagePath, filepath.Ext(imagePath)), width, height, filepath.Ext(imagePath))
			out, err := os.Create(outputFileName)
			if err != nil {
				fmt.Println("Error creating file:", err)
				continue
			}
			defer out.Close()

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
}