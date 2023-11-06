package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemytesting struct {
	x  int
	y  int
	up bool
	down bool
	left bool
	right bool
}


// Position returns the player p's position.
func (e *Enemytesting) position() (int, int) {
	return e.x, e.y
}


func (e *Enemytesting) Draw(screen *ebiten.Image) {

	imgEnemy := sprites["enemy_testing"]


	op := &ebiten.DrawImageOptions{}
	x, y := e.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(imgEnemy, op)
}

