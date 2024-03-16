package rgb

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

func Load(filePath string) *image.RGBA {
	imgFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	buffer := make([]byte, 512) // docs tell that it takes only the first 512 bytes into consideration
	_, err = imgFile.Read(buffer)
	if err != nil {
		panic(err)
	}
	format := http.DetectContentType(buffer)

	imgFile.Seek(0, io.SeekStart)

	fmt.Printf("format: %s\n", format)
	switch format {
	case "image/png":
		img, err := png.Decode(imgFile)
		if err != nil {
			panic(err)
		}
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)
		return rgba
	case "image/jpeg":
		img, err := jpeg.Decode(imgFile)
		if err != nil {
			panic(err)
		}
		bounds := img.Bounds()
		rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(rgba, rgba.Bounds(), img, bounds.Min, draw.Src)
		return rgba
	}
	return nil
}

func Load_grid(img *image.RGBA) (grid [][]color.Color) {
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			data := img.At(i, j)
			y = append(y, data)
		}
		grid = append(grid, y)
	}
	return
}

func ConvertToGrayscale(img *image.RGBA) *image.Gray {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	draw.Draw(grayImg, bounds, img, bounds.Min, draw.Src)
	return grayImg
}

func Save_grid(filepath string, grid [][]color.Gray) {
	len_x, len_y := len(grid), len(grid[0])
	rect := image.Rect(0, 0, len_x, len_y)
	img := image.NewRGBA(rect)

	for x := 0; x < len_x; x++ {
		for y := 0; y < len_y; y++ { // Perbaikan batas perulangan
			img.Set(x, y, grid[x][y])
		}
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		// Handle the error if image encoding fails
		panic(err)
	}
}
