package objects

import (
	_ "image/png"
	"math/rand"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Fuel struct {
	x				int
	y				int
	finalX			int
	finalY			int
	Snaps			bool
	collisionHitBox *ebiten.Image
	imgFuel			*ebiten.Image
	imgParachute	*ebiten.Image
}

const (
	startFuelX            = 0
	startFuelY            = 0
	FallingFuelVelocity   = 2
	offsetFuelLandingY    = 29
	minOffsetFuelLandingX = 20
)

func NewFuel(imgFuel *ebiten.Image, imgParachute *ebiten.Image) *Fuel {
	
	return &Fuel{
		x: 				 startFuelX,
		y: 				 startFuelY,
		Snaps: 			 false,
		collisionHitBox: imgFuel,
		imgFuel: 		 imgFuel,
		imgParachute:    imgParachute,
	}
}

func (f *Fuel) CollisionHitBox() *ebiten.Image {
	return f.collisionHitBox
}

func (f *Fuel) Position() (int, int) {
	return f.x, f.y
}

func (f *Fuel) Draw(screen *ebiten.Image) {

	f.drawFuel(screen)
	
	if f.isFalling() {
		f.drawParachute(screen)		
	}
	
}

func (f *Fuel) drawFuel(screen *ebiten.Image) {
	lib.DrawNormalImage(screen, f.imgFuel, f.x, f.y)
}

func (f *Fuel) drawParachute(screen *ebiten.Image) {
	lib.DrawNormalImage(screen, f.imgParachute, f.x - 18, f.y - 42)
}

func (f *Fuel) MoveTo(x, y int) {
	f.x = x
	f.y = y
}

func (f *Fuel) SetFinalPosition(x, y int) {
	f.finalX = x
	f.finalY = y
}

func (f *Fuel) isFalling() bool {
	return f.y < f.finalY && !f.Snaps
}

func (f *Fuel) Update() {
	if f.isFalling() {
		f.y += FallingFuelVelocity
	}
}

func (f* Fuel) SetFinalPositionIntoPlatform (platformPosX, platformPosY, PlatformWidth int) {
	randX := rand.Intn(PlatformWidth)
	if randX < 20 {
		randX = minOffsetFuelLandingX
	}

	fx := platformPosX + randX
	fy := platformPosY - offsetFuelLandingY

	f.SetFinalPosition(fx, fy)
	f.MoveTo(fx, 0)
}
