package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x				int
	y				int
	img 			*ebiten.Image
}

func NewPlatform(img *ebiten.Image) *Platform {
	
	return &Platform{
		x: 0,
		y: 0,
		img: img,
	}
}


func (p *Platform) Position() (int, int) {
	return p.x, p.y
}

func (p *Platform) MoveTo(x int, y int) {
	p.x = x
	p.y = y
}


func (p *Platform) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, p.img, p.x, p.y)
}
