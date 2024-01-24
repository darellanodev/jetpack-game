package main

import (
	"math/rand"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
)

type GameStatus int

const (
	GameStatusPlaying GameStatus = iota
	GameStatusPaused
	GameStatusGameOver
	GameStatusInit
	GameStatusLanding
	GameStatusFinishingLevel
	GameStatusTravelingToLevel
	GameStatusGameComplete
	GameStatusShowingError
)

type Game struct {
	player 			  		*Player
	enemy  			  		*objects.Enemy
	fuel   			  		*objects.Fuel
	rocket			  		*objects.Rocket
	level			  		*Level
	platforms		  		[]*Platform
	floors			  		[]*Floor
	blinkingStars	  		[]*BlinkingStar
	changeBlinkingStarsTime int
	hud				  		*Hud
	smoke			  		*ParticlesSystem
	showSmokeTime     		int
	explosion		  		*ParticlesSystem
	showExplosionTime 		int
	pauseTime 		  		int
	soundTime 		  		int
	soundTextTime	  		int
	pausePressed 	  		bool
	soundPressed	  		bool
	status			  		GameStatus
	travelingTextTime 		int
	count			  		int
}

func (g *Game) Init() error {

	if err := LoadSprites(); err != nil {
		return err
	}

	if err := lib.LoadFonts(assets); err != nil {
		return err
	}

	if err := LoadSounds(); err != nil {
		return err
	}

	if err := LoadLevels(); err != nil {
		return err
	}

	g.player = NewPlayer()
	g.enemy = objects.NewEnemy(sprites["enemy1"], sprites["enemy1_closing_eyes"], sprites["enemy1_closing_eyes"])
	g.fuel = objects.NewFuel(sprites["fuel"], sprites["parachute"])
	g.rocket = objects.NewRocket(sprites["fire_center"],sprites["rocket_fuel_indicator_on"],sprites["rocket_fuel_indicator_off"],sprites["rocket"])
	g.hud = NewHud()
	g.level = NewLevel()
	g.smoke = NewSmoke()
	g.explosion = NewExplosion()

	g.blinkingStars = []*BlinkingStar{NewBlinkingStar(), NewBlinkingStar()}
	g.floors = []*Floor{NewFloor(), NewFloor(), NewFloor(), NewFloor(), NewFloor(), NewFloor()}
	g.platforms = []*Platform{NewPlatform(), NewPlatform()}	

	return nil
}

func (g *Game) putFuelIntoRocket() {
	if (g.rocket.FuelIndicatorItems < 4) {
		g.rocket.FuelIndicatorItems++
		sounds["rocket_fuel_drop"].Play()
		g.player.hasFuel = false
		g.restartFuel()
	} else {
		g.rocket.FuelIndicatorItems++
		sounds["rocket_fuel_drop"].Play()
		sounds["rocket_move"].Play()
		g.smoke.creating = true
		g.showSmokeTime = 0
		g.status = GameStatusFinishingLevel
	}
	//TODO: else: level completed!

}

func (g *Game) placeLevelPlatforms() {
	indexPlatform := 0
	px := 0
	py := 0
	for _, platformPlace := range g.level.platformPlaces {
		px = 0
		for _, char := range platformPlace {
			if string(char) == platformLevelCharacter {
				g.platforms[indexPlatform].x = px * 210 + marginLeftPlatforms
				// fmt.Println("px", g.platforms[indexPlatform].x)

				g.platforms[indexPlatform].y = py * 210 + marginTopPlatforms
				// fmt.Println("py", g.platforms[indexPlatform].y)
				indexPlatform++
			}
			px ++
		}
		py++
	}
}

func (g *Game) placeLevelFloors() {

	px := 0
	py := appHeight - floorHeight

	indexFloor := 0
	for _, char := range g.level.floorPlaces {
		if string(char) == normalFloorLevelCharacter {
			g.floors[indexFloor].floorType = FloorNormal
		} else if string(char) == lavaFloorLevelCharacter {
			g.floors[indexFloor].floorType = FloorLava		
		}
		g.floors[indexFloor].MoveTo(px,py)
		g.floors[indexFloor].InitFloor()
		px += floorWidth
		indexFloor++
	}

}

func (g *Game) restartFuel() {
	g.fuel.Snaps = false

	randomIndex := rand.Intn(len(g.platforms))
	randomPlatform := g.platforms[randomIndex]

	px, py := randomPlatform.position()

	g.fuel.SetFinalPositionIntoPlatform(px, py, platformWidthLanding)

}

func (g *Game) restartPlayer() {
	g.player.x = g.rocket.GetX() - 30
	g.player.y = groundY - playerOffsetY
	g.player.hasFuel = false
}

func (g *Game) restartGame() {
	g.restartPlayer()
	g.restartFuel()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}

func NewGame() *Game {
	g := &Game{
		pausePressed: 		false,
		pauseTime: 			0,
		soundPressed:		false,
		soundTime:			0,
		status: 			GameStatusInit,
		travelingTextTime:  travelingTextMaxTime,
		count:				0,
		soundTextTime:		0,
		showSmokeTime:      0,
	}

	return g
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}
