package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tree struct {
	x				int
	y				int
	img 			*ebiten.Image
}

const(
	TreeHeight = 128
)

func NewTree(treeSprite *ebiten.Image, x int, y int) *Tree {
	
	return &Tree{
		x: x,
		y: y,
		img: treeSprite,
	}
}

func (t *Tree) Position() (int, int) {
	return t.x, t.y
}

func (t *Tree) MoveTo(x, y int) {
	t.x = x
	t.y = y
}

func (t *Tree) Draw(screen *ebiten.Image) {
	lib.DrawNormalImage(screen, t.img, t.x, t.y)
}

func CalculateTreePositionX(floorPosX, randValue int, isRocketFloor bool) int {

	if isRocketFloor || randValue < 40 {
		return -1000 // move the tree out of the screen
	}
	return floorPosX
}