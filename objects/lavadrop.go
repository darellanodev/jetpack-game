package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)


type Lavadrop struct {
	x  					int
	y  					int
	isMovingUp			bool
	collisionHitBox 	*ebiten.Image
	img     			*ebiten.Image
}

const (
	lavadropWith = 32
)

func NewLavadrop(lavaDropImg *ebiten.Image) *Lavadrop {
	
	return &Lavadrop{
		x:     				187,
		y:     				500,
		isMovingUp: 		true,
		collisionHitBox:	lavaDropImg,
		img:				lavaDropImg,
		
	}
}

func (l *Lavadrop) MoveTo(x int, y int) {
	l.x = x
	l.y = y
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

func (l *Lavadrop) Update() {

	l.y --
	
}