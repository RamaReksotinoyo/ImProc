package improc

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func CreateImageFromGrayPixels(pixels [][]color.Gray) *image.Gray {
	height := len(pixels)
	width := len(pixels[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, pixels[y][x])
		}
	}

	return img
}

// SaveImage menyimpan citra ke file
func SaveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func DummyRGB(h, w int) [][]color.Gray {
	img := image.NewGray(image.Rect(0, 0, w, h))

	pixels := [][]uint8{
		{4, 4, 4, 4, 4},
		{3, 4, 5, 4, 3},
		{3, 5, 5, 5, 3},
		{3, 4, 5, 4, 3},
		{4, 4, 4, 4, 4},
	}

	// Setel nilai piksel berdasarkan array
	for y := 0; y < len(pixels); y++ {
		for x := 0; x < len(pixels[y]); x++ {
			img.Set(x, y, color.Gray{Y: pixels[y][x]})
		}
	}

	// Simpan citra ke file PNG
	file, err := os.Create("proceed_img/dummy.png")
	if err != nil {
		fmt.Println("Gagal menyimpan citra:", err)
		return nil
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Gagal menyimpan citra:", err)
		return nil
	}

	// Konversi tipe data dari []uint8 ke [][]color.Gray
	pix := make([][]color.Gray, h)
	for y := 0; y < h; y++ {
		pix[y] = make([]color.Gray, w)
		for x := 0; x < w; x++ {
			pix[y][x] = img.GrayAt(x, y)
		}
	}

	return pix
}

func LoadImage(filename string) ([][]color.Gray, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca citra PNG: %v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := make([][]color.Gray, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]color.Gray, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = color.GrayModel.Convert(img.At(x, y)).(color.Gray)
		}
	}

	return pixels, nil
}
