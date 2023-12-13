package main

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Smoke struct {
	particles *list.List
	posX	 int
	posY	 int
	creating bool
}

func (sm *Smoke) MoveTo(posX int, posY int) {
	sm.posX = posX
	sm.posY = posY
}

func (sm *Smoke) Update() error {
	if (sm.particles == nil) {
		sm.particles = list.New()
	}

	if (sm.creating) {
		if sm.particles.Len() < 200 && rand.Intn(4) < 3 {
			sm.particles.PushBack(newParticle(sprites["smoke"], sm.posX, sm.posY, 100, 0.7, 0.1))
		}
	}

	for e := sm.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.update()
		if s.terminated() {
			defer sm.particles.Remove(e)
		}
	}
	return nil
}

func (sm *Smoke) Draw(screen *ebiten.Image) {
	for e := sm.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.draw(screen)
	}

}
