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

	op.GeoM.Translate(float64(x), float64(y))
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(p.currentSprite, op)
}
