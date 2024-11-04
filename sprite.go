package main

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (

	sprites map[string] *ebiten.Image

	playerFilenameSprites = []string{
		"player_walk_right_with_fuel",
		"player_walk_right",
		"player_right_with_fuel",
		"player_right",
		"player_center",
		"fire_right",
		"fire_center",
	}

	hudFilenameSprites = []string{
		"hud",
		"heart",
		"cloud",
		"live",
		"circle_planet_selector",
	}
	
	rocketFilenameSprites = []string{
		"rocket",
		"rocket_fuel_indicator_off",
		"rocket_fuel_indicator_on",
	}
	
	backgroundsFilenameSprites = []string{
		"background1",
		"background2",
		"blinking_star",
		"mainmenu",
	}
	
	floorsFilenameSprites = []string{
		"vulcan_floor",
		"lava_floor",
	}
	
	enemiesFilenameSprites = []string{
		"enemy1",
		"enemy1_closing_eyes",
		"enemy1_opening_eyes",
		"lava_drop",
	}

	planetFilenameSprites = []string{
		"fire_planet",
		"green_planet",
		"starfield",
		"starfield2",
	}
	
	othersFilenameSprites = []string{
		"fuel",
		"platform",
		"parachute",
		"pillar",
	}
	
	particlesFilenameSprites = []string{
		"smoke",
		"explosion",
		"fire",
	}

	treesSprites = []string{
		"fire_tree_01",
		"fire_tree_02",
		"fire_tree_03",
		"fire_tree_04",
		"fire_tree_05",
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

	if err := LoadFolderSprites("player/", playerFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("hud/", hudFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("rocket/", rocketFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("backgrounds/", backgroundsFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("floors/", floorsFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("enemies/", enemiesFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("others/", othersFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("particles/", particlesFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("planets/", planetFilenameSprites); err != nil {
		return err
	}
	if err := LoadFolderSprites("trees/", treesSprites); err != nil {
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

	if sprite1 == nil || sprite2 == nil {
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