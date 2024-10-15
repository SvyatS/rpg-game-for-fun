package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	currentOffset int
	img           *ebiten.Image
	width         int
	heigth        int
}

func (b *Background) init() {
	b.currentOffset = 0
	b.width = 576
	b.heigth = 324
	b.img = getImageFromFilePath("src/assets/background 1/background 1.png")
}

func (b *Background) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(screenWidth)/float64(b.width), float64(screenHeight)/float64(b.heigth))
	op.GeoM.Translate(0, 0)
	screen.DrawImage(b.img.SubImage(image.Rect(0, 0, b.width, b.heigth)).(*ebiten.Image), op)
}
