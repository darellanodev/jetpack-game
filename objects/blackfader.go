package objects

import (
	_ "image/png"
	"math"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Blackfader struct {
	active				bool
	increasing			bool
	alpha				float32
	imgBackground       *ebiten.Image
}

const (
	increment = 0.02
)


func NewBlackfader(BlackfaderSprites []*ebiten.Image) *Blackfader {
	return &Blackfader{
		alpha:     			0.0,
		imgBackground:      BlackfaderSprites[0],
	}
}

func (b *Blackfader) Activate() {
	b.active = true
	b.increasing = true
	b.alpha = 0.0
}

func (b *Blackfader) IsActive() bool {
	return b.active
}

func (b *Blackfader) Draw(screen *ebiten.Image) {
	if (b.active) {
		lib.DrawAlphaImage(screen, b.imgBackground, 0, 0, b.alpha)
	}
}

func (b *Blackfader) Update() {
	if (b.active) {
		if (b.increasing) {
			b.alpha += increment
			if (b.alpha >= 1) {
				b.increasing = false
				b.alpha = 1
			}
		} else {
			b.alpha -= increment
			if (b.alpha <= 0) {
				b.increasing = false
				b.alpha = 0.0
				b.active = false
			}
		}
		b.alpha = float32(math.Round(float64(b.alpha) * 100) / 100)

		if b.alpha < 1e-6 {
			b.alpha = 0
		}
	}
}