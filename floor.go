package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Floor struct {
	x				int
	y				int
	currentSprite 	*ebiten.Image
}

func (f *Floor) position() (int, int) {
	return f.x, f.y
}

func (f *Floor) MoveTo(x int, y int) {
	f.x = x
	f.y = y
}


func (f *Floor) Draw(screen *ebiten.Image) {

	f.currentSprite = sprites["floor1"]

	op := &ebiten.DrawImageOptions{}
	x, y := f.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(f.currentSprite, op)
}
