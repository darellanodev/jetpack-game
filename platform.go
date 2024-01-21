package main

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x				int
	y				int
}

func NewPlatform() *Platform {
	
	return &Platform{
		x: 0,
		y: 0,
	}
}

func (p *Platform) position() (int, int) {
	return p.x, p.y
}


func (p *Platform) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, sprites["platform"], p.x, p.y)
}
