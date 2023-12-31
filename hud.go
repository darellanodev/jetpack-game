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
	oxygen				int
	oxygenTimeToConsume int
	levelTitle			string
	lives			int
}

func (h *Hud) setTitle(levelTitle string) {
	h.levelTitle = levelTitle
}

func (h *Hud) setLives(lives int) {
	h.lives = lives
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
	text.Draw(screen, "Level " + h.levelTitle, mplusSmallFont, 105, 115, color.Black)
}

func (h *Hud) Draw(screen *ebiten.Image) {

	h.drawBackground(screen)
	h.drawOxigenBar(screen)
	h.drawTitle(screen)
	h.drawLives(screen)
}

func (h *Hud) drawLive(offset int, screen *ebiten.Image) {
	
	NewGame().drawNormalImage(screen, sprites["live"], h.x + 80 + offset, h.y + 27)

}

func (h *Hud) drawLives(screen *ebiten.Image) {

	for i := 0; i < h.lives; i++ {
		h.drawLive(0 + (i * 30), screen)
	}

}

func (h *Hud) drawBackground(screen *ebiten.Image) {

	NewGame().drawNormalImage(screen, sprites["hud"], h.x, h.y)
	
}

func (h *Hud) drawOxigenBar(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 375, 33, float32(h.oxygen), 18, color.RGBA{0xff, 0xff, 0xff, 0xff}, true)
}

func (h *Hud) MoveTo(x int, y int) {
	h.x = x
	h.y = y
}
