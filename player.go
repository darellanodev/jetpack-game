package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x  					int
	y  					int
	vx 					int
	vy 					int
	lives				int
	currentSprite 		*ebiten.Image
	engineOn 			bool
	engineTimeToTurnOff int
}

func (p *Player) LostLive() {
	if (p.lives > 0) {
		p.lives--
	}
}

func (p *Player) GetCenter() (int, int) {

	playerCenterX := p.x + ((playerWidth * unit) / 2)
	playerCenterY := p.y + ((playerHeight * unit) / 2)

	return playerCenterX, playerCenterY

}


func (p *Player) MoveRight() {

	if p.isInGround(){
		// player is walking
		p.x += walkSpeed
		p.vx = walkSpeed // to change sprite

	} else {
		// player is flying
		p.vx += acceleration
		if p.vx > maxVx {
			p.vx = maxVx
		}
	}

}

func (p *Player) MoveLeft() {

	if p.isInGround(){
		p.x -= walkSpeed
		p.vx = -walkSpeed // to change sprite
	} else {
		// player is flying
		p.vx -= acceleration
		if p.vx < -maxVx {
			p.vx = -maxVx
		}
	}

}

func (p *Player) MoveUp() {
	p.engineOn = true
	p.engineTimeToTurnOff = 3
	p.vy -= acceleration
	if p.vy < -maxVy {
		p.vy = -maxVy
	}
}


func (p *Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) isMovingToRight() bool {
	return p.vx > speedToChangeSprite
}

func (p *Player) isMovingToLeft() bool {
	return p.vx < -speedToChangeSprite
}

func (p *Player) Draw(screen *ebiten.Image) {

		p.currentSprite = sprites["player_center"]
		firePlayer := sprites["fire_center"]
		
		switch {
		case p.isMovingToRight():
			firePlayer = sprites["fire_right"]
			p.currentSprite = sprites["player_right"]
		case p.isMovingToLeft():
			firePlayer = sprites["fire_left"]
			p.currentSprite = sprites["player_left"]
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
				op2.GeoM.Translate(float64(x)/unit - 15, float64(y)/unit + 30)
			case p.isMovingToLeft():
				op2.GeoM.Translate(float64(x)/unit + 15, float64(y)/unit + 30)
			default:
				//center
				op2.GeoM.Translate(float64(x)/unit, float64(y)/unit + 30)
			}
			op2.GeoM.Scale(scale, scale)
			screen.DrawImage(firePlayer, op2)
		}
		screen.DrawImage(p.currentSprite, op)
}

func (p *Player) isInGround() bool {
	return p.y >= groundY * unit
}

func (p *Player) horizontalFriction() {
	if (p.vx < 0) {
		p.vx += 1
	}
	if (p.vx > 0){
		p.vx -= 1
	}
}

func (p *Player) gravity() {
	p.vy += gravitySpeed
	if p.vy > maxGravitySpeed {
		p.vy = maxGravitySpeed
	}
}

func (p *Player) Update() {

	isFlying := !p.isInGround()

	if (isFlying) {
		// Gravity
		p.gravity()
		// Horizontal friction
		p.horizontalFriction()
		// Horizontal movement when player is flying
		p.x += p.vx
	}

	p.y += p.vy
	// floor
	if p.isInGround() {
		p.y = groundY * unit
		p.horizontalFriction()
	}

	if p.engineOn {
		p.engineTimeToTurnOff -= 1
	}
	if p.engineTimeToTurnOff == 0 {
		p.engineOn = false
	}


}
