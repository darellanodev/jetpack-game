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
		posX: f.x/32,
		posY: f.y/32,
	}
}

func (f *Floor) position() (int, int) {
	return f.x, f.y
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
		f.fire.UpdateUp(randomWithFireLavaFloor)
	}
}


func (f *Floor) Draw(screen *ebiten.Image, spriteCount int) {


	i := (spriteCount / 5) % frameCount
	sx, sy := frameOX+i*lavaFloorFrameWidth, frameOY

	op := &ebiten.DrawImageOptions{}
	x, y := f.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)

	switch f.floorType {
		case FloorNormal:
			screen.DrawImage(sprites["floor1"], op)
		case FloorLava:
			screen.DrawImage(sprites["lava_floor"].SubImage(image.Rect(sx, sy, sx+lavaFloorFrameWidth, sy+lavaFloorFrameHeight)).(*ebiten.Image), op)
	}

	if (f.floorType == FloorLava && f.fire.creating) {
		f.fire.Draw(screen)
	}

	
}
