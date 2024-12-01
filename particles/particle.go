package particles

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

func (p *Particle) UpdateRandomDir() {
	if p.count == 0 {
		return
	}
	p.count--

	x := math.Cos(p.dir) * float64(2) + float64(p.posX)
	y := math.Sin(p.dir) * float64(2) + float64(p.posY)

	p.posX = int(x)
	p.posY = int(y)

	rate := float32(p.count) / float32(p.maxCount * 2)
	var alpha float32
	if rate < 0.2 {
		alpha = rate / 0.2
	} else if rate > 0.8 {
		alpha = (1 - rate) / 0.2
	} else {
		alpha = 1
	}

	p.alpha = alpha

}

func (p *Particle) UpdateUpDir() {
	if p.count == 0 {
		return
	}
	p.count--
	
	p.posY = int(p.posY) - 1

	rate := float32(p.count) / float32(p.maxCount * 2)
	var alpha float32
	if rate < 0.2 {
		alpha = rate / 0.2
	} else if rate > 0.8 {
		alpha = (1 - rate) / 0.2
	} else {
		alpha = 1
	}

	p.alpha = alpha

}

func (p *Particle) terminated() bool {
	return p.count == 0
}

func (p *Particle) draw(screen *ebiten.Image) {
	if p.count == 0 {
		return
	}

	op := &ebiten.DrawImageOptions{}

	sx, sy := p.img.Bounds().Dx(), p.img.Bounds().Dy()
	op.GeoM.Translate(-float64(sx)/2, -float64(sy)/2)
	op.GeoM.Rotate(p.angle)
	op.GeoM.Scale(p.scale, p.scale)
	op.GeoM.Translate(float64(p.posX), float64(p.posY))

	op.ColorScale.ScaleAlpha(p.alpha)

	screen.DrawImage(p.img, op)
}

func newParticle(img *ebiten.Image, posX, posY int, sizeMax, opaqueMax float32) *Particle {
	c := rand.Intn(50)
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