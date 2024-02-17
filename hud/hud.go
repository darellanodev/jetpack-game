package hud

import (
	"image/color"
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Hud struct {
	x					int
	y					int
	oxygen				int
	oxygenTimeToConsume int
	levelTitle			string
	lives				int
	img 				*ebiten.Image	
	liveImg				*ebiten.Image	
	oxygenBar			*Progressbar
}

func NewHud(img *ebiten.Image, liveImg *ebiten.Image) *Hud {
	
	return &Hud{
		x: 0,
		y: 0,
		oxygen: maxOxygenCapacity,
		oxygenTimeToConsume: maxOxygenTimeToConsume,
		img: img,
		liveImg: liveImg,
		oxygenBar: NewProgressbar(375, 33, 600, 33, 100, 120),
	}
}

const (
	maxOxygenCapacity      = 610
	maxOxygenTimeToConsume = 10
)

func (h *Hud) SetMaxOxygenCapacity() {
	h.oxygen = maxOxygenCapacity
}

func (h *Hud) SetTitle(levelTitle string) {
	h.levelTitle = levelTitle
}

func (h *Hud) SetLives(lives int) {
	h.lives = lives
}

func (h *Hud) Update() {

	if (h.oxygenBar.IsNotEmpty()) {
		h.oxygenBar.UpdateDecrease()
	}

}

func (h *Hud) drawTitle(screen *ebiten.Image) {
	text.Draw(screen, "Level " + h.levelTitle, lib.MplusSmallFont, 105, 115, color.Black)
}

func (h *Hud) Draw(screen *ebiten.Image) {

	h.drawBackground(screen)
	h.drawOxigenBar(screen)
	h.drawTitle(screen)
	h.drawLives(screen)
}

func (h *Hud) drawLive(offset int, screen *ebiten.Image) {
	
	lib.DrawNormalImage(screen, h.liveImg, h.x + 80 + offset, h.y + 27)

}

func (h *Hud) drawLives(screen *ebiten.Image) {

	for i := 0; i < h.lives; i++ {
		h.drawLive(0 + (i * 30), screen)
	}

}

func (h *Hud) drawBackground(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, h.img, h.x, h.y)
	
}

func (h *Hud) drawOxigenBar(screen *ebiten.Image) {
	h.oxygenBar.Draw(screen)
	// vector.DrawFilledRect(screen, 375, 33, float32(h.oxygen), 18, color.RGBA{0xff, 0xff, 0xff, 0xff}, true)
}

func (h *Hud) MoveTo(x int, y int) {
	h.x = x
	h.y = y
}
