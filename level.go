package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Level struct {
	number				int
	title				string
	platformPlaces		[5]string
	floorPlaces			string
}


func (f *Level) Draw(screen *ebiten.Image) {

	// f.currentSprite = sprites["level"]

	// op := &ebiten.DrawImageOptions{}
	// x, y := f.position()

	// op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	// op.GeoM.Scale(scale, scale)
	// screen.DrawImage(f.currentSprite, op)
}

func (f *Level) Next() {
	f.number++
	f.Load()
}

func (f *Level) Load() {

	switch f.number {
	case 1:
		f.title = "the kind planet"
		f.platformPlaces[0] = "00000"
		f.platformPlaces[1] = "10000"
		f.platformPlaces[2] = "01000"
		f.platformPlaces[3] = "00000"
		f.platformPlaces[4] = "00000"
		f.floorPlaces = 	  "111111"
	case 2:
		f.title = "fire everywhere"
		f.platformPlaces[0] = "00000"
		f.platformPlaces[1] = "00001"
		f.platformPlaces[2] = "01000"
		f.platformPlaces[3] = "00000"
		f.platformPlaces[4] = "00000"
		f.floorPlaces = 	  "121211"

	}
}


func (f *Level) Update() {
	
	// if f.right{
	// 	f.x += 1
	// }
	// if f.left{
	// 	f.x -= 1
	// }
	// if f.up{
	// 	f.y -= 1
	// }
	// if f.down{
	// 	f.y += 1
	// }
	
}