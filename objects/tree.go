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

func (t *Tree) MoveTo(x int, y int) {
	t.x = x
	t.y = y
}

func (t *Tree) Draw(screen *ebiten.Image) {
	lib.DrawNormalImage(screen, t.img, t.x, t.y)
}

func CalculateTreePositionX(floorPosX int, randValue int) int {
	posX := floorPosX
	// move the tree out of the screen
	if randValue < 40 {
		posX = -1000
	}
	return posX
}