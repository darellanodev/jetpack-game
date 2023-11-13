package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x				int
	y				int
	currentSprite 	*ebiten.Image
}

func (p *Platform) position() (int, int) {
	return p.x, p.y
}


func (p *Platform) Draw(screen *ebiten.Image) {

	p.currentSprite = sprites["platform"]

	op := &ebiten.DrawImageOptions{}
	x, y := p.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(p.currentSprite, op)
}

func (p *Platform) Update() {
	
	// if p.right{
	// 	p.x += 1
	// }
	// if p.left{
	// 	p.x -= 1
	// }
	// if p.up{
	// 	p.y -= 1
	// }
	// if p.down{
	// 	p.y += 1
	// }
	
}