#!/bin/bash
CURRENT_DIR=$(pwd)

# Directory containing PNG images
IMAGE_PATH="$CURRENT_DIR/test_images/image.png"

# Output directory for Resizer
RESIZER_OUTPUT_DIR="$CURRENT_DIR/test_images/resizer-output"

# Output directory for ImageMagick
IMAGEMAGICK_OUTPUT_DIR="$CURRENT_DIR/test_images/imagemagick-output"

# Ensure output directories exist
mkdir -p "$RESIZER_OUTPUT_DIR"
mkdir -p "$IMAGEMAGICK_OUTPUT_DIR"

# Path to the Resizer executable
RESIZER_PATH="./resizer"

# Number of iterations
ITERATIONS=100

# Start timing Resizer
START_TIME=$(date +%s)
for i in $(seq 1 $ITERATIONS); do
    $RESIZER_PATH "$IMAGE_PATH" 1024x1024
done
END_TIME=$(date +%s)
RESIZER_DURATION=$((END_TIME - START_TIME))
echo "Resizer took $RESIZER_DURATION seconds."

# Move Resizer output to its directory
mv "${IMAGE_PATH%.*}"*_1024x1024.png "$RESIZER_OUTPUT_DIR"

# Start timing ImageMagick
START_TIME=$(date +%s)
for i in $(seq 1 $ITERATIONS); do
    convert "$IMAGE_PATH" -resize 1024x1024 "$IMAGEMAGICK_OUTPUT_DIR/$(basename "${IMAGE_PATH%.*}")_$i.png"
done
END_TIME=$(date +%s)
IMAGEMAGICK_DURATION=$((END_TIME - START_TIME))
echo "ImageMagick took $IMAGEMAGICK_DURATION seconds."

# Compare times
echo "Comparison:"
echo "Resizer: $RESIZER_DURATION seconds"
echo "ImageMagick: $IMAGEMAGICK_DURATION seconds"