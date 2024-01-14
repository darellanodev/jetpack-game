package main

const (
	appWidth            = 1024
	appHeight           = 768
	scale               = 1
	groundY             = 665
	speedToChangeSprite = 6
	initialSoundEnabled = true
)

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
	playerMaxDown      = 800
	fuelHandXOffset    = 25
	fuelHandYOffset    = 6
	horizontalFriction = 0.1
	fireRightWidth     = 32
)

const (
	startEnemyX   = 187
	startEnemyY   = 500
	enemyHeight   = 128
	enemyWidth    = 128
	enemySpeed    = 3
	enemyMaxRight = 970
	enemyMaxLeft  = 60
	enemyMaxUp    = 170
	enemyMaxDown  = 670
)

const (
	startFuelX            = 0
	startFuelY            = 0
	FallingFuelVelocity   = 2
	offsetFuelLandingY    = 29
	minOffsetFuelLandingX = 20
)

const (
	startRocketX         = 750
	startRocketY         = -31
	landedRocketY        = 597
	rocketMaxSpeed       = 5
	startRocketFuelItems = 0
	rocketWidth          = 64
	rocketHeight         = 128
	rocketAcceleration   = 0.032
)

const (
	startingLevel             = 0
	totalRowsTxt              = 7
	totalLevelRowsTxt         = 6
	firstLevelRowTxt          = 1
	platformLevelCharacter    = "-"
	emptyLevelCharacter       = "0"
	normalFloorLevelCharacter = "1"
	lavaFloorLevelCharacter   = "2"

	travelingTextMaxTime = 160
	marginTopPlatforms   = 62
	marginLeftPlatforms  = 0
	platformWidthLanding = 150
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
	blinkingStarFrameSpeed     = 5
	changeBlinkingStarsMaxTime = 50
)
const (
	enemy1ClosingEyesFrameWidth  = 128
	enemy1ClosingEyesFrameHeight = 128
	enemy1ClosingEyesFrameSpeed  = 20
)
const (
	playerWalkFrameWidth  = 32
	playerWalkFrameHeight = 64
	playerWalkFrameSpeed  = 5
)

const (
	lavaFloorFrameWidth  = 180
	lavaFloorFrameHeight = 53
	lavaFloorFrameSpeed  = 5
)

const (
	maxOxygenCapacity      = 610
	maxOxygenTimeToConsume = 10
)

const (
	maxTimeToShowSmoke     = 175
	maxTimeToShowExplosion = 25
	smokeX                 = 800
	smokeOffsetY           = 320
)

const (
	offsetSecondTextLine = 50
	offsetThirdTextLine  = 100
	offsetFourthTextLine = 150
)

const (
	CRLF = "\r\n"
)