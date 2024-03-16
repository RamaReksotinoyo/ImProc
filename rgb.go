// package improc

// import (
// 	"image/color"
// )

// // Grayscale converts an RGB image grid to grayscale.
// //
// // It takes a 2D grid of color values representing an RGB image and returns
// // a new image grid where each pixel is converted to grayscale.
// //
// // Parameters:
// //   - grid: The input RGB image grid.
// //
// // Returns:
// //   - [][]color.Color: The grayscale image grid.
// func GrayscaleRGB(grid [][]color.Color) (grayImg [][]color.Color) {
// 	xlen, ylen := len(grid), len(grid[0])
// 	grayImg = make([][]color.Color, xlen)

// 	for i := 0; i < len(grayImg); i++ {
// 		grayImg[i] = make([]color.Color, ylen)
// 		for j := 0; j < len(grayImg[i]); j++ {
// 			r, g, b, _ := grid[i][j].RGBA()
// 			r8 := uint8(r >> 8)
// 			g8 := uint8(g >> 8)
// 			b8 := uint8(b >> 8)
// 			grayVal := uint8((uint32(r8) + uint32(g8) + uint32(b8)) / 3)
// 			grayImg[i][j] = color.RGBA{grayVal, grayVal, grayVal, 255}
// 		}
// 	}
// 	return
// }

// // Invers performs color inversion on an RGB image grid.
// //
// // It takes a 2D grid of color values representing an RGB image and returns
// // a new image grid where each color component is inverted.
// //
// // Parameters:
// //   - grid: The input RGB image grid.
// //
// // Returns:
// //   - [][]color.Color: The inverted image grid.
// func InversRGB(grid [][]color.Color) (inversImg [][]color.Color) {
// 	xlen, ylen := len(grid), len(grid[0])
// 	inversImg = make([][]color.Color, xlen)

// 	for i := 0; i < len(inversImg); i++ {
// 		inversImg[i] = make([]color.Color, ylen)
// 		for j := 0; j < len(inversImg[i]); j++ {
// 			r, g, b, _ := grid[i][j].RGBA()

// 			r8 := uint8(255 - (int(r) >> 8))
// 			g8 := uint8(255 - (int(g) >> 8))
// 			b8 := uint8(255 - (int(b) >> 8))

// 			inversImg[i][j] = color.RGBA{r8, g8, b8, 255}
// 		}
// 	}
// 	return
// }

// // Brightness adjusts the brightness of an RGB image grid.
// //
// // It takes a 2D grid of color values representing an RGB image and a
// // brightness value. The function returns a new image grid with the
// // applied brightness adjustment.
// //
// // Parameters:
// //   - grid: The input RGB image grid.
// //   - brightness_val: Brightness adjustment value.
// //
// // Returns:
// //   - [][]color.Color: The image grid with the applied brightness adjustment.
// func BrightnessRGB(grid [][]color.Color, brightness_val int) (inversImg [][]color.Color) {
// 	xlen, ylen := len(grid), len(grid[0])
// 	inversImg = make([][]color.Color, xlen)

// 	for i := 0; i < len(inversImg); i++ {
// 		inversImg[i] = make([]color.Color, ylen)
// 		for j := 0; j < len(inversImg[i]); j++ {
// 			r, g, b, _ := grid[i][j].RGBA()

// 			r8 := uint8(Clamp(brightness_val + (int(r) >> 8)))
// 			g8 := uint8(Clamp(brightness_val + (int(g) >> 8)))
// 			b8 := uint8(Clamp(brightness_val + (int(b) >> 8)))

// 			inversImg[i][j] = color.RGBA{r8, g8, b8, 255}
// 		}
// 	}
// 	return
// }

// // Contrast applies a contrast adjustment to the provided image grid.
// //
// // It takes an image grid, a contrast factor 'g', and a brightness shift 'p'.
// // The function returns a new image grid with the applied contrast adjustment.
// //
// // Parameters:
// //   - grid: The input image grid.
// //   - g: Contrast factor. A higher 'g' increases contrast.
// //   - p: Brightness shift. A positive 'p' increases brightness.
// //
// // Returns:
// //   - [][]color.Color: The image grid with the applied contrast adjustment.
// // func Contrast(grid [][]color.Color, g, p int) (contrastImg [][]color.Color) {
// // 	xlen, ylen := len(grid), len(grid[0])
// // 	contrastImg = make([][]color.Color, xlen)

// // 	for i := 0; i < len(contrastImg); i++ {
// // 		contrastImg[i] = make([]color.Color, ylen)
// // 		for j := 0; j < len(contrastImg[i]); j++ {
// // 			r, g, b, _ := grid[i][j].RGBA()

// // 			// Ensure that r, g, b are of type uint32
// // 			r, g, b = r>>8, g>>8, b>>8

// // 			r8 := uint8(helpers.Clamp(int(g)*(int(r)-p) + p))
// // 			g8 := uint8(helpers.Clamp(int(g)*(int(g)-p) + p))
// // 			b8 := uint8(helpers.Clamp(int(g)*(int(b)-p) + p))

// // 			contrastImg[i][j] = color.RGBA{r8, g8, b8, 255}
// // 		}
// // 	}
// // 	return
// // }

// func ContrastRGB(grid [][]color.Color, g, p int) (contrastImg [][]color.Color) {
// 	xlen, ylen := len(grid), len(grid[0])
// 	contrastImg = make([][]color.Color, xlen)

// 	for i := 0; i < xlen; i++ {
// 		contrastImg[i] = make([]color.Color, ylen)
// 		for j := 0; j < ylen; j++ {
// 			r, gVal, b, _ := grid[i][j].RGBA()

// 			// Konversi ke int sebelum melakukan operasi matematika
// 			rInt, gInt, bInt := int(r>>8), int(gVal>>8), int(b>>8)

// 			// Perhitungan kontras untuk setiap komponen warna
// 			r8 := uint8(Clamp(rInt * g / 0xffff))
// 			g8 := uint8(Clamp(gInt * g / 0xffff))
// 			b8 := uint8(Clamp(bInt * g / 0xffff))

// 			// Simpan hasil ke dalam citra kontras
// 			contrastImg[i][j] = color.RGBA{r8, g8, b8, 255}
// 		}
// 	}
// 	return
// }

// func To_Gray(grid [][]color.Color) (grayImg [][]color.Gray) {
// 	xlen, ylen := len(grid), len(grid[0])
// 	grayImg = make([][]color.Gray, xlen)

// 	for i := 0; i < len(grayImg); i++ {
// 		grayImg[i] = make([]color.Gray, ylen)
// 		for j := 0; j < len(grayImg[i]); j++ {
// 			r, g, b, _ := grid[i][j].RGBA()
// 			r8 := uint8(r >> 8)
// 			g8 := uint8(g >> 8)
// 			b8 := uint8(b >> 8)
// 			grayVal := uint8((uint32(r8) + uint32(g8) + uint32(b8)) / 3)
// 			grayImg[i][j] = color.Gray{Y: grayVal}
// 		}
// 	}
// 	return
// }

package improc

import (
	"image/color"
)

func InversRGB(grid [][]color.Color) (inversImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	inversImg = make([][]color.Color, xlen)

	for i := 0; i < len(inversImg); i++ {
		inversImg[i] = make([]color.Color, ylen)
		for j := 0; j < len(inversImg[i]); j++ {
			r, g, b, _ := grid[i][j].RGBA()

			r8 := uint8(255 - (int(r) >> 8))
			g8 := uint8(255 - (int(g) >> 8))
			b8 := uint8(255 - (int(b) >> 8))

			inversImg[i][j] = color.RGBA{r8, g8, b8, 255}
		}
	}
	return
}

func BrightnessRGB(grid [][]color.Color, brightness_val int) (brightnessImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0])
	brightnessImg = make([][]color.Color, xlen)

	for i := 0; i < len(brightnessImg); i++ {
		brightnessImg[i] = make([]color.Color, ylen)
		for j := 0; j < len(brightnessImg[i]); j++ {
			r, g, b, _ := grid[i][j].RGBA()

			r8 := uint8(Clamp(brightness_val + (int(r) >> 8)))
			g8 := uint8(Clamp(brightness_val + (int(g) >> 8)))
			b8 := uint8(Clamp(brightness_val + (int(b) >> 8)))

			brightnessImg[i][j] = color.RGBA{r8, g8, b8, 255}
		}
	}
	return
}
