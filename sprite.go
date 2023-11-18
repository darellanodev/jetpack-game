package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var sprites map[string] *ebiten.Image



func loadImage(file string) *ebiten.Image {
	var err error

	// Preload images
	img, _, err := ebitenutil.NewImageFromFile(file)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadSprites(){

	sprites = make(map[string]*ebiten.Image)

	// player folder
	for _, spriteName := range []string{
		"player_right",
		"player_left",
		"player_center",
		"fire_right",
		"fire_left",
		"fire_center",
	} {
		sprites[spriteName] = loadImage("assets/img/player/" + spriteName + ".png")
	}

	// hud folder
	for _, spriteName := range []string{
		"hud",
		"heart",
		"cloud",
	} {
		sprites[spriteName] = loadImage("assets/img/hud/" + spriteName + ".png")
	}

	// rocket folder
	for _, spriteName := range []string{
		"rocket",
		"rocket_fuel_indicator_off",
		"rocket_fuel_indicator_on",
	} {
		sprites[spriteName] = loadImage("assets/img/rocket/" + spriteName + ".png")
	}

	// backgrounds folder
	for _, spriteName := range []string{
		"background1",
		"background2",
	} {
		sprites[spriteName] = loadImage("assets/img/backgrounds/" + spriteName + ".png")
	}

	// others
	for _, spriteName := range []string{
		"enemy",
		"fuel",
		"platform",
	} {
		sprites[spriteName] = loadImage("assets/img/" + spriteName + ".png")
	}

}

func isColliding(sprite1 *ebiten.Image, x1, y1 float64, sprite2 *ebiten.Image, x2, y2 float64) (bool,string) {

	if (sprite1 == nil || sprite2 == nil) {
		return false, "sin sprites"
	}
	bounds1 := sprite1.Bounds()
	bounds2 := sprite2.Bounds()

	r1 := image.Rectangle{
		Min: image.Point{int(x1), int(y1)},
		Max: image.Point{int(x1) + bounds1.Dx(), int(y1) + bounds1.Dy()},
	}
	r2 := image.Rectangle{
		Min: image.Point{int(x2), int(y2)},
		Max: image.Point{int(x2) + bounds2.Dx(), int(y2) + bounds2.Dy()},
	}

	debugInfo := fmt.Sprintf("Sprite 1 bounds: %v\n", bounds1)
	debugInfo += fmt.Sprintf("Sprite 2 bounds: %v\n", bounds2)
	debugInfo += fmt.Sprintf("Sprite 1 position and size: %v\n", r1)
	debugInfo += fmt.Sprintf("Sprite 2 position and size: %v\n", r2)

	result := r1.Intersect(r2) != image.Rectangle{}
	debugInfo += fmt.Sprintf("Collision result: %v\n", result)

	// fmt.Println(debugInfo)

	return result, debugInfo
}