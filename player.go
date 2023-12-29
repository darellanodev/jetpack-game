package main

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerStatus int

const (
	WalkingLeft PlayerStatus = iota
	WalkingRight
	WalkingRightWithFuel
	WalkingLeftWithFuel
	FlyingLeft
	FlyingRight
	Center
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
	PlayerStatus		PlayerStatus
	timeToIdle			int
	hasFuel				bool
	inmuneToDamageTime	int
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

	p.timeToIdle = maxTimeToIdle



	if p.isInGround(){
		p.vx = walkSpeed
		p.x = p.x + p.vx
		if (p.hasFuel){
			p.PlayerStatus = WalkingRightWithFuel
		}else{
			p.PlayerStatus = WalkingRight
		}
		p.vx = 0

	} else {
		// player is flying
		p.vx += acceleration
		if p.vx > maxVx {
			p.vx = maxVx
		}
		p.PlayerStatus = FlyingRight
	}
}

func (p *Player) MoveLeft() {

	p.timeToIdle = maxTimeToIdle

	if p.isInGround(){
		p.vx = -walkSpeed
		p.x = p.x + p.vx

		if (p.hasFuel){
			p.PlayerStatus = WalkingLeftWithFuel
		} else {
			p.PlayerStatus = WalkingLeft
		}
		p.vx = 0

	} else {
		// player is flying
		p.vx -= acceleration
		if p.vx < -maxVx {
			p.vx = -maxVx
		}
		p.PlayerStatus = FlyingLeft
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

func (p *Player) isMovingToTheRight() bool {
	return p.PlayerStatus == WalkingRight || p.PlayerStatus == WalkingRightWithFuel || p.PlayerStatus == FlyingRight
}

func (p *Player) isMovingToTheLeft() bool {
	return p.PlayerStatus == WalkingLeft || p.PlayerStatus == WalkingLeftWithFuel || p.PlayerStatus == FlyingLeft
}

func (p *Player) HandsPosition() (int, int) {
	if (p.isMovingToTheRight()) {
		return p.x + 400, p.y + 100
	} else if (p.isMovingToTheLeft()) {
		return p.x - 400, p.y + 100
	}
	return p.x, p.y + 100
}

func (p *Player) drawFire(screen *ebiten.Image) {

	x, y := p.Position()
	op := &ebiten.DrawImageOptions{}

	if (p.engineOn) {
	
		if (p.isMovingToTheRight()) {

			op.GeoM.Translate(float64(x)/unit - 15, float64(y)/unit + 30)
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(sprites["fire_right"], op)

		} else if (p.isMovingToTheLeft()) {

			op.GeoM.Translate(float64(x)/unit + 15, float64(y)/unit + 30)
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(sprites["fire_left"], op)

		} else {

			op.GeoM.Translate(float64(x)/unit, float64(y)/unit + 30)
			op.GeoM.Scale(scale, scale)
			screen.DrawImage(sprites["fire_center"], op)
		}

	}
}

func (p *Player) drawPlayer(screen *ebiten.Image, spriteCount int) {

	x, y := p.Position()
	op := &ebiten.DrawImageOptions{}

	i := (spriteCount / 5) % frameCount
	sx, sy := frameOX+i*playerWalkFrameWidth, frameOY

	p.currentSprite = sprites["player_right"]

	withFuel := ""
	
	if (p.hasFuel) {
		withFuel = "_with_fuel"
	}

	op.GeoM.Translate(float64(x)/unit, float64(y)/unit)
	op.GeoM.Scale(scale, scale)

	switch p.PlayerStatus {

		case WalkingRightWithFuel:
			screen.DrawImage(sprites["player_walk_right_with_fuel"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image), op)

		case WalkingLeftWithFuel:
			screen.DrawImage(sprites["player_walk_left_with_fuel"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image), op)

		case WalkingRight:
			screen.DrawImage(sprites["player_walk_right"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image), op)

		case WalkingLeft:
			screen.DrawImage(sprites["player_walk_left"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image), op)

		case FlyingLeft:
			screen.DrawImage(sprites["player_left" + withFuel], op)
		
		case FlyingRight:
			screen.DrawImage(sprites["player_right" + withFuel], op)
			
		default:
			screen.DrawImage(sprites["player_center"], op)
	}
}


func (p *Player) Draw(screen *ebiten.Image, spriteCount int) {

	doDraw := true
	
	if (p.inmuneToDamageTime > 0) {
		if (p.inmuneToDamageTime % 3 == 0) {
			doDraw = false
		}
	}

	if (doDraw){
		p.drawFire(screen)
		p.drawPlayer(screen, spriteCount)
	}
		

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

	if (p.timeToIdle > 0) {
		p.timeToIdle--
	} else {
		p.PlayerStatus = Center
	}

	isFlying := !p.isInGround()

	if (isFlying) {
		p.gravity()
		p.horizontalFriction()
		p.x += p.vx
	}

	p.y += p.vy

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

	if p.x > playerMaxRight{
		p.x = playerMaxRight
	}

	if p.x < playerMaxLeft{
		p.x = playerMaxLeft
	}

	if p.y < playerMaxUp{
		p.y = playerMaxUp
	}

}