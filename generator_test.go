package improc

import (
	"image"
	"image/color"
	"os"
	"testing"
)

func TestCreateImageFromGrayPixels(t *testing.T) {
	// Persiapan data piksel grayscale untuk diuji
	pixels := [][]color.Gray{
		{{Y: 100}, {Y: 150}, {Y: 200}},
		{{Y: 50}, {Y: 100}, {Y: 150}},
		{{Y: 0}, {Y: 50}, {Y: 100}},
	}

	// Panggil fungsi yang akan diuji
	img := CreateImageFromGrayPixels(pixels)

	// Periksa apakah ukuran gambar sesuai dengan yang diharapkan
	expectedWidth := 3
	expectedHeight := 3
	if img.Bounds().Dx() != expectedWidth || img.Bounds().Dy() != expectedHeight {
		t.Errorf("CreateImageFromGrayPixels() returned image with incorrect dimensions, got (%d, %d), expected (%d, %d)", img.Bounds().Dx(), img.Bounds().Dy(), expectedWidth, expectedHeight)
	}

	// Periksa apakah nilai piksel sesuai dengan yang diharapkan
	expectedPixelValues := [][]uint8{
		{100, 150, 200},
		{50, 100, 150},
		{0, 50, 100},
	}
	for y := 0; y < expectedHeight; y++ {
		for x := 0; x < expectedWidth; x++ {
			if img.GrayAt(x, y).Y != expectedPixelValues[y][x] {
				t.Errorf("CreateImageFromGrayPixels() returned image with incorrect pixel value at position (%d, %d), got %d, expected %d", x, y, img.GrayAt(x, y).Y, expectedPixelValues[y][x])
			}
		}
	}
}

func TestSaveImage(t *testing.T) {
	// Persiapan gambar untuk disimpan
	img := image.NewGray(image.Rect(0, 0, 2, 2))
	img.SetGray(0, 0, color.Gray{Y: 100})
	img.SetGray(1, 0, color.Gray{Y: 150})
	img.SetGray(0, 1, color.Gray{Y: 200})
	img.SetGray(1, 1, color.Gray{Y: 250})

	// Simpan gambar ke file
	filename := "test_image.png"
	err := SaveImage(img, filename)
	if err != nil {
		t.Fatalf("SaveImage() failed: %v", err)
	}
	defer func() {
		// Hapus file yang telah dibuat setelah selesai pengujian
		err := os.Remove(filename)
		if err != nil {
			t.Fatalf("Failed to remove test image file: %v", err)
		}
	}()

	// Baca kembali gambar dari file yang disimpan
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open saved image file: %v", err)
	}
	defer file.Close()

	decodedImg, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("Failed to decode saved image file: %v", err)
	}

	// Periksa apakah ukuran gambar sama
	if decodedImg.Bounds().Dx() != img.Bounds().Dx() || decodedImg.Bounds().Dy() != img.Bounds().Dy() {
		t.Errorf("Saved image has incorrect dimensions, got (%d, %d), expected (%d, %d)", decodedImg.Bounds().Dx(), decodedImg.Bounds().Dy(), img.Bounds().Dx(), img.Bounds().Dy())
	}

	// Periksa apakah nilai piksel sesuai
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			if img.GrayAt(x, y).Y != decodedImg.(*image.Gray).GrayAt(x, y).Y {
				t.Errorf("Saved image has incorrect pixel value at position (%d, %d), got %d, expected %d", x, y, decodedImg.(*image.Gray).GrayAt(x, y).Y, img.GrayAt(x, y).Y)
			}
		}
	}
}
