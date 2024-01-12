package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	x  					int
	y  					int
	up 					bool
	down 				bool
	left 				bool
	right 				bool
	timeToCloseEyesMax 	int
	timeToCloseEyes 	int
	spriteCount 		int
	spriteSpeed 		int
	isClosingEyes 	    bool
	collisionHitBox     *ebiten.Image
}

func NewEnemy() *Enemy {
	
	return &Enemy{
		x:     				startEnemyX,
		y:     				startEnemyY,
		up:    				true,
		down:  				false,
		left:  				false,
		right: 				true,
		timeToCloseEyesMax: 200,
		timeToCloseEyes: 	0,
		spriteCount: 		0,
		spriteSpeed: 		20,
		isClosingEyes: 		false,
	}
}

func (e *Enemy) CollisionHitBox() *ebiten.Image {
	return e.collisionHitBox
}

func (e *Enemy) Position() (int, int) {
	return e.x, e.y
}

func (e *Enemy) Draw(screen *ebiten.Image) {

	var subImage *ebiten.Image

	drawNormalImage(screen, sprites["enemy1"], e.x, e.y)

	if (e.timeToCloseEyes < e.timeToCloseEyesMax) {
		e.timeToCloseEyes++
		drawNormalImage(screen, sprites["enemy1"], e.x, e.y)

	} else {

		i := (e.spriteCount / e.spriteSpeed) % frameCount
		// sx, sy := frameOX+i*enemy1ClosingEyesFrameWidth, frameOY
		e.spriteCount++
		
		if (!e.isClosingEyes && i < frameCount) {
			subImage = getSubImage(sprites["enemy1_closing_eyes"], enemy1ClosingEyesFrameWidth, enemy1ClosingEyesFrameHeight, e.spriteCount, frameCount, enemy1ClosingEyesFrameSpeed)
			// subImage = sprites["enemy1_closing_eyes"].SubImage(image.Rect(sx, sy, sx+enemy1ClosingEyesFrameWidth, sy+enemy1ClosingEyesFrameHeight)).(*ebiten.Image)
			drawNormalImage(screen, subImage, e.x, e.y)
			
			if (i == frameCount - 1) {
				e.isClosingEyes = true
				e.spriteCount = 0
				i = 0
				drawNormalImage(screen, sprites["enemy1"], e.x, e.y)
			}
		}
		
		if (e.isClosingEyes && i < frameCount) {
			subImage = getSubImage(sprites["enemy1_opening_eyes"], enemy1ClosingEyesFrameWidth, enemy1ClosingEyesFrameHeight, e.spriteCount, frameCount, enemy1ClosingEyesFrameSpeed)
			// subImage = sprites["enemy1_opening_eyes"].SubImage(image.Rect(sx, sy, sx+enemy1ClosingEyesFrameWidth, sy+enemy1ClosingEyesFrameHeight)).(*ebiten.Image)
			drawNormalImage(screen, subImage, e.x, e.y)

			if (i == frameCount - 1) {
				e.isClosingEyes = false
				e.spriteCount = 0
				i = 0
				drawNormalImage(screen, sprites["enemy1"], e.x, e.y)
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