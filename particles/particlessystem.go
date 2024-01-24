package particles

import (
	"container/list"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticlesSystem struct {
	PosX	 	  int
	PosY	 	  int
	Creating 	  bool
	CurrentSprite *ebiten.Image
	particles 	  *list.List
}

func NewSmoke(imgSmoke *ebiten.Image) *ParticlesSystem {
	
	return &ParticlesSystem{
		particles: nil,
		PosX: 100,
		PosY: 100,
		Creating: false,
		CurrentSprite: imgSmoke,
	}
}

func NewExplosion(imgExplosion *ebiten.Image) *ParticlesSystem {
	
	return &ParticlesSystem{
		particles: nil,
		PosX: 100,
		PosY: 100,
		Creating: false,
		CurrentSprite: imgExplosion,
	}
}

func (ps *ParticlesSystem) SetImg(img *ebiten.Image) {
	ps.CurrentSprite = img
}

func (ps *ParticlesSystem) MoveTo(PosX int, PosY int) {
	ps.PosX = PosX
	ps.PosY = PosY
}

func (ps *ParticlesSystem) createNewParticles() {
	if (ps.particles == nil) {
		ps.particles = list.New()
	}
	
	if ps.particles.Len() < 200 && rand.Intn(4) < 3 {
		ps.particles.PushBack(newParticle(ps.CurrentSprite, ps.PosX, ps.PosY, 100, 0.7, 0.1))
	}
}

func (ps *ParticlesSystem) createNewParticlesInLine(randomWidth int) {
	if (ps.particles == nil) {
		ps.particles = list.New()
	}
	
	if ps.particles.Len() < 200 && rand.Intn(4) < 3 {
		ps.particles.PushBack(newParticle(ps.CurrentSprite, ps.PosX + rand.Intn(randomWidth) + 4, ps.PosY, 100, 0.7, 0.1))
	}
}

func (ps *ParticlesSystem) UpdateExpanded() error {

	if (ps.Creating) {
		ps.createNewParticles()
	}
	if(ps.particles == nil) {
		return nil
	}

	for e := ps.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.UpdateRandomDir()
		if s.terminated() {
			defer ps.particles.Remove(e)
		}
	}
	return nil
}


func (ps *ParticlesSystem) UpdateUp(randomWidth int) error {

	if (ps.Creating) {
		ps.createNewParticlesInLine(randomWidth)
	}
	if(ps.particles == nil) {
		return nil
	}

	for e := ps.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.UpdateUpDir()
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
