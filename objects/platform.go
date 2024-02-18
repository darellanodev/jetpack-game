package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	x				int
	y				int
	platformImg		*ebiten.Image
	pillarImg		*ebiten.Image
}

const(
	platformWidth = 166
	pillarOffsetX = platformWidth/2 - 10
	pillarHeight = 75
)

func NewPlatform(platformSprites []*ebiten.Image) *Platform {
	
	return &Platform{
		x: 0,
		y: 0,
		platformImg: platformSprites[0],
		pillarImg: platformSprites[1],
	}
}


func (p *Platform) Position() (int, int) {
	return p.x, p.y
}

func (p *Platform) MoveTo(x int, y int) {
	p.x = x
	p.y = y
}

func (p *Platform) drawPillar(screen *ebiten.Image) {
	for i := 0; i < 10; i++ {
		lib.DrawNormalImage(screen, p.pillarImg, p.x + pillarOffsetX, p.y + pillarHeight * i)
	}
}

func (p *Platform) Draw(screen *ebiten.Image) {

	p.drawPillar(screen)
	lib.DrawNormalImage(screen, p.platformImg, p.x, p.y)
}
