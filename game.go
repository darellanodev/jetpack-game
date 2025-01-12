package main

import (
	"math/rand"

	"github.com/darellanodev/jetpack-game/hud"
	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
	"github.com/darellanodev/jetpack-game/particles"
	"github.com/darellanodev/jetpack-game/scenes"
	"github.com/hajimehoshi/ebiten/v2"
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
	GameStatusTravelingToPlanet
)

type Game struct {
	player 			  		*objects.Player
	blackfader 		  		*objects.Blackfader
	enemy  			  		*objects.Enemy
	fuel   			  		*objects.Fuel
	rocket			  		*objects.Rocket
	mainmenu 				*objects.Mainmenu
	trees			  		[]*objects.Tree
	platforms		  		[]*objects.Platform
	floors			  		[]*objects.Floor
	blinkingStars	  		[]*objects.BlinkingStar
	smoke			  		*particles.ParticlesSystem
	explosion		  		*particles.ParticlesSystem
	level			  		*Level
	hud				  		*hud.Hud
	planets					*scenes.Planets
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

	rocketSprites := []*ebiten.Image{
		sprites["fire_center"],
		sprites["rocket_fuel_indicator_on"],
		sprites["rocket_fuel_indicator_off"],
		sprites["rocket"],
	}

	playerSprites := []*ebiten.Image{
		sprites["player_center"],
		sprites["fire_right"],
		sprites["fire_center"],
		sprites["player_walk_right_with_fuel"],
		sprites["player_walk_right"],
		sprites["player_right"],
		sprites["player_right_with_fuel"],
	}

	enemySprites := []*ebiten.Image{
		sprites["enemy1"],
		sprites["enemy1_closing_eyes"],
		sprites["enemy1_opening_eyes"],
	}

	planetsSprites := []*ebiten.Image{
		sprites["fire_planet"],
		sprites["green_planet"],
	}

	planetsBackgrounds := []*ebiten.Image{
		sprites["starfield1"],
		sprites["starfield2"],
	}

	planetsHudSprites := []*ebiten.Image{
		sprites["circle_planet_selector"],
	}

	floorSprites := []*ebiten.Image{
		sprites["vulcan_floor"],
		sprites["lava_floor"],
		sprites["fire"],
		sprites["lava_drop"],
	}

	platformSprites := []*ebiten.Image{
		sprites["platform"],
		sprites["pillar"],
	}

	treeSprites := []*ebiten.Image{
		sprites["fire_tree_01"],
		sprites["fire_tree_02"],
		sprites["fire_tree_03"],
		sprites["fire_tree_04"],
		sprites["fire_tree_05"],
	}

	hudSprites := []*ebiten.Image{
		sprites["hud"],
		sprites["live"],
	}

	g.blackfader = objects.NewBlackfader(int(GameStatusMainMenu), appWidth, appHeight)

	g.player = objects.NewPlayer(playerSprites)
	g.enemy = objects.NewEnemy(enemySprites)
	g.fuel = objects.NewFuel(sprites["fuel"], sprites["parachute"])
	g.rocket = objects.NewRocket(rocketSprites)
	
	g.hud = hud.NewHud(hudSprites)
	g.level = NewLevel()
	g.smoke = particles.NewSmoke(sprites["smoke"])
	g.explosion = particles.NewExplosion(sprites["explosion"])
	
	g.planets = scenes.NewPlanets(planetsSprites, rocketSprites, planetsHudSprites, planetsBackgrounds)
	g.planets.Init()

	g.blinkingStars = []*objects.BlinkingStar{
		objects.NewBlinkingStar(sprites["blinking_star"]), 
		objects.NewBlinkingStar(sprites["blinking_star"]),
	}
	g.floors = []*objects.Floor{
		objects.NewFloor(floorSprites), 
		objects.NewFloor(floorSprites), 
		objects.NewFloor(floorSprites), 
		objects.NewFloor(floorSprites), 
		objects.NewFloor(floorSprites), 
		objects.NewFloor(floorSprites),
	}
	treeY := appHeight - floorHeight - objects.TreeHeight 
	g.trees = []*objects.Tree{
		objects.NewTree(treeSprites[0], 200, treeY),
		objects.NewTree(treeSprites[1], 300, treeY),
		objects.NewTree(treeSprites[2], 400, treeY),
		objects.NewTree(treeSprites[3], 500, treeY),
		objects.NewTree(treeSprites[4], 600, treeY),
	}
	g.platforms = []*objects.Platform{
		objects.NewPlatform(platformSprites),
		objects.NewPlatform(platformSprites),
	}	
	g.mainmenu = objects.NewMainmenu(sprites["mainmenu"])

	return nil

	
}

func (g *Game) putFuelIntoRocket() {
	if g.rocket.FuelIndicatorItems < 4 {
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
			if string(char) == LevelCharacters["platform"] {

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
		floorType := objects.GetFloorType(string(char), LevelCharacters)
		g.floors[indexFloor].FloorType = floorType
		objects.InitializeFloor(g.floors[indexFloor], px, py)
		px += floorWidth
		indexFloor++
	}
}



func (g *Game) placeLevelTrees() {
	i := 0
	for _, floor := range g.floors {
		if floor.FloorType == objects.FloorNormal {
			if i < len(g.trees) {
				randValue := rand.Intn(100)
				floorPosX, _ := floor.Position()
				rocketX := g.rocket.GetX()
				isRocketFloor := (floorPosX <= rocketX) && ((floorPosX + floorWidth) >= rocketX)
				posX := objects.CalculateTreePositionX(floorPosX, randValue, isRocketFloor)
				posY := appHeight - floorHeight - objects.TreeHeight
				g.trees[i].MoveTo(posX, posY)
				i++
			}
		}
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
		status: 			GameStatusMainMenu,
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
