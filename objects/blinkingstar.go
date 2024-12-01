package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type BlinkingStar struct {
	x				int
	y				int
	img				*ebiten.Image
}

const (
	blinkingStarFrameWidth     = 16
	blinkingStarFrameHeight    = 16
	blinkingStarFrameSpeed     = 5
	ChangeBlinkingStarsMaxTime = 50
)

func NewBlinkingStar(img *ebiten.Image) *BlinkingStar {
	
	return &BlinkingStar{
		x: 0,
		y: 0,
		img: img,
	}
}

func (b *BlinkingStar) MoveTo(x, y int) {
	b.x = x
	b.y = y
}

func (b *BlinkingStar) Draw(screen *ebiten.Image, spriteCount int) {

	subImage := lib.GetSubImage(b.img, blinkingStarFrameWidth, blinkingStarFrameHeight, spriteCount, frameCount, blinkingStarFrameSpeed)
	lib.DrawNormalImage(screen, subImage, b.x, b.y)
	
}

