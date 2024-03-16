package main

import (
	"fmt"
	"log"
	"os"

	improc "github.com/RamaReksotinoyo/ImProc"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`Args required.`)
		return
	}
	log.Printf("loading %s", os.Args[1])

	var ip improc.ImageProcessor
	imageData, err := ip.LoadImage(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var memo []uint8

	for y := 0; y < imageData.Height; y++ {
		for x := 0; x < imageData.Width; x++ {
			pixel := imageData.Pixels[y][x]
			memo = append(memo, uint8(pixel.Y))
		}
	}

	var hist improc.Histogram
	hist.Nk = ip.Nk(memo)
	hist.Pdf = ip.Pdf(hist.Nk, imageData.Height*imageData.Width)
	hist.Cdf = ip.Cdf(hist.Pdf)
	hist.CdfMultiply = ip.CdfMultiply(hist.Cdf, 255)
	hist.HistEQ = ip.HistogramEqualizationLevel(hist.CdfMultiply)

	equalizedImage := ip.ApplyHistogramEqualization(imageData.Pixels, hist.HistEQ)

	outputImage := improc.CreateImageFromGrayPixels(equalizedImage)

	err = improc.SaveImage(outputImage, "output_lena_v2.png")
	if err != nil {
		fmt.Println("Failed to save image:", err)
	}
}
