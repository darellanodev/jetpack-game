package objects

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Blackfader struct {
	active				bool
	gameStatus			int
	increasing			bool
	isMaxOpaque			bool
	alpha				uint8
	appHeight			int
	appWidth			int
	overlay				*ebiten.Image
}

const (
	increment = 5 
	topAlpha = 255
	minAlpha = 0
)


func NewBlackfader(gameStatus, appWidth, appHeight int) *Blackfader {
	return &Blackfader{
		alpha:     			0,
		gameStatus:		    gameStatus,
		overlay : ebiten.NewImage(appWidth, appHeight),
	}
}

func (b *Blackfader) Activate(gameStatus int) {
	b.active = true
	b.increasing = true
	b.alpha = 0
	b.gameStatus = gameStatus
	b.isMaxOpaque = false
}

func (b *Blackfader) IsActive() bool {
	return b.active
}

func (b *Blackfader) IsMaxOpaque() bool {
	return b.isMaxOpaque
}

func (b *Blackfader) GameStatus() int {
	return b.gameStatus
}

func (b *Blackfader) Draw(screen *ebiten.Image) {
	if (b.active) {
		b.overlay.Fill(color.RGBA{0, 0, 0, b.alpha}) 
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(b.overlay, op)
	}
}

func (b *Blackfader) Update() {
	if (b.active) {
		if (b.increasing) {
			b.alpha += increment
			if (b.alpha >= topAlpha) {
				b.increasing = false
				b.isMaxOpaque = true
			}
		} else {
			b.alpha -= increment
			if (b.alpha <= minAlpha) {
				b.increasing = false
				b.active = false
			}
		}
	}
}