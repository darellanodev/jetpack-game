package main

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticlesExpansion struct {
	particles 	  *list.List
	posX	 	  int
	posY	 	  int
	creating 	  bool
	currentSprite *ebiten.Image
}

func (psa *ParticlesExpansion) SetImg(img *ebiten.Image) {
	psa.currentSprite = img
}

func (psa *ParticlesExpansion) MoveTo(posX int, posY int) {
	psa.posX = posX
	psa.posY = posY
}

func (psa *ParticlesExpansion) Update() error {
	if (psa.particles == nil) {
		psa.particles = list.New()
	}

	if (psa.creating) {
		if psa.particles.Len() < 200 && rand.Intn(4) < 3 {
			psa.particles.PushBack(newParticle(psa.currentSprite, psa.posX, psa.posY, 100, 0.7, 0.1))
		}
	}

	for e := psa.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.update()
		if s.terminated() {
			defer psa.particles.Remove(e)
		}
	}
	return nil
}

func (psa *ParticlesExpansion) Draw(screen *ebiten.Image) {
	for e := psa.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.draw(screen)
	}

}
