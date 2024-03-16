package improc

// // CreateImageFromGrayPixels membuat citra grayscale dari piksel-piksel
// func (ip *ImageProcessor) CreateImageFromGrayPixels(pixels [][]color.Gray) *image.Gray {
// 	height := len(pixels)
// 	width := len(pixels[0])

// 	img := image.NewGray(image.Rect(0, 0, width, height))

// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			img.Set(x, y, pixels[y][x])
// 		}
// 	}

// 	return img
// }

// // SaveImage menyimpan citra ke file
// func (ip *ImageProcessor) SaveImage(img image.Image, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	err = png.Encode(file, img)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// CreateImageFromGrayPixels membuat citra grayscale dari piksel-piksel
// func CreateImageFromGrayPixels(pixels [][]color.Gray) *image.Gray {
// 	height := len(pixels)
// 	width := len(pixels[0])

// 	img := image.NewGray(image.Rect(0, 0, width, height))

// 	for y := 0; y < height; y++ {
// 		for x := 0; x < width; x++ {
// 			img.Set(x, y, pixels[y][x])
// 		}
// 	}

// 	return img
// }

// // SaveImage menyimpan citra ke file
// func SaveImage(img image.Image, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	err = png.Encode(file, img)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
