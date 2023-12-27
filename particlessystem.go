package main

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticlesSystem struct {
	particles 	  *list.List
	posX	 	  int
	posY	 	  int
	creating 	  bool
	currentSprite *ebiten.Image
}

func (ps *ParticlesSystem) SetImg(img *ebiten.Image) {
	ps.currentSprite = img
}

func (ps *ParticlesSystem) MoveTo(posX int, posY int) {
	ps.posX = posX
	ps.posY = posY
}

func (ps *ParticlesSystem) Update() error {

	if (ps.creating) {
		if (ps.particles == nil) {
			ps.particles = list.New()
		}
		
		if ps.particles.Len() < 200 && rand.Intn(4) < 3 {
			ps.particles.PushBack(newParticle(ps.currentSprite, ps.posX, ps.posY, 100, 0.7, 0.1))
		}
	}
	if(ps.particles == nil) {
		return nil
	}

	for e := ps.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.update()
		if s.terminated() {
			defer ps.particles.Remove(e)
		}
	}
	return nil
}

func (ps *ParticlesSystem) Draw(screen *ebiten.Image) {

	if (ps.particles != nil) {
		for e := ps.particles.Front(); e != nil; e = e.Next() {
			s := e.Value.(*Particle)
			s.draw(screen)
		}
	}

}
