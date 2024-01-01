package main

const (
	appWidth            = 1024
	appHeight           = 768
	scale               = 1
	unit                = 16
	groundY             = 665
	speedToChangeSprite = 6
	initialSoundEnabled = false
)

const (
	startPlayerX    = 3000
	startPlayerY    = 4000
	playerHeight    = 64
	playerWidth     = 32
	walkSpeed       = 50
	acceleration    = 9
	gravitySpeed    = 3
	maxGravitySpeed = 90
	maxVx           = 81
	maxVy           = 81
	maxTimeToIdle   = 5
	playerMaxRight  = 20000
	playerMaxLeft   = -70
	playerMaxUp     = 2000
	playerMaxDown   = 12800
)

const (
	startEnemyX   = 3000
	startEnemyY   = 8000
	enemyHeight   = 128
	enemyWidth    = 128
	enemySpeed    = 50
	enemyMaxRight = 14600
	enemyMaxLeft  = -100
	enemyMaxUp    = 1800
	enemyMaxDown  = 9800
)

const (
	startFuelX            = 0
	startFuelY            = 0
	FallingFuelVelocity   = 2
	offsetFuelLandingY    = 29
	minOffsetFuelLandingX = 20
)

const (
	startRocketX         = 12000
	startRocketY         = -500
	landedRocketY        = 9560
	rocketMaxSpeed       = 20
	startRocketFuelItems = 0
)

const (
	startingLevel        = 1
	totalGameLevels      = 2
	travelingTextMaxTime = 160
	marginTopPlatforms   = 62
	marginLeftPlatforms  = 0
	platformWidthLanding = 166
	floorWidth           = 180
	floorY               = 300
	floorHeight          = 48
)

const (
	frameOX    = 0
	frameOY    = 0
	frameCount = 4
)

const (
	blinkingStarFrameWidth     = 16
	blinkingStarFrameHeight    = 16
	changeBlinkingStarsMaxTime = 50
)
const (
	enemy1ClosingEyesFrameWidth  = 128
	enemy1ClosingEyesFrameHeight = 128
)
const (
	playerWalkFrameWidth  = 32
	playerWalkFrameHeight = 64
)

const (
	lavaFloorFrameWidth     = 180
	lavaFloorFrameHeight    = 53
	randomWithFireLavaFloor = 112
)

const (
	maxOxygenCapacity      = 457
	maxOxygenTimeToConsume = 10
)

const (
	maxTimeToShowSmoke     = 175
	maxTimeToShowExplosion = 15
	smokeX                 = 800
	smokeOffsetY           = 320
)
