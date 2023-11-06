package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fuel struct {
	x				int
	y				int
	currentSprite 	*ebiten.Image
	snaps			bool
}

func (f *Fuel) position() (int, int) {
	return f.x, f.y
}


func (f *Fuel) Draw(screen *ebiten.Image) {

	f.currentSprite = sprites["fuel"]

	op := &ebiten.DrawImageOptions{}
	x, y := f.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(f.currentSprite, op)
}

func (f *Fuel) MoveTo(x int, y int) {
	f.x = x
	f.y = y
}


func (f *Fuel) Update() {
	
	// if f.right{
	// 	f.x += 1
	// }
	// if f.left{
	// 	f.x -= 1
	// }
	// if f.up{
	// 	f.y -= 1
	// }
	// if f.down{
	// 	f.y += 1
	// }
	
}