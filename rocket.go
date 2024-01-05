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
}

func (r *Rocket) restartFuelItems() {
	r.fuelIndicatorItems = 0
}

func (r *Rocket) drawFire(screen *ebiten.Image) {

	NewGame().drawNormalImage(screen, sprites["fire_center"], r.x + 17, r.y + 120)
}

func (r *Rocket) drawIndicators(screen *ebiten.Image) {
		
	for i := 0; i < 5; i++ {
		if (i < r.fuelIndicatorItems){
			NewGame().drawNormalImage(screen, sprites["rocket_fuel_indicator_on"], r.x + 17, r.y - (8 * i) + 80)
		} else {
			NewGame().drawNormalImage(screen, sprites["rocket_fuel_indicator_off"], r.x + 17, r.y - (8 * i) + 80)
		}
	}
}

func (r *Rocket) Draw(screen *ebiten.Image) {

	NewGame().drawNormalImage(screen, sprites["rocket"], r.x, r.y)	
	
	if (r.y < r.landedY) {
		r.drawFire(screen)
	}
	r.drawIndicators(screen)
	

}

func (r *Rocket) MoveTo(x int, y int) {
	r.x = x
	r.y = y
}
