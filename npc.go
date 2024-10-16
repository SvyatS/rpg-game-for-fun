package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type EventEnum struct {
	Idle     uint8
	Walk     uint8
	Run      uint8
	Jump     uint8
	Attack_1 uint8
	Attack_2 uint8
	Flame    uint8
	Fireball uint8
	Dead     uint8
}

var eventEnum = EventEnum{
	Idle:     0,
	Walk:     1,
	Run:      2,
	Jump:     3,
	Attack_1: 4,
	Attack_2: 5,
	Flame:    6,
	Fireball: 7,
	Dead:     8,
}

var rejectedEvent = []int{
	0, 1, 2,
}

type Animation struct {
	img         *ebiten.Image
	frameCounts int
}

type Hero struct {
	currentEvent          uint8
	currentAnimationFrame int
	events                [9]Animation
	position              [2]int
	prevPosition          [2]int
	health                uint8
	mana                  uint8
	moveSpeed             int
	attackSpeed           int
	jumpHeight            int
	moveEvent             int
	eventAnimating        bool
	rightSide             bool
	attacking             bool
}

func (h *Hero) init() {
	// TODO: add sprite width height
	h.currentEvent = 0
	h.currentAnimationFrame = 0
	h.position[0], h.position[1] = 200, 340
	h.health = 100
	h.mana = 100
	h.moveSpeed = 60
	h.attackSpeed = 17
	h.jumpHeight = 30
	h.moveEvent = 2
	h.eventAnimating = false
	h.rightSide = true
	h.attacking = false
	h.events = [9]Animation{
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Idle.png"),
			frameCounts: 7,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Walk.png"),
			frameCounts: 6,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Run.png"),
			frameCounts: 8,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Jump.png"),
			frameCounts: 9,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Attack_1.png"),
			frameCounts: 4,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Attack_2.png"),
			frameCounts: 4,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Flame_jet.png"),
			frameCounts: 14,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Fireball.png"),
			frameCounts: 8,
		},
		Animation{
			img:         getImageFromFilePath("src/assets/Fire vizard/Dead.png"),
			frameCounts: 6,
		},
	}
}

func (h *Hero) moveRecognize(keys []ebiten.Key) int {
	var side int = 0
	var up int = 0

	if slices.Contains(keys, ebiten.KeyA) {
		side--
	}
	if slices.Contains(keys, ebiten.KeyD) {
		side++
	}
	if slices.Contains(keys, ebiten.KeyShiftLeft) {
		side = side * 2
	}
	if slices.Contains(keys, ebiten.KeySpace) {
		up++
	}

	side = side + 2
	return side + up*5
}

// func (h *Hero) attckRecognize(keys []ebiten.Key) int {
// 	if slices.Contains(keys, ebiten.mo) {
// 		side--
// 	}
// }

func (h *Hero) doMove(moveEventIdx int) {
	switch moveEventIdx {
	case 0:
		h.runleft()
	case 1:
		h.walkLeft()
	case 2:
		h.step()
	case 3:
		h.walkRight()
	case 4:
		h.runRight()
	case 5:
		h.jumpRunLeft()
	case 6:
		h.jumpWalkLeft()
	case 7:
		h.jumpUp()
	case 8:
		h.jumpWalkRight()
	case 9:
		h.jumpRunRight()
	}
}

func (h *Hero) updateFrame(tick int) {
	if tick%(maxFps/(h.moveSpeed)*12) == 0 {
		if !h.attacking {
			h.doMove(h.moveEvent)
		}

		fmt.Println(h.currentAnimationFrame)
		h.currentAnimationFrame++
	}

	if h.currentAnimationFrame >= h.events[h.currentEvent].frameCounts {
		h.currentAnimationFrame = 0
		h.eventAnimating = false
		if h.attacking {
			h.attacking = false
		}
	}
}

var eventForAnimation = [10]uint8{
	2, 1, 0, 1, 2, 3, 3, 3, 3, 3,
}

func (h *Hero) update(keys []ebiten.Key, tick int) {
	if !h.eventAnimating || slices.Contains(rejectedEvent, int(h.currentEvent)) {
		h.moveEvent = h.moveRecognize(keys)
		if h.currentEvent != eventForAnimation[h.moveEvent] {
			h.currentAnimationFrame = 0
			h.prevPosition = h.position
		}
	}
	h.attack()
	h.eventAnimating = true

	h.updateFrame(tick)
}

func (h *Hero) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if h.rightSide {
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(float64(h.position[0]), float64(h.position[1]))
	} else {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(h.position[0]+110), float64(h.position[1]))
	}

	sx := h.currentAnimationFrame * 128
	screen.DrawImage(h.events[h.currentEvent].img.SubImage(image.Rect(sx, 0, sx+128, 128)).(*ebiten.Image), op)
}

func (h *Hero) attack() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		h.currentEvent = eventEnum.Attack_1
		h.currentAnimationFrame = 0
		h.attacking = true
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButton1) {
		h.currentEvent = eventEnum.Attack_2
		h.currentAnimationFrame = 0
		h.attacking = true
	} else if ebiten.IsKeyPressed(ebiten.KeyQ) {
		h.currentEvent = eventEnum.Fireball
		h.currentAnimationFrame = 0
		h.attacking = true
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		h.currentEvent = eventEnum.Flame
		h.currentAnimationFrame = 0
		h.attacking = true
	}
}
