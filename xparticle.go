package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)


type Xparticle struct {
	count    int
	maxCount int
	dir      float64

	img   *ebiten.Image
	scale float64
	angle float64
	alpha float32

	posX	int
	posY 	int
}

func (xp *Xparticle) update() {
	if xp.count == 0 {
		return
	}
	xp.count--
}

func (xp *Xparticle) terminated() bool {
	return xp.count == 0
}

func (xp *Xparticle) draw(screen *ebiten.Image) {
	if xp.count == 0 {
		return
	}

	x := math.Cos(xp.dir) * float64(xp.maxCount-xp.count)
	y := math.Sin(xp.dir) * float64(xp.maxCount-xp.count)

	op := &ebiten.DrawImageOptions{}

	sx, sy := xp.img.Bounds().Dx(), xp.img.Bounds().Dy()
	op.GeoM.Translate(-float64(sx)/2, -float64(sy)/2)
	op.GeoM.Rotate(xp.angle)
	op.GeoM.Scale(xp.scale, xp.scale)
	op.GeoM.Translate(x + float64(xp.posX), y + float64(xp.posY))

	rate := float32(xp.count) / float32(xp.maxCount)
	var alpha float32
	if rate < 0.2 {
		alpha = rate / 0.2
	} else if rate > 0.8 {
		alpha = (1 - rate) / 0.2
	} else {
		alpha = 1
	}

	op.ColorScale.ScaleAlpha(alpha)

	screen.DrawImage(xp.img, op)
}

func newParticle(img *ebiten.Image, posX int, posY int, life int, sizeMax float32, opaqueMax float32) *Xparticle {
	c := rand.Intn(50) + life
	dir := rand.Float64() * 2 * math.Pi
	a := rand.Float64() * 2 * math.Pi
	s := rand.Float64() * float64(sizeMax)
	al := rand.Float32() * opaqueMax
	return &Xparticle{
		img: img,

		maxCount: c,
		count:    c,
		dir:      dir,

		angle: a,
		scale: s,
		alpha: al,
		posX: posX,
		posY: posY,
	}
}