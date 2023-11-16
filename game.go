package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type GameStatus int

const (
	GameStatusPlaying GameStatus = iota
	GameStatusPaused
	GameStatusGameOver
	GameStatusInit
)

type Game struct {
	player 			*Player
	enemy  			*Enemy
	fuel   			*Fuel
	rocket			*Rocket
	platforms		[]*Platform
	hud				*Hud
	pauseTime 		int
	soundTime 		int
	pausePressed 	bool
	soundPressed	bool
	debugMsg 		string
	status			GameStatus
}

func (g *Game) Update() error {

	if (g.status == GameStatusInit) {
		g.restartFuel()

		g.status = GameStatusPlaying
	}

	if (g.status == GameStatusPlaying) {

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
		}
		
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		if (g.pauseTime == 0) {
			
			if (g.status != GameStatusPaused) {
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
	if (g.soundPressed && g.soundTime > 0) {
		g.soundTime--
	}

	if (g.status == GameStatusPaused) {
		return nil
	}

	g.player.Update()
	g.enemy.Update()
	g.fuel.Update()
	g.rocket.Update()

	// collision with enemy
	// isCollidingPlayerWithEnemy, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.enemy.currentSprite, float64(g.enemy.x)/unit, float64(g.enemy.y)/unit)
	isCollidingPlayerWithEnemy, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.enemy.currentSprite, float64(g.enemy.x)/unit, float64(g.enemy.y)/unit)

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

	if (isCollidingPlayerWithEnemy){
		sounds["die"].Play()
		g.player.LostLive()
		if (g.player.lives == 0) {
			g.status = GameStatusGameOver
		}
		g.restartGame()
		return nil
	} 
	
	if isCollidingPlayerWithFuel && !g.fuel.snaps{
		g.fuel.snaps = true
		isCollidingPlayerWithFuel = false
		sounds["fuel_pick"].Play()
	}

	if g.fuel.snaps {
		g.fuel.MoveTo(g.player.Position())
	}

	return nil
}

func (g *Game) Init() error {

	LoadSprites()
	LoadFonts()
	LoadSounds()

	sounds["start"].Play()
	return nil
}

func (g *Game) putFuelIntoRocket() {
	if (g.rocket.fuelIndicatorItems < 5) {
		g.rocket.fuelIndicatorItems++
		g.restartFuel()
	} 
	//TODO: else: level completed!

}

func (g *Game) restartFuel() {
	g.fuel.snaps = false

	randomIndex := rand.Intn(len(g.platforms))
	randomPlatform := g.platforms[randomIndex]

	px, py := randomPlatform.position()
	g.fuel.MoveTo(px + rand.Intn(3000) + 300 , py - 300)

}

func (g *Game) restartGame() {
	g.player.x = startPlayerX
	g.player.y = startPlayerY
	g.restartFuel()
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(sprites["background"], op)

	if (g.status == GameStatusPlaying || g.status == GameStatusPaused) {
		g.player.Draw(screen)
	}
	g.enemy.Draw(screen)
	g.fuel.Draw(screen)
	g.rocket.Draw(screen)
	for _, platform := range g.platforms {
		platform.Draw(screen)
	}
	g.hud.Draw(screen)

	if (g.status == GameStatusGameOver) {
		text.Draw(screen, "Game Over", mplusNormalFont, 220, 220, color.White)
	}

	// msg := fmt.Sprintf("posX:%d posY:%d, fuelX:%d fuelY:%d, enemyX:%d enemyY:%d", g.player.x, g.player.y, g.fuel.x, g.fuel.y, g.enemy.x, g.enemy.y)
	// msg := fmt.Sprintf("debugMsg:%s soundEnabled: %v lives: %v",g.debugMsg, soundEnabled, g.player.lives)
	// ebitenutil.DebugPrint(screen, msg)
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}


func NewGame() *Game {
	g := &Game{
		player: &Player{
			x: 				startPlayerX,
			y: 				startPlayerY,
			lives:			3,
			currentSprite: 	nil,
		},
		enemy: &Enemy{
			x:     			startEnemyX,
			y:     			startEnemyY,
			up:    			true,
			down:  			false,
			left:  			false,
			right: 			true,
			currentSprite: 	nil,
		},
		fuel: &Fuel{
			x: 				startFuelX,
			y: 				startFuelY,
			currentSprite: 	nil,
			snaps: 			false,
		},
		rocket: &Rocket{
			x: 				startRocketX,
			y: 				startRocketY,
			currentSprite: 	nil,
			snaps: 			false,
			fuelIndicatorItems: 		0,
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
		hud: &Hud{
			x: 0,
			y: 0,
		},
		pausePressed: 		false,
		pauseTime: 			0,
		soundPressed:		false,
		soundTime:			0,
		status: 			GameStatusInit,

	}

	return g
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}
