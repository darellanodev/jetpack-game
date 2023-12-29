package main

import (
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Hud struct {
	x					int
	y					int
	currentSprite 		*ebiten.Image
	oxygen				int
	oxygenTimeToConsume int
	levelTitle			string
}

func (h *Hud) position() (int, int) {
	return h.x, h.y
}

func (h *Hud) setTitle(levelTitle string) {
	h.levelTitle = levelTitle
}

func (h *Hud) Update() {
	if (h.oxygenTimeToConsume > 0) {
		h.oxygenTimeToConsume--
	} else {
		h.oxygenTimeToConsume = maxOxygenTimeToConsume
		if (h.oxygen > 0) {
			h.oxygen--
		}
	}
}

func (h *Hud) drawTitle(screen *ebiten.Image) {
	text.Draw(screen, "Level " + h.levelTitle, mplusHudFont, 50, 53, color.Black)
}

func (h *Hud) Draw(screen *ebiten.Image) {

	h.drawBackground(screen)
	h.drawOxigenBar(screen)
	h.drawTitle(screen)
}


func (h *Hud) drawBackground(screen *ebiten.Image) {

	h.currentSprite = sprites["hud"]

	op := &ebiten.DrawImageOptions{}
	x, y := h.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(h.currentSprite, op)
}

func (h *Hud) drawOxigenBar(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 167, 14, float32(h.oxygen), 11, color.RGBA{0xff, 0xff, 0xff, 0xff}, true)
}

func (h *Hud) MoveTo(x int, y int) {
	h.x = x
	h.y = y
}
