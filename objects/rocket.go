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
	scaleX					  float32
	scaleY					  float32
	snaps					  bool
	FuelIndicatorItems		  int
	allwaysShowFire			  bool
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

func NewRocket(rocketSprites []*ebiten.Image) *Rocket {
	
	return &Rocket{
		x: 						   StartRocketX,
		y: 						   StartRocketY,
		scaleX: 				   1,
		scaleY: 				   1,
		LandedY:				   landedRocketY,
		LandingSpeed: 			   RocketMaxSpeed,
		snaps: 					   false,
		allwaysShowFire: 		   false,
		FuelIndicatorItems: 	   startRocketFuelItems,
		collisionHitBox:		   rocketSprites[3],
		imgFireCenter: 			   rocketSprites[0],
		imgRocketFuelIndicatorOn:  rocketSprites[1],
		imgRocketFuelIndicatorOff: rocketSprites[2],
		imgRocket: 				   rocketSprites[3],
		
	}
}

func (r *Rocket) ReduceScale() {
	r.scaleX -= 0.001
	r.scaleY -= 0.001
}

func (r *Rocket) SetFireAllways() {
	r.allwaysShowFire = true
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
	posX := int(float32(float32(float32(r.x) + 17 * r.scaleX) / r.scaleX) * r.scaleX)
	posY := int(float32(float32(float32(r.y) + 120 * r.scaleY) / r.scaleY) * r.scaleY)

	lib.DrawNormalScaledImage(screen, r.imgFireCenter, posX, posY, r.scaleX, r.scaleY)
}

func (r *Rocket) drawIndicators(screen *ebiten.Image) {
		
	var posX int
	var posY int
	for i := 0; i < 5; i++ {
	posX = int(float32(float32(float32(r.x) + 17 * r.scaleX) / r.scaleX) * r.scaleX)
	posY = int(float32(float32(float32(r.y) - float32(8 * i) + 80 * r.scaleY) / r.scaleY) * r.scaleY)
		if i < r.FuelIndicatorItems {
			lib.DrawNormalScaledImage(screen, r.imgRocketFuelIndicatorOn, posX, posY, r.scaleX, r.scaleY)
		} else {
			lib.DrawNormalScaledImage(screen, r.imgRocketFuelIndicatorOff, posX, posY, r.scaleX, r.scaleY)
		}
	}
}

func (r *Rocket) Draw(screen *ebiten.Image) {

	lib.DrawNormalScaledImage(screen, r.imgRocket, r.x, r.y, r.scaleX, r.scaleY)	
	
	if r.y < r.LandedY || r.allwaysShowFire {
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