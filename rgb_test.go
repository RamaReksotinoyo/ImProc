package improc

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGrayscaleRGB(t *testing.T) {
// 	img := [][]color.Color{
// 		{
// 			color.RGBA{40, 54, 25, 255},
// 			color.RGBA{44, 22, 14, 255},
// 			color.RGBA{66, 66, 66, 255},
// 		},
// 	}

// 	expected_img := [][]color.Color{
// 		{
// 			color.RGBA{39, 39, 39, 255},
// 			color.RGBA{26, 26, 26, 255},
// 			color.RGBA{66, 66, 66, 255},
// 		},
// 	}

// 	result := GrayscaleRGB(img)

// 	if !reflect.DeepEqual(result, expected_img) {
// 		t.Errorf("Grayscale result does not match expected value")
// 	}
// }

// func TestInversRGB(t *testing.T) {
// 	img := [][]color.Color{
// 		{
// 			color.RGBA{40, 54, 25, 255},
// 			color.RGBA{66, 66, 66, 255},
// 			color.RGBA{255, 255, 255, 255},
// 		},
// 	}

// 	expected_img := [][]color.Color{
// 		{
// 			color.RGBA{215, 201, 230, 255},
// 			color.RGBA{189, 189, 189, 255},
// 			color.RGBA{0, 0, 0, 255},
// 		},
// 	}

// 	result := InversRGB(img)

// 	if !reflect.DeepEqual(result, expected_img) {
// 		t.Errorf("Invers result does not match expected value")
// 	}
// }

// func TestBrightnessRGB(t *testing.T) {
// 	img := [][]color.Color{
// 		{
// 			color.RGBA{40, 40, 40, 255},
// 			color.RGBA{200, 200, 200, 255},
// 		},
// 	}

// 	expected_img := [][]color.Color{
// 		{
// 			color.RGBA{140, 140, 140, 255},
// 			color.RGBA{255, 255, 255, 255},
// 		},
// 	}

// 	result := BrightnessRGB(img, 100)

// 	if !reflect.DeepEqual(result, expected_img) {
// 		t.Errorf("Brightness result does not match expected value")
// 	}
// }

// func TestContrastRGB(t *testing.T) {
// 	img := [][]color.Color{
// 		{
// 			color.RGBA{100, 120, 150, 255},
// 		},
// 	}

// 	expected_img := [][]color.Color{
// 		{
// 			color.RGBA{54, 60, 68, 255}, // Sesuaikan nilai dengan hasil fungsi Contrast
// 		},
// 	}

// 	result := ContrastRGB(img, 2, 50)

// 	if !reflect.DeepEqual(result, expected_img) {
// 		t.Errorf("Contrast result does not match expected value")
// 	}
// }

func TestInversRGB(t *testing.T) {
	// Mock input image
	inputImg := [][]color.Color{
		{
			color.RGBA{100, 150, 200, 255},
			color.RGBA{50, 100, 150, 255},
		},
		{
			color.RGBA{200, 50, 100, 255},
			color.RGBA{150, 200, 50, 255},
		},
	}

	// Expected output image
	expectedImg := [][]color.Color{
		{
			color.RGBA{155, 105, 55, 255},
			color.RGBA{205, 155, 105, 255},
		},
		{
			color.RGBA{55, 205, 155, 255},
			color.RGBA{105, 55, 205, 255},
		},
	}

	// Call the function
	result := InversRGB(inputImg)

	// Compare the result with the expected output
	assert.Equal(t, expectedImg, result)
}

func TestBrightnessRGB(t *testing.T) {
	// Mock input image
	inputImg := [][]color.Color{
		{
			color.RGBA{100, 150, 200, 255},
			color.RGBA{50, 100, 150, 255},
		},
		{
			color.RGBA{200, 50, 100, 255},
			color.RGBA{150, 200, 50, 255},
		},
	}

	// Expected output image with brightness +50
	expectedImg := [][]color.Color{
		{
			color.RGBA{150, 200, 250, 255},
			color.RGBA{100, 150, 200, 255},
		},
		{
			color.RGBA{250, 100, 150, 255},
			color.RGBA{200, 250, 100, 255},
		},
	}

	// Call the function with brightness +50
	result := BrightnessRGB(inputImg, 50)

	// Compare the result with the expected output
	assert.Equal(t, expectedImg, result)
}
