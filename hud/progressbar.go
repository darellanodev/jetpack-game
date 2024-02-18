package hud

import (
	"image/color"
	_ "image/png"
	"strconv"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Progressbar struct {
	x		   		 int
	y		   		 int
	width	   		 int
	height	   		 int
	percentage 		 int
	delayTime  		 int
	initialDelayTime int
	drawOutsideLine   bool
}

func NewProgressbar(x int, y int, width int, height int, percentage int, delayTime int, drawOutsideLine bool) *Progressbar {
	
	return &Progressbar{
		x: x,
		y: y,
		width: width,
		height: height,
		percentage: percentage,
		delayTime: 0,
		initialDelayTime: delayTime,
		drawOutsideLine: drawOutsideLine,
	}
}

const (
	maxPercentage = 100
	progressBarMargin = 10
	progressBarStokeWidth = 2
)

func (p *Progressbar) UpdateIncrease() {
	if p.delayTime < p.initialDelayTime {
		p.delayTime++
	} else {
		p.delayTime = 0
		if p.IsNotFull() {
			p.percentage++
		}
	}
	
}


func (p *Progressbar) UpdateDecrease() {
	
	if p.delayTime < p.initialDelayTime {
		p.delayTime++
	} else {
		p.delayTime = 0
		if p.IsNotEmpty() {
			p.percentage--
		}
	}
}

func (p *Progressbar) IsFull() bool {
	return p.percentage >= maxPercentage
}

func (p *Progressbar) IsNotEmpty() bool {
	return p.percentage > 0
}

func (p *Progressbar) IsNotFull() bool {
	return p.percentage < maxPercentage
}

func (p *Progressbar) IsEmpty() bool {
	return p.percentage <= 0
}

func (p *Progressbar) drawProgressBarPercentage(screen *ebiten.Image) {
	
	x:= p.x + p.percentage * (p.width / 100) - progressBarMargin - 50
	y:= p.y + p.height / 2 + 2
	
	if p.percentage < 30 {
		x = x + 70
		text.Draw(screen, strconv.Itoa(p.percentage) + "%", lib.MplusSmallFont, x, y, color.White)
	} else {
		text.Draw(screen, strconv.Itoa(p.percentage) + "%", lib.MplusSmallFont, x, y, color.Black)
	}
}

func (p *Progressbar) drawOutside(screen *ebiten.Image) {

	x := float32(p.x) - progressBarMargin
	y := float32(p.y) - progressBarMargin
	width := float32(p.width) + progressBarMargin
	height := float32(p.height) + progressBarMargin

	vector.StrokeRect(screen, x, y, width, height, progressBarStokeWidth, color.White, true)
}

func (p *Progressbar) drawFilledBar(screen *ebiten.Image) {
	
	x := float32(p.x)
	y := float32(p.y)
	width := (float32(p.percentage) * float32(p.width) / 100) - progressBarMargin
	height := float32(p.height) - progressBarMargin

	if width < 0 {
		width = 0
	}
	
	vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{0xff, 0xff, 0xff, 0xff}, true)
}

func (p *Progressbar) drawFilledBarBackground(screen *ebiten.Image) {
	
	x := float32(p.x)
	y := float32(p.y)
	width := float32(p.width) - progressBarMargin
	height := float32(p.height) - progressBarMargin
	
	vector.DrawFilledRect(screen, x, y, width, height, color.RGBA{0x66, 0x66, 0x66, 0x66}, true)
}

func (p *Progressbar) Draw(screen *ebiten.Image) {

	if p.drawOutsideLine {
		p.drawOutside(screen)
	}
	p.drawFilledBarBackground(screen)
	p.drawFilledBar(screen)
	p.drawProgressBarPercentage(screen)
	
}

