package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Playertesting struct {
	x  int
	y  int
	vx int
	vy int
	engineOn bool
	engineTimeToTurnOff int
}

func (p *Playertesting) GetCenter() (int, int) {

	// playerCenterX := p.x + 1400
	playerCenterX := p.x + (playerHeight * 4)
	playerCenterY := p.y + (playerHeight * 2)

	return playerCenterX, playerCenterY

}



func (p *Playertesting) MoveUp() {
	p.engineOn = true
	p.engineTimeToTurnOff = 3
	p.vy -= acceleration
	if p.vy < -maxVy {
		p.vy = -maxVy
	}
}


func (p *Playertesting) Position() (int, int) {
	return p.x, p.y
}

func (p *Playertesting) isMovingToRight() bool {
	return p.vx > speedToChangeSprite
}

func (p *Playertesting) isMovingToLeft() bool {
	return p.vx < -speedToChangeSprite
}

func (p *Playertesting) Draw(screen *ebiten.Image) {

	imgPlayer := sprites["player_center_testing"]
	firePlayer := sprites["fire_center"]
	
	switch {
	case p.isMovingToRight():
		firePlayer = sprites["fire_right"]
		imgPlayer = sprites["player_right"]
	case p.isMovingToLeft():
		firePlayer = sprites["fire_left"]
		imgPlayer = sprites["player_left"]
	}

	op := &ebiten.DrawImageOptions{}
	x, y := p.Position()

	op2 := &ebiten.DrawImageOptions{}
	
	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)
	// fire
	if p.engineOn{

		switch {
		case p.isMovingToRight():
			op2.GeoM.Translate(float64(x)/unit, float64(y)/unit + 30)
		case p.isMovingToLeft():
			op2.GeoM.Translate(float64(x)/unit + 35, float64(y)/unit + 30)
		default:
			//center
			op2.GeoM.Translate(float64(x)/unit + 17, float64(y)/unit + 30)
		}
		op2.GeoM.Scale(scale, scale)
		screen.DrawImage(firePlayer, op2)
	}
	screen.DrawImage(imgPlayer, op)
}

