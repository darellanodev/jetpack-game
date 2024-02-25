package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)


type Lavadrop struct {
	x  					int
	y  					int
	initialX  			int
	initialY            int
	vx 					float64
	vy 					float64
	isMovingUp			bool
	collisionHitBox 	*ebiten.Image
	img     			*ebiten.Image
}

const (
	lavadropWith = 32
	lavadropMaxGravitySpeed = 10
)

func NewLavadrop(lavaDropImg *ebiten.Image) *Lavadrop {
	
	return &Lavadrop{
		x:     			 0,
		y:     			 0,
		vx:    			 0,
		vy:    			 -15,
		isMovingUp: 	 true,
		collisionHitBox: lavaDropImg,
		img:			 lavaDropImg,
	}
}

func (l *Lavadrop) reinit() {
	l.x = l.initialX
	l.y = l.initialY
	l.vx = 0
	l.vy = -15
	l.isMovingUp = true
}

func (l *Lavadrop) SetInitialPosition(x int, y int) {
	l.x = x
	l.y = y
	l.initialX = x
	l.initialY = y
}

func (l *Lavadrop) CollisionHitBox() *ebiten.Image {
	return l.collisionHitBox
}

func (l *Lavadrop) Position() (int, int) {
	return l.x, l.y
}

func (l *Lavadrop) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, l.img, l.x, l.y)	
}

func (l *Lavadrop) gravity() {
	l.vy += gravitySpeed
	if l.vy > lavadropMaxGravitySpeed {
		l.vy = lavadropMaxGravitySpeed
	}
}

func (l *Lavadrop) Update() bool {

	l.gravity()
	l.y += int(l.vy)

	if l.y > 1000 {
		l.reinit()
		return false
	}
	return true
	
}