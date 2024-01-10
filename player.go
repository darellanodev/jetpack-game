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
	vx 					float64
	vy 					float64
	lives				int
	engineOn 			bool
	engineTimeToTurnOff int
	PlayerStatus		PlayerStatus
	timeToIdle			int
	hasFuel				bool
	inmuneToDamageTime	int
	collisionHitBox		*ebiten.Image
}

func (p *Player) CollisionHitBox() *ebiten.Image {
	return p.collisionHitBox
}

func (p *Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) LostLive() {
	if (p.lives > 0) {
		p.lives--
	}
}

func (p *Player) GetCenter() (int, int) {

	playerCenterX := p.x + (playerWidth / 2)
	playerCenterY := p.y + (playerHeight / 2)

	return playerCenterX, playerCenterY

}

func (p *Player) MoveRight() {

	p.timeToIdle = maxTimeToIdle



	if p.isInGround(){
		p.vx = walkSpeed
		p.x = int(float64(p.x) + p.vx)
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
		p.x = int(float64(p.x) + p.vx)

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

func (p *Player) isMovingToTheRight() bool {
	return p.PlayerStatus == WalkingRight || p.PlayerStatus == WalkingRightWithFuel || p.PlayerStatus == FlyingRight
}

func (p *Player) isMovingToTheLeft() bool {
	return p.PlayerStatus == WalkingLeft || p.PlayerStatus == WalkingLeftWithFuel || p.PlayerStatus == FlyingLeft
}

func (p *Player) HandsPosition() (int, int) {
	if (p.isMovingToTheRight()) {
		return p.x + fuelHandXOffset, p.y + fuelHandYOffset
	} else if (p.isMovingToTheLeft()) {
		return p.x - fuelHandXOffset, p.y + fuelHandYOffset
	}
	return p.x, p.y + fuelHandYOffset
}

func (p *Player) drawFire(screen *ebiten.Image) {

	if (p.engineOn) {
	
		if (p.isMovingToTheRight()) {

			drawNormalImage(screen, sprites["fire_right"], p.x - 15, p.y + 30)

		} else if (p.isMovingToTheLeft()) {

			drawHorizontalFlippedImage(screen, sprites["fire_right"], fireRightWidth, p.x + 15, p.y + 30)

		} else {

			drawNormalImage(screen, sprites["fire_center"], p.x, p.y + 30)
		}

	}
}

func (p *Player) drawPlayer(screen *ebiten.Image, spriteCount int) {

	i := (spriteCount / 5) % frameCount
	sx, sy := frameOX+i*playerWalkFrameWidth, frameOY

	withFuel := ""
	
	if (p.hasFuel) {
		withFuel = "_with_fuel"
	}

	switch p.PlayerStatus {

		case WalkingRightWithFuel:
			subImage := sprites["player_walk_right_with_fuel"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image)
			drawNormalImage(screen, subImage, p.x, p.y)
			
		case WalkingLeftWithFuel:
			subImage := sprites["player_walk_right_with_fuel"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image)
			drawHorizontalFlippedImage(screen, subImage, playerWalkFrameWidth, p.x, p.y)
			
		case WalkingRight:
			subImage := sprites["player_walk_right"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image)
			drawNormalImage(screen, subImage, p.x, p.y)

		case WalkingLeft:
			subImage := sprites["player_walk_right"].SubImage(image.Rect(sx, sy, sx+playerWalkFrameWidth, sy+playerWalkFrameHeight)).(*ebiten.Image)
			drawHorizontalFlippedImage(screen, subImage, playerWalkFrameWidth, p.x, p.y)
			
		case FlyingRight:
			drawNormalImage(screen,sprites["player_right" + withFuel], p.x, p.y)

		case FlyingLeft:
			drawHorizontalFlippedImage(screen, sprites["player_right" + withFuel], playerWidth, p.x, p.y)
		
		default:
			drawNormalImage(screen,sprites["player_center"], p.x, p.y)
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
	return p.y >= groundY - playerOffsetY
}

func (p *Player) horizontalFriction() {
	if (p.vx < 0) {
		p.vx += horizontalFriction
	}
	if (p.vx > 0){
		p.vx -= horizontalFriction
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
		p.x += int(p.vx)
	}

	p.y += int(p.vy)

	if p.isInGround() {
		p.y = groundY - playerOffsetY
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