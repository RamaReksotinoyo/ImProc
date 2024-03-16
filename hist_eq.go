package improc

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

// ImageData menyimpan informasi tentang citra grayscale
type ImageData struct {
	Pixels [][]color.Gray
	Width  int
	Height int
}

// Histogram menyimpan informasi tentang histogram suatu citra
type Histogram struct {
	Nk          map[uint8]int
	Pdf         map[uint8]float64
	Cdf         map[uint8]float64
	CdfMultiply map[uint8]float64
	HistEQ      map[uint8]int
}

// ImageProcessor menyimpan fungsi-fungsi untuk memproses citra
type ImageProcessor struct {
}

// LoadImage untuk membaca citra dari file
func (ip *ImageProcessor) LoadImage(filepath string) (*ImageData, error) {
	// gray, err := ioutils.LoadImage(filepath)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to load file: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open png format: %v", err)
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

	if err != nil {
		return nil, err
	}

	new_width := len(pixels[0])
	new_height := len(pixels)

	return &ImageData{
		Pixels: pixels,
		Width:  new_width,
		Height: new_height,
	}, nil
}

// Nk menghitung nilai Nk (jumlah piksel dengan intensitas tertentu)
func (ip *ImageProcessor) Nk(pix []uint8) map[uint8]int {
	count := make(map[uint8]int)
	for i := 0; i <= 255; i++ {
		count[uint8(i)] = 0
	}
	for _, v := range pix {
		count[v]++
	}
	return count
}

// Pdf menghitung nilai Pdf (probabilitas piksel)
func (ip *ImageProcessor) Pdf(pix map[uint8]int, nPix int) map[uint8]float64 {
	prob := make(map[uint8]float64)
	// maxPixValue := 215

	var maxPixValue int
	for _, v := range pix {
		if v > maxPixValue {
			maxPixValue = v
		}
	}

	for i := 0; i <= maxPixValue; i++ {
		pixel := uint8(i)
		v, exists := pix[pixel]
		if exists {
			prob[pixel] = float64(v) / float64(nPix)
		} else {
			prob[pixel] = 0.0
		}
	}
	return prob
}

func (ip *ImageProcessor) Cdf(prob map[uint8]float64) map[uint8]float64 {
	sk := make(map[uint8]float64)
	var temp float64

	// Ubah tipe data variabel i menjadi int dan batasi iterasi sesuai panjang prob
	for i := 0; i <= len(prob)-1; i++ {
		pixel := uint8(i)
		if i == 0 {
			sk[pixel] = prob[pixel]
		} else {
			temp = sk[uint8(i-1)]
			sk[pixel] = temp + prob[pixel]
		}
	}
	return sk
}

// CdfMultiply menghitung nilai CdfMultiply (perkalian distribusi kumulatif dengan nilai maksimum)
func (ip *ImageProcessor) CdfMultiply(cdf map[uint8]float64, maxValue int) map[uint8]float64 {
	sk := make(map[uint8]float64)
	for i := 0; i <= maxValue; i++ {
		sk[uint8(i)] = cdf[uint8(i)] * float64(maxValue)
	}
	return sk
}

// HistogramEqualizationLevel menghitung tingkat ekualisasi histogram
func (ip *ImageProcessor) HistogramEqualizationLevel(cdfMultiply map[uint8]float64) map[uint8]int {
	histEQ := make(map[uint8]int)

	for i := 0; i <= 255; i++ {
		histEQ[uint8(i)] = int(cdfMultiply[uint8(i)] + 0.5)
	}

	return histEQ
}

// ApplyHistogramEqualization menerapkan ekualisasi histogram pada citra grayscale
func (ip *ImageProcessor) ApplyHistogramEqualization(originalImage [][]color.Gray, eqLevel map[uint8]int) [][]color.Gray {
	height := len(originalImage)
	width := len(originalImage[0])

	// Buat salinan baru dari gambar asli
	equalizedImage := make([][]color.Gray, height)
	for y := 0; y < height; y++ {
		equalizedImage[y] = make([]color.Gray, width)
		copy(equalizedImage[y], originalImage[y])
	}

	// Terapkan pemetaan EQ Level pada setiap piksel
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := originalImage[y][x]
			equalizedPixelValue := eqLevel[uint8(pixel.Y)]
			equalizedImage[y][x].Y = uint8(equalizedPixelValue)
		}
	}

	return equalizedImage
}
