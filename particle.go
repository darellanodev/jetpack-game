package main

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)


type Particle struct {
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

func (pa *Particle) update() {
	if pa.count == 0 {
		return
	}
	pa.count--
}

func (pa *Particle) terminated() bool {
	return pa.count == 0
}

func (pa *Particle) draw(screen *ebiten.Image) {
	if pa.count == 0 {
		return
	}

	x := math.Cos(pa.dir) * float64(pa.maxCount-pa.count)
	y := math.Sin(pa.dir) * float64(pa.maxCount-pa.count)

	op := &ebiten.DrawImageOptions{}

	sx, sy := pa.img.Bounds().Dx(), pa.img.Bounds().Dy()
	op.GeoM.Translate(-float64(sx)/2, -float64(sy)/2)
	op.GeoM.Rotate(pa.angle)
	op.GeoM.Scale(pa.scale, pa.scale)
	op.GeoM.Translate(x + float64(pa.posX), y + float64(pa.posY))

	rate := float32(pa.count) / float32(pa.maxCount)
	var alpha float32
	if rate < 0.2 {
		alpha = rate / 0.2
	} else if rate > 0.8 {
		alpha = (1 - rate) / 0.2
	} else {
		alpha = 1
	}

	op.ColorScale.ScaleAlpha(alpha)

	screen.DrawImage(pa.img, op)
}

func newParticle(img *ebiten.Image, posX int, posY int, life int, sizeMax float32, opaqueMax float32) *Particle {
	c := rand.Intn(50) + life
	dir := rand.Float64() * 2 * math.Pi
	a := rand.Float64() * 2 * math.Pi
	s := rand.Float64() * float64(sizeMax)
	al := rand.Float32() * opaqueMax
	return &Particle{
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