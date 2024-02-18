package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	enemy1ClosingEyesFrameWidth  = 128
	enemy1ClosingEyesFrameHeight = 128
	enemy1ClosingEyesFrameSpeed  = 20

	enemyHeight   = 128
	enemyWidth    = 128
	enemySpeed    = 3
	enemyMaxRight = 970
	enemyMaxLeft  = 60
	enemyMaxUp    = 170
	enemyMaxDown  = 670
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
	collisionHitBox 	*ebiten.Image
	img     			*ebiten.Image
	imgAnimClosingEyes  *ebiten.Image
	imgAnimOpeningEyes  *ebiten.Image
}

func NewEnemy(enemySprites []*ebiten.Image) *Enemy {
	
	return &Enemy{
		x:     				187,
		y:     				500,
		up:    				true,
		down:  				false,
		left:  				false,
		right: 				true,
		timeToCloseEyesMax: 200,
		timeToCloseEyes: 	0,
		spriteCount: 		0,
		spriteSpeed: 		20,
		isClosingEyes: 		false,
		collisionHitBox:	enemySprites[0],
		img:				enemySprites[0],
		imgAnimClosingEyes: enemySprites[1],
		imgAnimOpeningEyes: enemySprites[2],
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

	lib.DrawNormalImage(screen, e.img, e.x, e.y)

	if (e.timeToCloseEyes < e.timeToCloseEyesMax) {
		e.timeToCloseEyes++
		lib.DrawNormalImage(screen, e.img, e.x, e.y)

	} else {

		i := (e.spriteCount / e.spriteSpeed) % frameCount
		e.spriteCount++
		
		if (!e.isClosingEyes && i < frameCount) {
			subImage = lib.GetSubImage(e.imgAnimClosingEyes, enemy1ClosingEyesFrameWidth, enemy1ClosingEyesFrameHeight, e.spriteCount, frameCount, enemy1ClosingEyesFrameSpeed)
			lib.DrawNormalImage(screen, subImage, e.x, e.y)
			
			if (i == frameCount - 1) {
				e.isClosingEyes = true
				e.spriteCount = 0
				i = 0
				lib.DrawNormalImage(screen, e.img, e.x, e.y)
			}
		}
		
		if (e.isClosingEyes && i < frameCount) {
			subImage = lib.GetSubImage(e.imgAnimOpeningEyes, enemy1ClosingEyesFrameWidth, enemy1ClosingEyesFrameHeight, e.spriteCount, frameCount, enemy1ClosingEyesFrameSpeed)
			lib.DrawNormalImage(screen, subImage, e.x, e.y)

			if (i == frameCount - 1) {
				e.isClosingEyes = false
				e.spriteCount = 0
				i = 0
				lib.DrawNormalImage(screen, e.img, e.x, e.y)
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