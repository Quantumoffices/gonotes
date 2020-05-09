package utils

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const (
	qrcodeSize = 150 //二维码大小
	offsetX    = 250 //x轴偏移
	offsetY    = 560 //y轴偏移
)

func GenerateQRCode(info string) error {
	// Create the barcode
	qrCode, err := qr.Encode(info, qr.M, qr.Auto)
	if err != nil {
		return err
	}
	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, qrcodeSize, qrcodeSize)
	if err != nil {
		return err
	}
	// create the output file
	file, err := os.Create("qrcode.png")
	if err != nil {
		return err
	}
	defer file.Close()
	// encode the barcode as png
	err = png.Encode(file, qrCode)
	if err != nil {
		return err
	}
	return nil
}

func MixImages(bgImg, frontImg string, mixImgName string) error {
	//背景图
	bgImageFile, err := os.Open(bgImg)
	if err != nil {
		return err
	}
	defer bgImageFile.Close()

	bgImage, err := png.Decode(bgImageFile)
	if err != nil {
		return err
	}
	bgImageBound := bgImage.Bounds()

	preImageFile, _ := os.Open(frontImg)
	preImage, err := png.Decode(preImageFile)
	if err != nil {
		return err
	}
	preImageBound := preImage.Bounds()
	defer preImageFile.Close()

	offset := image.Pt((bgImageBound.Max.X-preImageBound.Max.X)/2+offsetX, (bgImageBound.Max.Y-preImageBound.Max.Y)/2+offsetY)
	b := bgImage.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImage, image.Point{}, draw.Src)
	draw.Draw(m, preImage.Bounds().Add(offset), preImage, image.Point{}, draw.Over)
	imgw, err := os.Create(mixImgName)
	if err != nil {
		return err
	}
	defer imgw.Close()
	err = png.Encode(imgw, m)
	if err != nil {
		return err
	}
	return nil
}
