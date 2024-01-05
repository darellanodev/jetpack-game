package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x				int
	y				int
}

func (p *Platform) position() (int, int) {
	return p.x, p.y
}


func (p *Platform) Draw(screen *ebiten.Image) {

	NewGame().drawNormalImage(screen, sprites["platform"], p.x, p.y)
}
