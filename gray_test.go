package improc

import (
	"image/color"
	"reflect"
	"testing"
)

func TestInvers(t *testing.T) {
	img := [][]color.Gray{
		{
			color.Gray{Y: 40},
			color.Gray{Y: 66},
			color.Gray{Y: 255},
		},
	}

	expectedImg := [][]color.Gray{
		{
			color.Gray{Y: 215},
			color.Gray{Y: 189},
			color.Gray{Y: 0},
		},
	}

	result := Invers(img)

	if !reflect.DeepEqual(result, expectedImg) {
		t.Errorf("Invers result does not match expected value")
	}
}

func TestBrightness(t *testing.T) {
	img := [][]color.Gray{
		{
			color.Gray{Y: 40},
			color.Gray{Y: 200},
		},
	}

	expectedImg := [][]color.Gray{
		{
			color.Gray{Y: 140},
			color.Gray{Y: 255},
		},
	}

	result := Brightness(img, 100)

	if !reflect.DeepEqual(result, expectedImg) {
		t.Errorf("Brightness result does not match expected value")
	}
}
