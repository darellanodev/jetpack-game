package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type FloorType int

const (
	FloorNormal FloorType = iota
	FloorLava
)

type Floor struct {
	x				int
	y				int
	floorType		FloorType
	fire		  	*ParticlesSystem
	collisionHitBox *ebiten.Image
}

func NewFloor() *Floor {
	
	return &Floor{
		x: 0,
		y: 0,
		collisionHitBox: sprites["floor1"],
	}
}

func (f *Floor) CollisionHitBox() *ebiten.Image {
	return f.collisionHitBox
}

func (f *Floor) Position() (int, int) {
	return f.x, f.y
}

func (f *Floor) InitFloor() {
	f.fire = &ParticlesSystem{
		currentSprite: sprites["fire"],
		creating: true,
		posX: f.x,
		posY: f.y,
	}
}

func (f *Floor) MoveTo(x int, y int) {
	f.x = x
	f.y = y
}

func (f *Floor) SetLavaType() {
	f.floorType = FloorLava
}

func (f *Floor) SetNormalType() {
	f.floorType = FloorNormal
}

func (f *Floor) Update() {
	if (f.floorType == FloorLava && f.fire.creating) {
		f.fire.UpdateUp(lavaFloorFrameWidth)
	}
}

func (f *Floor) Draw(screen *ebiten.Image, spriteCount int) {

	switch f.floorType {
		case FloorNormal:
			drawNormalImage(screen,sprites["floor1"],f.x,f.y)
		case FloorLava:
			subImage := getSubImage(sprites["lava_floor"], lavaFloorFrameWidth, lavaFloorFrameHeight, spriteCount, frameCount, lavaFloorFrameSpeed)
			drawNormalImage(screen, subImage, f.x, f.y)
	}

	if (f.floorType == FloorLava && f.fire.creating) {
		f.fire.Draw(screen)
	}
	
}