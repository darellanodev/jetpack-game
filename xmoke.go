package main

import (
	"container/list"
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
	alpha *= xp.alpha
	op.ColorScale.ScaleAlpha(alpha)

	screen.DrawImage(xp.img, op)
}

func newSprite(img *ebiten.Image, posX int, posY int) *Xparticle {
	c := rand.Intn(50) + 300
	dir := rand.Float64() * 2 * math.Pi
	a := rand.Float64() * 2 * math.Pi
	s := rand.Float64()*0.1 + 0.4
	return &Xparticle{
		img: img,

		maxCount: c,
		count:    c,
		dir:      dir,

		angle: a,
		scale: s,
		alpha: 0.5,
		posX: posX,
		posY: posY,
	}
}

type Xmoke struct {
	particles *list.List
	posX	 int
	posY	 int
	creating bool
}

func (x *Xmoke) MoveTo(posX int, posY int) {
	x.posX = posX
	x.posY = posY
}

func (x *Xmoke) Update() error {
	if (x.particles == nil) {
		x.particles = list.New()
	}

	if (x.creating) {
		if x.particles.Len() < 500 && rand.Intn(4) < 3 {
			x.particles.PushBack(newSprite(sprites["smoke"], x.posX, x.posY))
		}
	}

	for e := x.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Xparticle)
		s.update()
		if s.terminated() {
			defer x.particles.Remove(e)
		}
	}
	return nil
}

func (x *Xmoke) Draw(screen *ebiten.Image) {
	for e := x.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Xparticle)
		s.draw(screen)
	}

}
