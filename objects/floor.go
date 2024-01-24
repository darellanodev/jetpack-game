package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/particles"
	"github.com/hajimehoshi/ebiten/v2"
)

type FloorType int

const (
	FloorNormal FloorType = iota
	FloorLava
)

const (
	lavaFloorFrameWidth  = 180
	lavaFloorFrameHeight = 53
	lavaFloorFrameSpeed  = 5
)

type Floor struct {
	x				int
	y				int
	FloorType		FloorType
	fire		  	*particles.ParticlesSystem
	collisionHitBox *ebiten.Image
	imgFloor1    	*ebiten.Image
	imgLavaFloor	*ebiten.Image
	imgAnimFire		*ebiten.Image
}

func NewFloor(imgFloor1 *ebiten.Image, imgLavaFloor *ebiten.Image, imgAnimFire *ebiten.Image) *Floor {
	
	return &Floor{
		x: 0,
		y: 0,
		collisionHitBox: imgFloor1,
		imgFloor1: imgFloor1,
		imgLavaFloor: imgLavaFloor,
		imgAnimFire: imgAnimFire,
	
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
}


func (f *Floor) Update() {
	if (f.FloorType == FloorLava && f.fire.Creating) {
		f.fire.UpdateUp(lavaFloorFrameWidth)
	}
}

func (f *Floor) Draw(screen *ebiten.Image, spriteCount int) {

	switch f.FloorType {
		case FloorNormal:
			lib.DrawNormalImage(screen,f.imgFloor1,f.x,f.y)
		case FloorLava:
			subImage := lib.GetSubImage(f.imgLavaFloor, lavaFloorFrameWidth, lavaFloorFrameHeight, spriteCount, frameCount, lavaFloorFrameSpeed)
			lib.DrawNormalImage(screen, subImage, f.x, f.y)
	}

	if (f.FloorType == FloorLava && f.fire.Creating) {
		f.fire.Draw(screen)
	}
	
}