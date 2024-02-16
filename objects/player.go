package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
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
	Lives				int
	engineOn 			bool
	engineTimeToTurnOff int
	PlayerStatus		PlayerStatus
	timeToIdle			int
	HasFuel				bool
	InmuneToDamageTime	int
	collisionHitBox		*ebiten.Image
	imgPlayerCenter		*ebiten.Image
	imgPlayerRight		*ebiten.Image
	imgPlayerRightWithFuel *ebiten.Image
	imgFireRight		*ebiten.Image
	imgFireCenter		*ebiten.Image
	imgPlayerWalkRightWithFuel *ebiten.Image
	imgPlayerWalkRight  *ebiten.Image
}

const (
	playerOffsetY      = 5
	playerHeight       = 64
	playerWidth        = 32
	walkSpeed          = 4
	acceleration       = 0.4
	gravitySpeed       = 0.3
	maxGravitySpeed    = 4
	maxVx              = 5
	maxVy              = 8
	maxTimeToIdle      = 5
	playerMaxRight     = 998
	playerMaxLeft      = -3
	playerMaxUp        = 135
	fuelHandXOffset    = 25
	fuelHandYOffset    = 6
	horizontalFriction = 0.1
	fireRightWidth     = 32
	playerWalkFrameWidth  = 32
	playerWalkFrameHeight = 64
	playerWalkFrameSpeed  = 5	
	groundY = 665
	initialLives = 3

)

func NewPlayer(imgPlayerCenter *ebiten.Image, imgFireRight *ebiten.Image, imgFireCenter *ebiten.Image, imgPlayerWalkWithFuel *ebiten.Image, imgPlayerWalkRight *ebiten.Image, imgPlayerRight *ebiten.Image, imgPlayerRightWithFuel *ebiten.Image) *Player {

	return &Player{
		x: 				    0,
		y: 				    0,
		Lives:			    initialLives,
		PlayerStatus:       Center,
		timeToIdle:		    maxTimeToIdle,
		HasFuel:		    false,
		InmuneToDamageTime: 0,
		collisionHitBox: 	imgPlayerCenter,
		imgPlayerCenter:	imgPlayerCenter,
		imgFireRight:		imgFireRight,
		imgFireCenter:		imgFireCenter,
		imgPlayerWalkRightWithFuel: imgPlayerWalkWithFuel,
		imgPlayerWalkRight: imgPlayerWalkRight,
		imgPlayerRight:		imgPlayerRight,
		imgPlayerRightWithFuel: imgPlayerRightWithFuel,
	}
}

func (p *Player) CollisionHitBox() *ebiten.Image {
	return p.collisionHitBox
}

func (p *Player) Position() (int, int) {
	return p.x, p.y
}

func (p *Player) GetY() int {
	return p.y
}

func (p *Player) GetX() int {
	return p.x
}

func (p *Player) LostLive() {
	if (p.Lives > 0) {
		p.Lives--
	}
}

func (p *Player) GetCenter() (int, int) {

	playerCenterX := p.x + (playerWidth / 2)
	playerCenterY := p.y + (playerHeight / 2)

	return playerCenterX, playerCenterY

}

func (p *Player) MoveTo(x int, y int) {
	p.x = x
	p.y = y
}

func (p *Player) RestartLives() {
	p.Lives = initialLives
}

func (p *Player) Restart(posX int) {
	p.x = posX
	p.y = groundY - playerOffsetY
	p.HasFuel = false
}

func (p *Player) MoveRight() {

	p.timeToIdle = maxTimeToIdle

	if p.isInGround(){
		p.vx = walkSpeed
		p.x = int(float64(p.x) + p.vx)
		if (p.HasFuel){
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

		if (p.HasFuel){
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

			lib.DrawNormalImage(screen, p.imgFireRight, p.x - 15, p.y + 30)

		} else if (p.isMovingToTheLeft()) {

			lib.DrawHorizontalFlippedImage(screen, p.imgFireRight, fireRightWidth, p.x + 15, p.y + 30)

		} else {

			lib.DrawNormalImage(screen, p.imgFireCenter, p.x, p.y + 30)
		}

	}
}

func (p *Player) getImgFlyingRight() *ebiten.Image {
	
	imgFlyingRight := p.imgPlayerRight
	if p.HasFuel {
		imgFlyingRight = p.imgPlayerRightWithFuel
	}

	return imgFlyingRight
}

func (p *Player) drawPlayer(screen *ebiten.Image, spriteCount int) {

	walkingRightWithFuelSubImage := lib.GetSubImage(p.imgPlayerWalkRightWithFuel, playerWalkFrameWidth, playerWalkFrameHeight, spriteCount, frameCount, playerWalkFrameSpeed)
	walkingRightSubImage := lib.GetSubImage(p.imgPlayerWalkRight, playerWalkFrameWidth, playerWalkFrameHeight, spriteCount, frameCount, playerWalkFrameSpeed)

	switch p.PlayerStatus {

		case WalkingRightWithFuel:
			lib.DrawNormalImage(screen, walkingRightWithFuelSubImage, p.x, p.y)
			
		case WalkingLeftWithFuel:
			lib.DrawHorizontalFlippedImage(screen, walkingRightWithFuelSubImage, playerWalkFrameWidth, p.x, p.y)
			
		case WalkingRight:
			lib.DrawNormalImage(screen, walkingRightSubImage, p.x, p.y)

		case WalkingLeft:
			lib.DrawHorizontalFlippedImage(screen, walkingRightSubImage, playerWalkFrameWidth, p.x, p.y)
			
		case FlyingRight:
			lib.DrawNormalImage(screen, p.getImgFlyingRight(), p.x, p.y)

		case FlyingLeft:
			lib.DrawHorizontalFlippedImage(screen, p.getImgFlyingRight(), playerWidth, p.x, p.y)
		
		default:
			lib.DrawNormalImage(screen,p.imgPlayerCenter, p.x, p.y)
	}
}


func (p *Player) Draw(screen *ebiten.Image, spriteCount int) {

	doDraw := true
	
	if (p.InmuneToDamageTime > 0) {
		if (p.InmuneToDamageTime % 3 == 0) {
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