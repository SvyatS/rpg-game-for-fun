package main

import (
	// "image"

	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth    = 320 * 2
	screenHeight   = 240 * 2
	maxFps         = 60
	frameOX        = 0
	frameOY        = 128
	frameWidth     = 128
	frameHeight    = 128
	frameCount     = 7
	animationSpeed = 10
)

type Game struct {
	count      int
	hero       Hero
	background Background
	keys       []ebiten.Key
}

func (g *Game) init() {
	// g.hero.init()
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.count = g.count%maxFps + 1
	g.hero.update(g.keys, g.count)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.draw(screen)
	g.hero.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	var hero Hero
	var background Background
	hero.init()
	background.init()
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{hero: hero, background: background}); err != nil {
		log.Fatal(err)
	}
}
