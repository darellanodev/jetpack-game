package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fuel struct {
	x				int
	y				int
	finalX			int
	finalY			int
	snaps			bool
	collisionHitBox *ebiten.Image
}

func (f *Fuel) Draw(screen *ebiten.Image) {

	f.drawFuel(screen)
	
	if (f.isFalling()) {
		f.drawParachute(screen)		
	}
	
}

func (f *Fuel) drawFuel(screen *ebiten.Image) {
	NewGame().drawNormalImage(screen, sprites["fuel"], f.x, f.y)
}

func (f *Fuel) drawParachute(screen *ebiten.Image) {
	NewGame().drawNormalImage(screen, sprites["parachute"], f.x - 18, f.y - 42)
}

func (f *Fuel) MoveTo(x int, y int) {
	f.x = x
	f.y = y
}

func (f *Fuel) SetFinalPosition(x int, y int) {
	f.finalX = x
	f.finalY = y
}

func (f *Fuel) isFalling() bool {
	return f.y < f.finalY && !f.snaps
}

func (f *Fuel) Update() {
	if (f.isFalling()) {
		f.y += FallingFuelVelocity
	}
}
