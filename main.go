package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/gen2brain/go-fitz"
	"github.com/oliamb/cutter"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("syntax: %s PDF-file [...PDF-file...]\n", os.Args[0])
		return
	}

	for _, arg := range os.Args[1:] {
		process(arg)
	}
}

func process(fn string) {
	fmt.Printf("INFO: Processing %s -> %s.png\n", fn, fn)

	doc, err := fitz.New(fn)
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return
	}

	defer doc.Close()

	img, err := doc.ImageDPI(0, float64(300))
	if err != nil {
		panic(err)
	}

	// Cut actual card
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  970,
		Height: 546,
		Anchor: image.Point{1560, 2700},
		Mode:   cutter.TopLeft, // optional, default value
	})

	f, err := os.Create(fn + ".png")
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return
	}

	err = png.Encode(f, croppedImg)
	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		return
	}

	f.Close()
}
