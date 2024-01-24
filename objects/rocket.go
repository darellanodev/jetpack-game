package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Rocket struct {
	LandedY					  int
	LandingSpeed			  float32
	x						  int
	y						  int
	snaps					  bool
	FuelIndicatorItems		  int
	collisionHitBox			  *ebiten.Image
	imgRocketFuelIndicatorOn  *ebiten.Image
	imgRocketFuelIndicatorOff *ebiten.Image
	imgRocket 				  *ebiten.Image
	imgFireCenter			  *ebiten.Image
}

const (
	StartRocketX         = 750
	StartRocketY         = -31
	RocketMaxSpeed       = 5
	RocketWidth  		 = 64
	RocketHeight 		 = 128
	landedRocketY        = 597
	startRocketFuelItems = 0
	rocketAcceleration   = 0.032
)

func NewRocket(imgFireCenter *ebiten.Image, imgRocketFuelIndicatorOn *ebiten.Image, imgRocketFuelIndicatorOff *ebiten.Image, imgRocket *ebiten.Image) *Rocket {
	
	return &Rocket{
		x: 						   StartRocketX,
		y: 						   StartRocketY,
		LandedY:				   landedRocketY,
		LandingSpeed: 			   RocketMaxSpeed,
		snaps: 					   false,
		FuelIndicatorItems: 	   startRocketFuelItems,
		collisionHitBox:		   imgRocket,
		imgFireCenter: 			   imgFireCenter,
		imgRocketFuelIndicatorOn:  imgRocketFuelIndicatorOn,
		imgRocketFuelIndicatorOff: imgRocketFuelIndicatorOff,
		imgRocket: 				   imgRocket,
		
	}
}

func (r *Rocket) CollisionHitBox() *ebiten.Image {
	return r.collisionHitBox
}

func (r *Rocket) Position() (int, int) {
	return r.x, r.y
}

func (r *Rocket) GetY() int {
	return r.y
}

func (r *Rocket) GetX() int {
	return r.x
}

func (r *Rocket) RestartFuelItems() {
	r.FuelIndicatorItems = 0
}

func (r *Rocket) drawFire(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, r.imgFireCenter, r.x + 17, r.y + 120)
}

func (r *Rocket) drawIndicators(screen *ebiten.Image) {
		
	for i := 0; i < 5; i++ {
		if (i < r.FuelIndicatorItems){
			lib.DrawNormalImage(screen, r.imgRocketFuelIndicatorOn, r.x + 17, r.y - (8 * i) + 80)
		} else {
			lib.DrawNormalImage(screen, r.imgRocketFuelIndicatorOff, r.x + 17, r.y - (8 * i) + 80)
		}
	}
}

func (r *Rocket) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, r.imgRocket, r.x, r.y)	
	
	if (r.y < r.LandedY) {
		r.drawFire(screen)
	}
	r.drawIndicators(screen)
	

}

func (r *Rocket) MoveTo(x int, y int) {
	r.x = x
	r.y = y
}

func (r* Rocket) Landing() {
	r.MoveTo(r.GetX(), r.GetY() + (2)*int(r.LandingSpeed))
	r.LandingSpeed -= rocketAcceleration

}

func (r* Rocket) TakeOff() {
	r.MoveTo(r.x, r.y-(2)*int(r.LandingSpeed))
	r.LandingSpeed += rocketAcceleration
}