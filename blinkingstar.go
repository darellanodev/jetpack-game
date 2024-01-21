package main

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type BlinkingStar struct {
	x				int
	y				int
}

func NewBlinkingStar() *BlinkingStar {
	
	return &BlinkingStar{
		x: 0,
		y: 0,
	}
}

func (bs *BlinkingStar) MoveTo(x int, y int) {
	bs.x = x
	bs.y = y
}

func (bs *BlinkingStar) Draw(screen *ebiten.Image, spriteCount int) {

	subImage := lib.GetSubImage(sprites["blinking_star"], blinkingStarFrameWidth, blinkingStarFrameHeight, spriteCount, frameCount, blinkingStarFrameSpeed)
	lib.DrawNormalImage(screen, subImage, bs.x, bs.y)
	
}

