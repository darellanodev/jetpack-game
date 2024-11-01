package main

const (
	appWidth            = 1024
	appHeight           = 768
	scale               = 1
	speedToChangeSprite = 6
	initialSoundEnabled = true
)

const (
	startingLevel     = 0
	totalRowsTxt      = 7
	totalLevelRowsTxt = 6
	firstLevelRowTxt  = 1

	travelingTextMaxTime = 160
	marginTopPlatforms   = 62
	marginLeftPlatforms  = 0
	platformWidthLanding = 150
	floorWidth           = 180
	floorY               = 300
	floorHeight          = 48
)

var LevelCharacters = map[string]string{
	"platform":           "-",
	"empty":              "0",
	"normalFloor":        "1",
	"lavaFloor":          "2",
	"lavaFloorWithDrops": "3",
}

const (
	maxTimeToShowSmoke     = 175
	maxTimeToShowExplosion = 25
	smokeX                 = 800
	smokeOffsetY           = 320
)

const (
	CRLF = "\r\n"
)