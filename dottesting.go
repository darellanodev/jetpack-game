package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Dottesting struct {
	x  int
	y  int
}


// Position returns the player p's position.
func (e *Dottesting) position() (int, int) {
	return e.x, e.y
}


func (e *Dottesting) Draw(screen *ebiten.Image) {

	imgDot := sprites["dot_testing"]


	op := &ebiten.DrawImageOptions{}
	x, y := e.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(imgDot, op)
}

