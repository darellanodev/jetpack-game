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
	currentSprite 	*ebiten.Image
	snaps			bool
}

func (f *Fuel) position() (int, int) {
	return f.x, f.y
}

func (f *Fuel) Draw(screen *ebiten.Image) {

	f.drawFuel(screen)
	
	if (f.isFalling()) {
		f.drawParachute(screen)		
	}
	
}

func (f *Fuel) drawFuel(screen *ebiten.Image) {
	f.currentSprite = sprites["fuel"]

	op := &ebiten.DrawImageOptions{}
	x, y := f.position()
	
	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(f.currentSprite, op)

}

func (f *Fuel) drawParachute(screen *ebiten.Image) {
	x, y := f.position()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x-250)/unit, float64(y-700)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(sprites["parachute"], op)
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
