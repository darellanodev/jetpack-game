package main

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type BlinkingStar struct {
	x				int
	y				int
}

func (bs *BlinkingStar) MoveTo(x int, y int) {
	bs.x = x
	bs.y = y
}

func (bs *BlinkingStar) Draw(screen *ebiten.Image, spriteCount int) {

	i := (spriteCount / 5) % frameCount
	sx, sy := frameOX+i*blinkingStarFrameWidth, frameOY

	subImage := sprites["blinking_star"].SubImage(image.Rect(sx, sy, sx+blinkingStarFrameWidth, sy+blinkingStarFrameHeight)).(*ebiten.Image)
	drawNormalImage(screen, subImage, bs.x, bs.y)
	
}