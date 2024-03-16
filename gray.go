package improc

import (
	"image"
	"image/color"
)

func Invers(grid [][]color.Gray) (inversImg [][]color.Gray) {
	xlen, ylen := len(grid), len(grid[0])
	inversImg = make([][]color.Gray, xlen)

	for i := 0; i < len(inversImg); i++ {
		inversImg[i] = make([]color.Gray, ylen)
		for j := 0; j < len(inversImg[i]); j++ {
			grayVal := 255 - grid[i][j].Y
			inversImg[i][j] = color.Gray{Y: grayVal}
		}
	}
	return
}

func Brightness(grid [][]color.Gray, brightness_val int) (brightnessImg [][]color.Gray) {
	xlen, ylen := len(grid), len(grid[0])
	brightnessImg = make([][]color.Gray, xlen)

	for i := 0; i < len(brightnessImg); i++ {
		brightnessImg[i] = make([]color.Gray, ylen)
		for j := 0; j < len(brightnessImg[i]); j++ {
			grayVal := uint8(Clamp(int(grid[i][j].Y) + brightness_val))
			brightnessImg[i][j] = color.Gray{Y: grayVal}
		}
	}
	return
}

// I STILL HAVE AN ERROR FOR THIS FUNCTION :D
func Contrast(img *image.Gray, g, p int) *image.Gray {
	bounds := img.Bounds()
	contrastImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayVal := img.GrayAt(x, y).Y

			adjustedP := Clamp(p)

			contrastFactor := float64(g) / 255.0

			adjustedContrastFactor := contrastFactor * (1 - 2*float64(adjustedP)/255.0)

			grayVal = uint8(Clamp(int(float64(grayVal) * adjustedContrastFactor)))

			contrastImg.SetGray(x, y, color.Gray{Y: grayVal})
		}
	}

	return contrastImg
}
