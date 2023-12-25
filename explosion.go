package main

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Explosion struct {
	particles *list.List
	posX	 int
	posY	 int
	creating bool
}

func (ex *Explosion) MoveTo(posX int, posY int) {
	ex.posX = posX
	ex.posY = posY
}

func (ex *Explosion) Update() error {
	if (ex.particles == nil) {
		ex.particles = list.New()
	}

	if (ex.creating) {
		if ex.particles.Len() < 200 && rand.Intn(4) < 3 {
			ex.particles.PushBack(newParticle(sprites["explosion"], ex.posX, ex.posY, 100, 0.7, 0.1))
		}
	}

	for e := ex.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.update()
		if s.terminated() {
			defer ex.particles.Remove(e)
		}
	}
	return nil
}

func (ex *Explosion) Draw(screen *ebiten.Image) {
	for e := ex.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.draw(screen)
	}

}
