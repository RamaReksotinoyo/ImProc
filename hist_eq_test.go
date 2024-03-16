package improc

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageProcessor_LoadImage(t *testing.T) {
	ip := ImageProcessor{}

	// Test cases
	testCases := []struct {
		filePath                      string
		expectedWidth, expectedHeight int
	}{
		{"dummy_output.png", 5, 5}, // Adjust the expected width and height based on your test image
		// Add more test cases if needed
	}

	for _, tc := range testCases {
		imageData, err := ip.LoadImage("dummy_output.png")

		assert.NoError(t, err)
		assert.NotNil(t, imageData)
		assert.Equal(t, tc.expectedWidth, imageData.Width)
		assert.Equal(t, tc.expectedHeight, imageData.Height)
	}
}

func TestImageProcessor_ApplyHistogramEqualization(t *testing.T) {
	ip := ImageProcessor{}

	// Mock grayscale image data
	grayscaleImageData := [][]color.Gray{
		{
			{Y: 50},
			{Y: 100},
			{Y: 150},
		},
		{
			{Y: 200},
			{Y: 100},
			{Y: 50},
		},
	}

	// Mock histogram equalization level
	histEQ := map[uint8]int{
		50:  25,
		100: 75,
		150: 125,
		200: 175,
	}

	// Expected result after applying histogram equalization
	expectedEqualizedImage := [][]color.Gray{
		{
			{Y: 25},
			{Y: 75},
			{Y: 125},
		},
		{
			{Y: 175},
			{Y: 75},
			{Y: 25},
		},
	}

	// Apply histogram equalization
	equalizedImage := ip.ApplyHistogramEqualization(grayscaleImageData, histEQ)

	// Assert the result
	assert.Equal(t, expectedEqualizedImage, equalizedImage)
}
