package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"math"
	"os"
	"strings"

	"github.com/aybabtme/rgbterm"
)

const grayScale = " .'`^,:;Il!i><~+_-?][}{1)(|/tfjrxnuvczXYUJCLQ0OZnwqpdbkhao*#MW&8%B@$"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of arguments!")
		fmt.Println("Usage: ./GoASCII image.jpg")
	}

	img_path := args[0]
	reader, err := os.Open(img_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reader.Name())

	decodedImage, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}

	if decodedImage == nil {
		fmt.Println("Failed to load Image")
	}

	printImageAsASCIIArray(decodedImage)
}

func printImageAsASCIIArray(src image.Image) {
	bounds := src.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	srcAsrgba := image.NewRGBA(src.Bounds())
	draw.Draw(srcAsrgba, bounds, src, bounds.Min, draw.Src)

	for y := 0; y < height; y++ {
		row := ""
		for x := 0; x < width; x++ {
			pixelIndex := (y*width + x) * 4
			pix := srcAsrgba.Pix[pixelIndex : pixelIndex+4]
			r := pix[0]
			g := pix[1]
			b := pix[2]
			luminosity := (r + g + b) / 3
			// Trying to keep image to scale by repeating the Character to widen it
			row += rgbterm.FgString(getCharForLuminosity(luminosity), r, g, b) + rgbterm.FgString(getCharForLuminosity(luminosity), r, g, b)
		}
		fmt.Println(row)
	}
}

func getCharForLuminosity(luminosity uint8) string {
	charArray := strings.Split(grayScale, "")
	scale := float64(len(charArray)) / 255
	luminosityValue := float64(luminosity) * scale
	luminosityIndex := int(math.Floor(luminosityValue))

	return charArray[luminosityIndex]
}
