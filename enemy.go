package main

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x  int
	y  int
	currentSprite *ebiten.Image
	up bool
	down bool
	left bool
	right bool
	timeToCloseEyesMax int
	timeToCloseEyes int
	spriteCount int
	spriteSpeed int
	isClosingEyes bool
}

func (e *Enemy) position() (int, int) {
	return e.x, e.y
}

func (e *Enemy) Draw(screen *ebiten.Image) {

	e.currentSprite = sprites["enemy1"]

	op := &ebiten.DrawImageOptions{}
	x, y := e.position()

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)

	if (e.timeToCloseEyes < e.timeToCloseEyesMax) {
		e.timeToCloseEyes++
		screen.DrawImage(e.currentSprite, op)

	} else {

		i := (e.spriteCount / e.spriteSpeed) % frameCount
		sx, sy := frameOX+i*enemy1ClosingEyesFrameWidth, frameOY
		e.spriteCount++
		
		if (!e.isClosingEyes && i < frameCount) {
			screen.DrawImage(sprites["enemy1_closing_eyes"].SubImage(image.Rect(sx, sy, sx+enemy1ClosingEyesFrameWidth, sy+enemy1ClosingEyesFrameHeight)).(*ebiten.Image), op)
			if (i == frameCount - 1) {
				e.isClosingEyes = true
				e.spriteCount = 0
				i = 0
				screen.DrawImage(e.currentSprite, op)
			}
		}
		
		if (e.isClosingEyes && i < frameCount) {
			screen.DrawImage(sprites["enemy1_opening_eyes"].SubImage(image.Rect(sx, sy, sx+enemy1ClosingEyesFrameWidth, sy+enemy1ClosingEyesFrameHeight)).(*ebiten.Image), op)
			if (i == frameCount - 1) {
				e.isClosingEyes = false
				e.spriteCount = 0
				i = 0
				screen.DrawImage(e.currentSprite, op)
				e.timeToCloseEyes = 0
			}

		}

	}
}

func (e *Enemy) Update() {

	enemyCenterX := e.x + (enemyWidth / 2)
	enemyCenterY := e.y + (enemyHeight / 2)

	if e.right{
		e.x += enemySpeed
	}
	if e.left{
		e.x -= enemySpeed
	}
	if e.up{
		e.y -= enemySpeed
	}
	if e.down{
		e.y += enemySpeed
	}
	
	if enemyCenterX > enemyMaxRight {
		e.left = true
		e.right = false
	}
	if enemyCenterX < enemyMaxLeft {
		e.right = true
		e.left = false
	}
	if enemyCenterY < enemyMaxUp {
		e.down = true
		e.up = false
	}
	if enemyCenterY > enemyMaxDown {
		e.up = true
		e.down = false
	}
	
}