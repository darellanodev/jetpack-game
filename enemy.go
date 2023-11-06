package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x  int
	y  int
	currentSprite *ebiten.Image
	up bool
	down bool
	left bool
	right bool
}


// Position returns the player p's position.
func (e *Enemy) position() (int, int) {
	return e.x, e.y
}


func (e *Enemy) Draw(screen *ebiten.Image) {

	e.currentSprite = sprites["enemy"]


	op := &ebiten.DrawImageOptions{}
	x, y := e.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(e.currentSprite, op)
}


func (e *Enemy) Update() {

	enemyCenterX := e.x + (enemyWidth / 2)
	enemyCenterY := e.y + (enemyHeight / 2)

	if e.right{
		e.x += enemySpeed
	}
	if e.left{
		e.x -= enemySpeed
	}
	if e.up{
		e.y -= enemySpeed
	}
	if e.down{
		e.y += enemySpeed
	}
	
	if enemyCenterX > 20000 {
		e.left = true
		e.right = false
	}
	if enemyCenterX < 0 {
		e.right = true
		e.left = false
	}
	if enemyCenterY < 0 {
		e.down = true
		e.up = false
	}
	if enemyCenterY > 14400 {
		e.up = true
		e.down = false
	}
	
}