package main

import (
	"flag"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path"
)

var (
	bg      = flag.String("bg", "bg.png", "背景图片")
	pt      = flag.String("pt", "pt.png", "前景图片")
	offsetX = flag.Int("offsetX", 0, "x轴偏移值")
	offsetY = flag.Int("offsetY", 0, "y轴偏移值")
	prefix  = flag.String("prefix", "test_", "文件名前缀")
)

func main() {
	flag.Parse()
	mergeImage(*pt)
}

func mergeImage(file string) {
	bgImageFile, err := os.Open("imge/bg3.png")
	if err != nil {
		panic(err)
	}
	bgImage, err := png.Decode(bgImageFile)
	if err != nil {
		panic(err)
	}
	defer bgImageFile.Close()
	bgImageBound := bgImage.Bounds()

	preImageFile, _ := os.Open("imge/pt.png")
	preImage, _ := png.Decode(preImageFile)
	preImageBound := preImage.Bounds()
	defer preImageFile.Close()

	offset := image.Pt((bgImageBound.Max.X-preImageBound.Max.X)/2+*offsetX, (bgImageBound.Max.Y-preImageBound.Max.Y)/2+*offsetY)
	b := bgImage.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImage, image.ZP, draw.Src)
	draw.Draw(m, preImage.Bounds().Add(offset), preImage, image.ZP, draw.Over)
	imgw, _ := os.Create(*prefix + path.Base(file))
	png.Encode(imgw, m)
	defer imgw.Close()
}
