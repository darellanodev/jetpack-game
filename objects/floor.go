package objects

import (
	_ "image/png"
	"math/rand"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/particles"
	"github.com/hajimehoshi/ebiten/v2"
)

type FloorType int

const (
	FloorNormal FloorType = iota
	FloorLava
	FloorLavaWithDrops
)

const (
	lavaFloorFrameWidth  = 180
	lavaFloorFrameHeight = 53
	lavaFloorFrameSpeed  = 5
	lavadropFloorLightMax = 0.2
	lavadropLightIncrement = 0.03
	lavadropMaxWarningTimes = 5
	maxRandomDropWarningTime = 200
	minRandomDropWarningTime = 200
)

type Floor struct {
	x					 	  		 int
	y					 	  		 int
	FloorType			 	  		 FloorType
	fire		  		 	  		 *particles.ParticlesSystem
	collisionHitBox 	 	  		 *ebiten.Image
	imgFloor1    		 	  		 *ebiten.Image
	imgLavaFloor		 	  		 *ebiten.Image
	imgAnimFire			 	  		 *ebiten.Image
	Lavadrop	 		 	  		 *Lavadrop
	lavadropFloorLight 	 	  		 float64
	lavadropFloorWarning 	  		 bool
	lavadropFloorWarningTimes 		 int
	lavadropTimeToActivateWarning    int
	lavadropMaxTimeToActivateWarning int
	jumpLavadrop					 bool
}

func NewFloor(floorSprites []*ebiten.Image) *Floor {
	
	return &Floor{
		x: 0,
		y: 0,
		collisionHitBox: floorSprites[0],
		imgFloor1: floorSprites[0],
		imgLavaFloor: floorSprites[1],
		imgAnimFire: floorSprites[2],
		Lavadrop: NewLavadrop(floorSprites[3]),
		lavadropFloorLight: 0,
		lavadropFloorWarning: false,
		lavadropTimeToActivateWarning: 0,
		lavadropMaxTimeToActivateWarning: rand.Intn(maxRandomDropWarningTime) + minRandomDropWarningTime,
		jumpLavadrop: false,
	}
}

func (f *Floor) CollisionHitBox() *ebiten.Image {
	return f.collisionHitBox
}

func (f *Floor) Position() (int, int) {
	return f.x, f.y
}

func (f *Floor) InitFloor() {
	f.fire = &particles.ParticlesSystem{
		CurrentSprite: f.imgAnimFire,
		Creating: true,
		PosX: f.x,
		PosY: f.y,
	}
}

func (f *Floor) MoveTo(x int, y int) {
	f.x = x
	f.y = y

	lavadropX := x + lavaFloorFrameWidth / 2 - lavadropWith / 2

	f.Lavadrop.SetInitialPosition(lavadropX, y)
}

func (f *Floor) IsLavaFloor() bool {
	return f.FloorType == FloorLava || f.FloorType == FloorLavaWithDrops
}


func (f *Floor) Update() {

	
	if f.lavadropFloorWarning {

		if (f.lavadropFloorWarningTimes < lavadropMaxWarningTimes) && f.lavadropFloorLight >= lavadropFloorLightMax {
			f.lavadropFloorWarningTimes++
			f.lavadropFloorLight = 0
		}

		if (f.lavadropFloorLight < lavadropFloorLightMax) {
			f.lavadropFloorLight += lavadropLightIncrement
		}

		if (f.lavadropFloorWarningTimes >= lavadropMaxWarningTimes) {
			f.lavadropFloorWarningTimes = 0
			f.lavadropFloorLight = 0
			f.lavadropFloorWarning = false
			f.jumpLavadrop = true
		}

	} else {
		if f.lavadropTimeToActivateWarning >= f.lavadropMaxTimeToActivateWarning {
			f.lavadropFloorWarning = true
			f.lavadropTimeToActivateWarning = 0
			f.lavadropMaxTimeToActivateWarning = rand.Intn(maxRandomDropWarningTime) + minRandomDropWarningTime
		} else {
			f.lavadropTimeToActivateWarning++
		}
	}

	
	
	if f.IsLavaFloor() && f.fire.Creating {
		f.fire.UpdateUp(lavaFloorFrameWidth)
	}

	if f.FloorType == FloorLavaWithDrops && f.jumpLavadrop {
		f.jumpLavadrop = f.Lavadrop.Update()
	}
}

func (f *Floor) drawLavaFloor(screen *ebiten.Image, subImage *ebiten.Image) {
	lib.DrawNormalImage(screen, subImage, f.x, f.y)
}

func (f *Floor) drawLavaFloorDrop(screen *ebiten.Image, subImage *ebiten.Image) {
	lib.DrawLightenImage(screen, subImage, f.x, f.y, f.lavadropFloorLight)
}

func (f *Floor) Draw(screen *ebiten.Image, spriteCount int) {

	var subImage *ebiten.Image
	
	if f.IsLavaFloor() {
		subImage = lib.GetSubImage(f.imgLavaFloor, lavaFloorFrameWidth, lavaFloorFrameHeight, spriteCount, frameCount, lavaFloorFrameSpeed)
	}

	switch f.FloorType {
		case FloorNormal:
			lib.DrawNormalImage(screen, f.imgFloor1, f.x, f.y)
		case FloorLava:
			f.drawLavaFloor(screen, subImage)
		case FloorLavaWithDrops:
			f.Lavadrop.Draw(screen)		
			f.drawLavaFloorDrop(screen, subImage)
	}

	if f.IsLavaFloor() && f.fire.Creating {
		f.fire.Draw(screen)
	}
	
}