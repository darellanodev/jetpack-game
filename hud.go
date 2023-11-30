package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hud struct {
	x				int
	y				int
	currentSprite 	*ebiten.Image
}

func (h *Hud) position() (int, int) {
	return h.x, h.y
}


func (h *Hud) Draw(screen *ebiten.Image) {

	h.currentSprite = sprites["hud"]

	op := &ebiten.DrawImageOptions{}
	x, y := h.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(h.currentSprite, op)
}

func (h *Hud) MoveTo(x int, y int) {
	h.x = x
	h.y = y
}
