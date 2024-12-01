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

func (p *ParticlesSystem) SetImg(img *ebiten.Image) {
	p.CurrentSprite = img
}

func (p *ParticlesSystem) MoveTo(PosX, PosY int) {
	p.PosX = PosX
	p.PosY = PosY
}

func (p *ParticlesSystem) createNewParticles() {
	if p.particles == nil {
		p.particles = list.New()
	}
	
	if p.particles.Len() < 200 && rand.Intn(4) < 3 {
		p.particles.PushBack(newParticle(p.CurrentSprite, p.PosX, p.PosY, 0.7, 0.1))
	}
}

func (p *ParticlesSystem) createNewParticlesInLine(randomWidth int) {
	if p.particles == nil {
		p.particles = list.New()
	}
	
	if p.particles.Len() < 200 && rand.Intn(4) < 3 {
		p.particles.PushBack(newParticle(p.CurrentSprite, p.PosX + rand.Intn(randomWidth) + 4, p.PosY, 0.7, 0.1))
	}
}

func (p *ParticlesSystem) UpdateExpanded() error {

	if p.Creating {
		p.createNewParticles()
	}
	if(p.particles == nil) {
		return nil
	}

	for e := p.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.UpdateRandomDir()
		if s.terminated() {
			defer p.particles.Remove(e)
		}
	}
	return nil
}


func (p *ParticlesSystem) UpdateUp(randomWidth int) error {

	if p.Creating {
		p.createNewParticlesInLine(randomWidth)
	}
	if(p.particles == nil) {
		return nil
	}

	for e := p.particles.Front(); e != nil; e = e.Next() {
		s := e.Value.(*Particle)
		s.UpdateUpDir()
		if s.terminated() {
			defer p.particles.Remove(e)
		}
	}
	return nil
}

func (p *ParticlesSystem) Draw(screen *ebiten.Image) {

	if p.particles != nil {
		for e := p.particles.Front(); e != nil; e = e.Next() {
			s := e.Value.(*Particle)
			s.draw(screen)
		}
	}

}
