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
	currentSprite 		*ebiten.Image
	snaps				bool
	fuelIndicatorItems	int
	fuelIndicatorOn 	*ebiten.Image
	fuelIndicatorOff	*ebiten.Image
}

func (r *Rocket) position() (int, int) {
	return r.x, r.y
}

func (r *Rocket) restartFuelItems() {
	r.fuelIndicatorItems = 0
}

func (r *Rocket) drawFire(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	x, y := r.position()

	op.GeoM.Translate(float64(x)/unit + 17, float64(y)/unit + 120)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(sprites["fire_center"], op)

}

func (r *Rocket) drawIndicators(screen *ebiten.Image) {
		
	op := &ebiten.DrawImageOptions{}
	x, y := r.position()
	
	op.GeoM.Translate(float64(x)/unit + 17, float64(y)/unit + 80)
	op.GeoM.Scale(scale, scale)
	for i := 0; i < 5; i++ {
		op.GeoM.Translate(0, -5)
		if (i < r.fuelIndicatorItems){
			screen.DrawImage(r.fuelIndicatorOn, op)
			
		} else {
			
			screen.DrawImage(r.fuelIndicatorOff, op)
		}
	}
}

func (r *Rocket) Draw(screen *ebiten.Image) {

	r.currentSprite = sprites["rocket"]
	r.fuelIndicatorOn = sprites["rocket_fuel_indicator_on"]
	r.fuelIndicatorOff = sprites["rocket_fuel_indicator_off"]

	op := &ebiten.DrawImageOptions{}
	x, y := r.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(r.currentSprite, op)

	
	if (r.y < r.landedY) {
		r.drawFire(screen)
	}
	r.drawIndicators(screen)
	

}

func (r *Rocket) MoveTo(x int, y int) {
	r.x = x
	r.y = y
}
