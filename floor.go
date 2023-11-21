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
	currentSprite 	*ebiten.Image
	floorType		FloorType
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


func (f *Floor) Draw(screen *ebiten.Image) {

	switch f.floorType {
		case FloorNormal:
			f.currentSprite = sprites["floor1"]
		case FloorLava:
			f.currentSprite = sprites["lava"]
	}

	op := &ebiten.DrawImageOptions{}
	x, y := f.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(f.currentSprite, op)
}
