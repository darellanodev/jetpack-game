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
	FloorLavaWithDrops
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
	lavadrop	 	*Lavadrop
}

func NewFloor(floorSprites []*ebiten.Image) *Floor {
	
	return &Floor{
		x: 0,
		y: 0,
		collisionHitBox: floorSprites[0],
		imgFloor1: floorSprites[0],
		imgLavaFloor: floorSprites[1],
		imgAnimFire: floorSprites[2],
		lavadrop: NewLavadrop(floorSprites[3]),
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

	f.lavadrop.MoveTo(lavadropX, y)
}

func (f *Floor) IsLavaFloor() bool {
	return f.FloorType == FloorLava || f.FloorType == FloorLavaWithDrops
}


func (f *Floor) Update() {
	
	if f.IsLavaFloor() && f.fire.Creating {
		f.fire.UpdateUp(lavaFloorFrameWidth)
	}

	if f.FloorType == FloorLavaWithDrops {
		f.lavadrop.Update()
	}
}

func (f *Floor) drawLavaFloor(screen *ebiten.Image, spriteCount int) {
	subImage := lib.GetSubImage(f.imgLavaFloor, lavaFloorFrameWidth, lavaFloorFrameHeight, spriteCount, frameCount, lavaFloorFrameSpeed)
	lib.DrawNormalImage(screen, subImage, f.x, f.y)
}

func (f *Floor) Draw(screen *ebiten.Image, spriteCount int) {

	switch f.FloorType {
		case FloorNormal:
			lib.DrawNormalImage(screen, f.imgFloor1, f.x, f.y)
		case FloorLava:
			f.drawLavaFloor(screen, spriteCount)
		case FloorLavaWithDrops:
			f.lavadrop.Draw(screen)		
			f.drawLavaFloor(screen, spriteCount)
	}

	if f.IsLavaFloor() && f.fire.Creating {
		f.fire.Draw(screen)
	}
	
}