package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rocket struct {
	landedY				int
	landingSpeed		float32
	x					int
	y					int
	snaps				bool
	fuelIndicatorItems	int
	collisionHitBox		*ebiten.Image
}

func NewRocket() *Rocket {
	
	return &Rocket{
		x: 					startRocketX,
		y: 					startRocketY,
		landedY:			landedRocketY,
		landingSpeed: 		rocketMaxSpeed,
		snaps: 				false,
		fuelIndicatorItems: startRocketFuelItems,
	}
}

func (r *Rocket) CollisionHitBox() *ebiten.Image {
	return r.collisionHitBox
}

func (r *Rocket) Position() (int, int) {
	return r.x, r.y
}

func (r *Rocket) restartFuelItems() {
	r.fuelIndicatorItems = 0
}

func (r *Rocket) drawFire(screen *ebiten.Image) {

	drawNormalImage(screen, sprites["fire_center"], r.x + 17, r.y + 120)
}

func (r *Rocket) drawIndicators(screen *ebiten.Image) {
		
	for i := 0; i < 5; i++ {
		if (i < r.fuelIndicatorItems){
			drawNormalImage(screen, sprites["rocket_fuel_indicator_on"], r.x + 17, r.y - (8 * i) + 80)
		} else {
			drawNormalImage(screen, sprites["rocket_fuel_indicator_off"], r.x + 17, r.y - (8 * i) + 80)
		}
	}
}

func (r *Rocket) Draw(screen *ebiten.Image) {

	drawNormalImage(screen, sprites["rocket"], r.x, r.y)	
	
	if (r.y < r.landedY) {
		r.drawFire(screen)
	}
	r.drawIndicators(screen)
	

}

func (r *Rocket) MoveTo(x int, y int) {
	r.x = x
	r.y = y
}
