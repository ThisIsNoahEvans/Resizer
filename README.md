# Resizer
Easily resize images

## Usage:
`Usage: ./resizer <image_path> <size> ...`

Sizes can be a single number, to resize into a square:
`./resizer <image_path> 1024 ...`

Or they can be exact, using 'x' as a splitter:
`./resizer <image_path> 500x300 ...`

Individual sizes are separated by a space.

Add as many sizes as you want, mixing between both square & exact measurements.

Supports JPEG, PNG, and GIF.

Results will be stored in the same directory as the input file, with the size at the end:

`./resizer /path/to/my/image.png 1024`

will result in a file named 

`/path/to/my/image_1024x1024.png`