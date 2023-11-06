package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rocket struct {
	x				int
	y				int
	currentSprite 	*ebiten.Image
	snaps			bool
}

func (r *Rocket) position() (int, int) {
	return r.x, r.y
}


func (r *Rocket) Draw(screen *ebiten.Image) {

	r.currentSprite = sprites["rocket"]

	op := &ebiten.DrawImageOptions{}
	x, y := r.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(r.currentSprite, op)
}

func (r *Rocket) MoveTo(x int, y int) {
	r.x = x
	r.y = y
}


func (r *Rocket) Update() {
	
	// if r.right{
	// 	r.x += 1
	// }
	// if r.left{
	// 	r.x -= 1
	// }
	// if r.up{
	// 	r.y -= 1
	// }
	// if r.down{
	// 	r.y += 1
	// }
	
}