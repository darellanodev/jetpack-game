package main

import (
	"math/rand"

	"github.com/darellanodev/jetpack-game/hud"
	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
	"github.com/darellanodev/jetpack-game/particles"
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
	GameStatusResetGame
	GameStatusMainMenu
	GameStatusPreloadingGame
)

type Game struct {
	player 			  		*objects.Player
	enemy  			  		*objects.Enemy
	fuel   			  		*objects.Fuel
	rocket			  		*objects.Rocket
	mainmenu 				*objects.Mainmenu
	platforms		  		[]*objects.Platform
	floors			  		[]*objects.Floor
	blinkingStars	  		[]*objects.BlinkingStar
	smoke			  		*particles.ParticlesSystem
	explosion		  		*particles.ParticlesSystem
	level			  		*Level
	hud				  		*hud.Hud
	preloadingProgressBar   *hud.Progressbar
	isGamePreloaded			bool
	timeToPreloadGame		int
	changeBlinkingStarsTime int
	showSmokeTime     		int
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

	g.player = objects.NewPlayer(sprites["player_center"], sprites["fire_right"], sprites["fire_center"], sprites["player_walk_right_with_fuel"], sprites["player_walk_right"], sprites["player_right"], sprites["player_right_with_fuel"])
	g.enemy = objects.NewEnemy(sprites["enemy1"], sprites["enemy1_closing_eyes"], sprites["enemy1_closing_eyes"])
	g.fuel = objects.NewFuel(sprites["fuel"], sprites["parachute"])
	g.rocket = objects.NewRocket(sprites["fire_center"],sprites["rocket_fuel_indicator_on"],sprites["rocket_fuel_indicator_off"],sprites["rocket"])
	g.hud = hud.NewHud(sprites["hud"], sprites["live"])
	g.preloadingProgressBar = hud.NewProgressbar(appWidth / 2 - 100, appHeight / 2 + 50, 200, 50, 0, 0, true)
	g.level = NewLevel()
	g.smoke = particles.NewSmoke(sprites["smoke"])
	g.explosion = particles.NewExplosion(sprites["explosion"])

	g.blinkingStars = []*objects.BlinkingStar{
		objects.NewBlinkingStar(sprites["blinking_star"]), 
		objects.NewBlinkingStar(sprites["blinking_star"]),
	}
	g.floors = []*objects.Floor{
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]), 
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]), 
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]), 
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]), 
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]), 
		objects.NewFloor(sprites["floor1"], sprites["lava_floor"], sprites["fire"]),
	}
	g.platforms = []*objects.Platform{objects.NewPlatform(sprites["platform"], sprites["pillar"]), objects.NewPlatform(sprites["platform"], sprites["pillar"])}	
	g.mainmenu = objects.NewMainmenu(sprites["mainmenu"])

	return nil

	
}

func (g *Game) putFuelIntoRocket() {
	if (g.rocket.FuelIndicatorItems < 4) {
		g.rocket.FuelIndicatorItems++
		sounds["rocket_fuel_drop"].Play()
		g.player.HasFuel = false
		g.restartFuel()
	} else {
		g.rocket.FuelIndicatorItems++
		sounds["rocket_fuel_drop"].Play()
		sounds["rocket_move"].Play()
		g.smoke.Creating = true
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

				g.platforms[indexPlatform].MoveTo(px * 210 + marginLeftPlatforms, py * 210 + marginTopPlatforms)
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
			g.floors[indexFloor].FloorType = objects.FloorNormal
		} else if string(char) == lavaFloorLevelCharacter {
			g.floors[indexFloor].FloorType = objects.FloorLava		
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

	px, py := randomPlatform.Position()

	g.fuel.SetFinalPositionIntoPlatform(px, py, platformWidthLanding)

}


func (g *Game) restartLevel() {
	g.restartPlayer()
	g.restartFuel()
}

func (g* Game) restartPlayer() {
	g.player.Restart(g.rocket.GetX() - 30)
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
		isGamePreloaded:	false,
		timeToPreloadGame:  0,
		status: 			GameStatusPreloadingGame,
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
