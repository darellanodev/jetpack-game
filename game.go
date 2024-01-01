package main

import (
	"image"
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
)

type Game struct {
	player 			  		*Player
	enemy  			  		*Enemy
	fuel   			  		*Fuel
	rocket			  		*Rocket
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
	debugMsg 		  		string
	status			  		GameStatus
	travelingTextTime 		int
	count			  		int
}

func (g *Game) Update() error {

	if (g.status == GameStatusInit) {

		g.smoke.SetImg(sprites["smoke"])
		g.explosion.SetImg(sprites["explosion"])

		g.level.Next()
		g.placeLevelPlatforms()
		g.placeLevelFloors()
		g.restartFuel()
		g.restartPlayer()
		g.rocket.restartFuelItems()
		sounds["start"].Play()
		
		g.hud.oxygen = maxOxygenCapacity
		g.hud.setTitle(strconv.Itoa(g.level.number) + ": " + g.level.title)
		g.hud.setLives(g.player.lives)
		g.status = GameStatusLanding

		g.rocket.landingSpeed = rocketMaxSpeed
		g.smoke.MoveTo(g.rocket.x, startPlayerY)
		g.showSmokeTime = 0
		g.smoke.creating = true
		g.explosion.creating = false
		g.showExplosionTime = 0
		g.changeBlinkingStarsTime = 0

	}

	if (ebiten.IsKeyPressed(ebiten.KeyP) && (g.status == GameStatusPlaying || g.status == GameStatusPaused)) {
		if (g.pauseTime == 0) {
			
			if (g.status == GameStatusPlaying) {
				g.status = GameStatusPaused
			} else {
				g.status = GameStatusPlaying
			}
			g.pausePressed = true
			g.pauseTime = 20
		}
	}

	if (g.pausePressed && g.pauseTime > 0) {
		g.pauseTime--
	}

	if (g.status == GameStatusPaused) {
		return nil
	}

	if (g.status == GameStatusLanding) {
		
		if (g.rocket.y < g.rocket.landedY) {
			g.rocket.MoveTo(g.rocket.x, g.rocket.y + (10) * int(g.rocket.landingSpeed))
			g.rocket.landingSpeed -= 0.146
		} else {
			g.rocket.MoveTo(g.rocket.x, g.rocket.landedY)
			g.status = GameStatusPlaying
			g.smoke.creating = false
			g.showSmokeTime = 0
		}
	}

	if (g.status == GameStatusFinishingLevel) {
		
		if (g.rocket.y > startRocketY) {
			g.rocket.MoveTo(g.rocket.x, g.rocket.y - (10) * int(g.rocket.landingSpeed))
			g.rocket.landingSpeed += 0.30
		} else {

			if (g.level.number == totalGameLevels) {
				g.status = GameStatusGameComplete
			} else {
				
				g.rocket.MoveTo(g.rocket.x, startRocketY)
				g.travelingTextTime = travelingTextMaxTime
				sounds["traveling"].Play()
				g.smoke.creating = false
				g.status = GameStatusTravelingToLevel
			}
		}
	}

	if (g.status == GameStatusTravelingToLevel) {
		g.travelingTextTime--
		if (g.travelingTextTime == 0) {
			g.travelingTextTime = travelingTextMaxTime
			g.status = GameStatusInit
		}

	}


	if (g.status == GameStatusPlaying) {

		g.count++

		if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
			g.player.MoveRight()
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
			g.player.MoveLeft()
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
			g.player.MoveUp()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if (g.soundTime == 0) {
			soundEnabled = !soundEnabled
			g.soundPressed = true
			g.soundTime = 20
			g.soundTextTime = 200
		}
		
	}

	if (g.soundTextTime > 0) {
		g.soundTextTime--
	}

	if (g.soundPressed && g.soundTime > 0) {
		g.soundTime--
	}

	g.player.Update()
	g.enemy.Update()
	g.fuel.Update()

	if (g.changeBlinkingStarsTime < changeBlinkingStarsMaxTime) {
		g.changeBlinkingStarsTime++
	} else {
		g.changeBlinkingStarsTime = 0
		for _, blinkingStar := range g.blinkingStars {
			if (rand.Intn(100) < 20) {
				blinkingStar.MoveTo(rand.Intn(20000), rand.Intn(7000))
			} else {
				blinkingStar.MoveTo(50, 50) // dont show (behind the hud)
			}
		}
	}


	for _, floor := range g.floors {
		floor.Update()
	}

	g.smoke.MoveTo(500,100 + g.rocket.y / 32)
	g.smoke.UpdateExpanded()
	g.explosion.UpdateExpanded()

	if (g.showSmokeTime < maxTimeToShowSmoke) {
		g.showSmokeTime++
		
	}

	if (g.explosion.creating && g.showExplosionTime < maxTimeToShowExplosion) {
		g.showExplosionTime++
		
	}

	if (g.showExplosionTime >= maxTimeToShowExplosion) {
		g.explosion.creating = false
		g.showExplosionTime = 0
	}

	if (g.status == GameStatusPlaying){
		g.hud.Update()
	}

	if (g.status == GameStatusGameOver || g.status == GameStatusGameComplete) {
		return nil
	}	
	// collision with enemy
	isCollidingPlayerWithEnemy, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.enemy.currentSprite, float64(g.enemy.x)/unit, float64(g.enemy.y)/unit)

	// collision with lava floors
	isCollidingPlayerWithLavaFloors := false
	imgLavaFloor := sprites["lava_floor"].SubImage(image.Rect(0, 0, lavaFloorFrameWidth, lavaFloorFrameHeight)).(*ebiten.Image)
	for _, floor := range g.floors {
		if (floor.floorType == FloorLava) {
			isCollidingPlayerWithLavaFloor, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, imgLavaFloor, float64(floor.x)/unit, float64(floor.y)/unit)
			if (isCollidingPlayerWithLavaFloor) {
				isCollidingPlayerWithLavaFloors = true
			}
		}
	}

	// collision with fuel
	isCollidingPlayerWithFuel := false
	debugMsg2 := ""
	if (!g.fuel.snaps) {
		isCollidingPlayerWithFuel, debugMsg2 = isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.fuel.currentSprite, float64(g.fuel.x)/unit, float64(g.fuel.y)/unit)
	}
	g.debugMsg = debugMsg2

	// collision with rocket when the player has the fuel
	if (g.fuel.snaps) {
		isCollidingPlayerAndFuelWithRocket, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.rocket.currentSprite, float64(g.rocket.x)/unit, float64(g.rocket.y)/unit)

		if (isCollidingPlayerAndFuelWithRocket) {
			g.putFuelIntoRocket()
			isCollidingPlayerAndFuelWithRocket = false
		}
	}

	if (g.player.inmuneToDamageTime > 0) {
		g.player.inmuneToDamageTime--
	}

	if ((isCollidingPlayerWithEnemy || isCollidingPlayerWithLavaFloors) && g.player.inmuneToDamageTime == 0){
		sounds["die"].Play()
		g.player.LostLive()
		g.player.inmuneToDamageTime = 200

		g.hud.setLives(g.player.lives)
		
		g.explosion.MoveTo(g.player.x / 32, g.player.y / 32)
		g.explosion.creating = true

		if (g.player.lives == 0) {
			g.status = GameStatusGameOver
		}
		g.restartGame()
		return nil
	} 
	
	if isCollidingPlayerWithFuel && !g.fuel.snaps{
		g.fuel.snaps = true
		isCollidingPlayerWithFuel = false
		g.player.hasFuel = true
		sounds["fuel_pick"].Play()
	}

	if g.fuel.snaps {
		g.fuel.MoveTo(g.player.HandsPosition())
	}

	return nil
}

func (g *Game) Init() error {

	LoadSprites()
	LoadFonts()
	LoadSounds()

	return nil
}

func (g *Game) putFuelIntoRocket() {
	if (g.rocket.fuelIndicatorItems < 4) {
		g.rocket.fuelIndicatorItems++
		sounds["rocket_fuel_drop"].Play()
		g.player.hasFuel = false
		g.restartFuel()
	} else {
		g.rocket.fuelIndicatorItems++
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
			px ++
			if char == '1' {
				g.platforms[indexPlatform].x = px * 4000 + marginLeftPlatforms
				// fmt.Println("px", g.platforms[indexPlatform].x)

				g.platforms[indexPlatform].y = py * 3000 + marginTopPlatforms
				// fmt.Println("py", g.platforms[indexPlatform].y)
				indexPlatform++
			}
		}
		py++
	}
}

func (g *Game) placeLevelFloors() {

	px := 0
	py := floorY

	indexFloor := 0
	for _, char := range g.level.floorPlaces {
		if char == '1' {
			g.floors[indexFloor].floorType = FloorNormal
		} else if char == '2' {
			g.floors[indexFloor].floorType = FloorLava		
		}
		g.floors[indexFloor].MoveTo(px,py)
		g.floors[indexFloor].InitFloor()
		px += floorWidth
		indexFloor++
	}

}

func (g *Game) restartFuel() {
	g.fuel.snaps = false

	randomIndex := rand.Intn(len(g.platforms))
	randomPlatform := g.platforms[randomIndex]

	px, py := randomPlatform.position()

	fx := px + rand.Intn(platformWidthLanding) + 300
	fy := py - 300

	g.fuel.SetFinalPosition(fx, fy)

	g.fuel.MoveTo(fx, 0)

}

func (g *Game) restartPlayer() {
	g.player.x = g.rocket.x - 300
	g.player.y = startPlayerY
	g.player.hasFuel = false
}

func (g *Game) restartGame() {
	g.restartPlayer()
	g.restartFuel()
}

func (g *Game) Draw(screen *ebiten.Image) {

	backgroundSpriteName := "background" + strconv.Itoa(g.level.number)
	NewGame().drawNormalImage(screen, sprites[backgroundSpriteName], 0, 0)

	for _, blinkingStar := range g.blinkingStars {
		blinkingStar.Draw(screen, g.count)
	}

	if (g.status == GameStatusPlaying || g.status == GameStatusPaused) {
		g.player.Draw(screen, g.count)
	}

	if (g.status != GameStatusTravelingToLevel && g.status != GameStatusFinishingLevel) {
		g.enemy.Draw(screen)
		g.fuel.Draw(screen)
	}

	if (g.showSmokeTime < maxTimeToShowSmoke) {
		g.smoke.Draw(screen)
	}

	if (g.explosion.creating && g.showExplosionTime < 50) {
		g.explosion.Draw(screen)
	}

	g.rocket.Draw(screen)

	for _, platform := range g.platforms {
		platform.Draw(screen)
	}



	//draw first lava floors (because then normal floors will be in front of lava floors and it will look better)
	for _, floor := range g.floors {
		if (floor.floorType == FloorLava) {
			floor.Draw(screen, g.count)
		}
	}
	//then draw normal floors
	for _, floor := range g.floors {
		if (floor.floorType == FloorNormal) {
			floor.Draw(screen, g.count)
		}		
	}

	if (g.status != GameStatusInit) {
		g.hud.Draw(screen)
	}

	if (g.status == GameStatusGameComplete) {
		text.Draw(screen, "Game Complete!", mplusNormalFont, 150, 220, color.White)
		text.Draw(screen, "Thanks for playing, this game is in an early stage of development", mplusSmallFont, 50, 300, color.White)
		text.Draw(screen, "More stuff coming soon", mplusSmallFont, 250, 350, color.White)
	}
	if (g.status == GameStatusGameOver) {
		text.Draw(screen, "Game Over", mplusNormalFont, 220, 220, color.White)
	}

	if (g.status == GameStatusPaused) {
		text.Draw(screen, "Paused", mplusNormalFont, 240, 220, color.White)
		text.Draw(screen, "Press P to continue", mplusNormalFont, 90, 260, color.White)
	}

	if (g.soundTextTime > 0 && g.status == GameStatusPlaying) {
		if (soundEnabled) {
			text.Draw(screen, "Sound ON", mplusNormalFont, 220, 220, color.White)
		} else {
			text.Draw(screen, "Sound OFF", mplusNormalFont, 220, 220, color.White)
		}
		text.Draw(screen, "Press S to ON/OFF", mplusNormalFont, 110, 260, color.White)
	}

	if (g.status == GameStatusTravelingToLevel) {
		text.Draw(screen, "Traveling to the", mplusNormalFont, 120, 220, color.White)
		text.Draw(screen, "next level...", mplusNormalFont, 160, 260, color.White)
	}
	

	// msg := fmt.Sprintf("posX:%d posY:%d, fuelX:%d fuelY:%d, enemyX:%d enemyY:%d", g.player.x, g.player.y, g.fuel.x, g.fuel.y, g.enemy.x, g.enemy.y)
	// msg := fmt.Sprintf("debugMsg:%s soundEnabled: %v lives: %v",g.debugMsg, soundEnabled, g.player.lives)
	// ebitenutil.DebugPrint(screen, msg)
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return appWidth, appHeight
}


func NewGame() *Game {
	g := &Game{
		player: &Player{
			x: 				    startPlayerX,
			y: 				    startPlayerY,
			lives:			    3,
			currentSprite: 	    nil,
			PlayerStatus:       Center,
			timeToIdle:		    maxTimeToIdle,
			hasFuel:		    false,
			inmuneToDamageTime: 0,
		},
		enemy: &Enemy{
			x:     				startEnemyX,
			y:     				startEnemyY,
			up:    				true,
			down:  				false,
			left:  				false,
			right: 				true,
			currentSprite: 		nil,
			timeToCloseEyesMax: 200,
			timeToCloseEyes: 	0,
			spriteCount: 		0,
			spriteSpeed: 		20,
			isClosingEyes: 		false,
		},
		fuel: &Fuel{
			x: 				startFuelX,
			y: 				startFuelY,
			currentSprite: 	nil,
			snaps: 			false,
		},
		rocket: &Rocket{
			x: 					startRocketX,
			y: 					startRocketY,
			landedY:			landedRocketY,
			landingSpeed: 		rocketMaxSpeed,
			currentSprite: 		nil,
			snaps: 				false,
			fuelIndicatorItems: startRocketFuelItems,
		},
		platforms: []*Platform{
			{
				x: 2000,
				y: 4000,
			},
			{
				x: 5000,
				y: 6000,
			},
		},
		floors: []*Floor{
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 0,
			},
		},
		blinkingStars: []*BlinkingStar{
			{
				x: 220,
				y: 110,
			},
			{
				x: 330,
				y: 1410,
			},
			
		},
		hud: &Hud{
			x: 0,
			y: 0,
			oxygen: maxOxygenCapacity,
			oxygenTimeToConsume: maxOxygenTimeToConsume,
		},
		level: &Level{
			number: startingLevel,
			title:  "",
		},
		smoke: &ParticlesSystem{
			particles: nil,
			posX: 100,
			posY: 100,
			creating: false,
		},
		explosion: &ParticlesSystem{
			particles: nil,
			posX: 100,
			posY: 100,
			creating: false,
		},
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
