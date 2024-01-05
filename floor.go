package main

import (
	"image"
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

	i := (spriteCount / 5) % frameCount
	sx, sy := frameOX+i*lavaFloorFrameWidth, frameOY

	switch f.floorType {
		case FloorNormal:
			NewGame().drawNormalImage(screen,sprites["floor1"],f.x,f.y)
		case FloorLava:
			subImage := sprites["lava_floor"].SubImage(image.Rect(sx, sy, sx+lavaFloorFrameWidth, sy+lavaFloorFrameHeight)).(*ebiten.Image)
			NewGame().drawNormalImage(screen, subImage, f.x, f.y)
	}

	if (f.floorType == FloorLava && f.fire.creating) {
		f.fire.Draw(screen)
	}
	
}