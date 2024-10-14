package main

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func getImageFromFilePath(filePath string) *ebiten.Image {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		return nil
	}

	return ebiten.NewImageFromImage(image)
}
