package main

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

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
		if x.particles.Len() < 200 && rand.Intn(4) < 3 {
			x.particles.PushBack(newParticle(sprites["smoke"], x.posX, x.posY, 100, 0.7, 0.1))
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
