# Resizer
A simple and fast tool to resize images in various formats.

## Usage
Resize images by specifying the file paths and desired sizes:

- **Square Size:** Resize to a square by providing a single number:
  
`./resizer <image_path> 1024`

- **Exact Size:** Specify width and height with 'x':
  
`./resizer <image_path> 500x300`

- **Multiple Sizes:** Separate sizes with spaces:
  
`./resizer <image_path> 512x512 1024x1024`

- **Multiple Images:** Include multiple images, even using wildcards:
  
`./resizer images/* 1024x1024`
`./resizer img1.png img2.png 512x512 1024x1024`

Supported formats: JPEG, PNG, GIF.

Output files are saved in the same directory as the source, appended with the specified size. For example, `./resizer /path/to/image.png 1024` creates `/path/to/image_1024x1024.png`.
