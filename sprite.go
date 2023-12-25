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

func LoadFolderSprites(directory string, spriteNames []string) {

	for _, spriteName := range spriteNames{
		sprites[spriteName] = loadImage("assets/img/" + directory + spriteName + ".png")
	}

}

func LoadSprites(){

	sprites = make(map[string]*ebiten.Image)

	playerSprites := []string{
		"player_walk_left_with_fuel",
		"player_walk_right_with_fuel",
		"player_walk_left",
		"player_walk_right",
		"player_right",
		"player_left",
		"player_right_with_fuel",
		"player_left_with_fuel",
		"player_center",
		"fire_right",
		"fire_left",
		"fire_center",
	}

	hudSprites := []string{
		"hud",
		"heart",
		"cloud",
	}

	rocketSprites := []string{
		"rocket",
		"rocket_fuel_indicator_off",
		"rocket_fuel_indicator_on",
	}

	backgroundsSprites := []string{
		"background1",
		"background2",
	}

	floorsSprites := []string{
		"floor1",
		"lava_floor",
	}

	enemiesSprites := []string{
		"enemy1",
		"enemy1_closing_eyes",
		"enemy1_opening_eyes",
	}

	othersSprites := []string{
		"fuel",
		"platform",
		"parachute",
	}

	particlesSprites := []string{
		"smoke",
		"explosion",
	}

	LoadFolderSprites("player/", playerSprites)
	LoadFolderSprites("hud/", hudSprites)
	LoadFolderSprites("rocket/", rocketSprites)
	LoadFolderSprites("backgrounds/", backgroundsSprites)
	LoadFolderSprites("floors/", floorsSprites)
	LoadFolderSprites("enemies/", enemiesSprites)
	LoadFolderSprites("others/", othersSprites)
	LoadFolderSprites("particles/", particlesSprites)

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