package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	generateQRCode()
	mixImages()
}

func generateQRCode() {
	// Create the barcode
	qrCode, _ := qr.Encode("Hello World", qr.M, qr.Auto)
	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, qrcodeSize, qrcodeSize)
	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()
	// encode the barcode as png
	png.Encode(file, qrCode)
}

const (
	qrcodeSize = 120
	offsetX    = 250 //x轴偏移
	offsetY    = 560 //y轴偏移
)

func mixImages() {
	//背景图
	bgImageFile, err := os.Open("bg.png")
	if err != nil {
		panic(err)
	}
	defer bgImageFile.Close()

	bgImage, err := png.Decode(bgImageFile)
	if err != nil {
		panic(err)
	}
	bgImageBound := bgImage.Bounds()

	preImageFile, _ := os.Open("qrcode.png")
	preImage, _ := png.Decode(preImageFile)
	preImageBound := preImage.Bounds()
	defer preImageFile.Close()

	offset := image.Pt((bgImageBound.Max.X-preImageBound.Max.X)/2+offsetX, (bgImageBound.Max.Y-preImageBound.Max.Y)/2+offsetY)
	b := bgImage.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImage, image.Point{}, draw.Src)
	draw.Draw(m, preImage.Bounds().Add(offset), preImage, image.Point{}, draw.Over)
	imgw, _ := os.Create("mix.png")
	png.Encode(imgw, m)
	defer imgw.Close()
}
