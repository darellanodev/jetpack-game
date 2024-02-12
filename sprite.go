package main

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (

	sprites map[string] *ebiten.Image

	playerSprites = []string{
		"player_walk_right_with_fuel",
		"player_walk_right",
		"player_right_with_fuel",
		"player_right",
		"player_center",
		"fire_right",
		"fire_center",
	}

	hudSprites = []string{
		"hud",
		"heart",
		"cloud",
		"live",
	}
	
	rocketSprites = []string{
		"rocket",
		"rocket_fuel_indicator_off",
		"rocket_fuel_indicator_on",
	}
	
	backgroundsSprites = []string{
		"background1",
		"background2",
		"blinking_star",
	}
	
	floorsSprites = []string{
		"floor1",
		"lava_floor",
	}
	
	enemiesSprites = []string{
		"enemy1",
		"enemy1_closing_eyes",
		"enemy1_opening_eyes",
	}
	
	othersSprites = []string{
		"fuel",
		"platform",
		"parachute",
		"pillar",
	}
	
	particlesSprites = []string{
		"smoke",
		"explosion",
		"fire",
	}
)

func loadImage(filesystem embed.FS, file string) (*ebiten.Image, error) {
	var err error

	// Preload images
	img, _, err := ebitenutil.NewImageFromFileSystem(assets,file)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func LoadFolderSprites(directory string, spriteNames []string) error {
	var err error

	for _, spriteName := range spriteNames{
		path := "assets/img/" + directory + spriteName + ".png"
		sprites[spriteName], err = loadImage(assets, path)
		if err != nil {
			return err
		}
	}

	return nil

}

func LoadSprites() error{

	sprites = make(map[string]*ebiten.Image)

	if err := LoadFolderSprites("player/", playerSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("hud/", hudSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("rocket/", rocketSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("backgrounds/", backgroundsSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("floors/", floorsSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("enemies/", enemiesSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("others/", othersSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("particles/", particlesSprites); err != nil {
		return err
	}

	return nil
}

type GameObject interface {
    CollisionHitBox() *ebiten.Image
	Position() (int,int)
}

func checkCollision [T1 GameObject, T2 GameObject](a T1, b T2) bool {

	aX, aY := a.Position()
	bX, bY := b.Position()

	return isColliding(a.CollisionHitBox(), float64(aX), float64(aY), b.CollisionHitBox(), float64(bX), float64(bY))

}


func isColliding(sprite1 *ebiten.Image, x1, y1 float64, sprite2 *ebiten.Image, x2, y2 float64) bool {

	if (sprite1 == nil || sprite2 == nil) {
		return false
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

	result := r1.Intersect(r2) != image.Rectangle{}

	return result
}