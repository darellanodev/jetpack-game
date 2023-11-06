package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player 			*Player
	enemy  			*Enemy
	fuel   			*Fuel
	rocket			*Rocket
	pause  			bool
	pauseTime 		int
	soundTime 		int
	pausePressed 	bool
	soundPressed	bool
	debugMsg 		string
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.MoveRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.MoveUp()
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
			g.pause = !g.pause
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

	if (g.pause) {
		return nil
	}

	g.player.Update()
	g.enemy.Update()
	g.fuel.Update()
	g.rocket.Update()

	isCollidingPlayerWithEnemy, _ := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.enemy.currentSprite, float64(g.enemy.x)/unit, float64(g.enemy.y)/unit)

	isCollidingPlayerWithFuel, debugMsg2 := isColliding(g.player.currentSprite, float64(g.player.x)/unit, float64(g.player.y)/unit, g.fuel.currentSprite, float64(g.fuel.x)/unit, float64(g.fuel.y)/unit)


	g.debugMsg = debugMsg2

	if (isCollidingPlayerWithEnemy){
		sounds["die"].Play()
		g.restartGame()
	}

	if isCollidingPlayerWithFuel && !g.fuel.snaps{
		g.fuel.snaps = true
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

func (g *Game) restartGame() {
	g.player.x = startPlayerX
	g.player.y = startPlayerY
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(sprites["background"], op)

	g.player.Draw(screen)
	g.enemy.Draw(screen)
	g.fuel.Draw(screen)
	g.rocket.Draw(screen)

	text.Draw(screen, "this is a sample", mplusNormalFont, 10, 120, color.White)

	// msg := fmt.Sprintf("posX:%d posY:%d, fuelX:%d fuelY:%d, enemyX:%d enemyY:%d", g.player.x, g.player.y, g.fuel.x, g.fuel.y, g.enemy.x, g.enemy.y)

	msg := fmt.Sprintf("debugMsg:%s soundEnabled: %v",g.debugMsg, soundEnabled)
	
	ebitenutil.DebugPrint(screen, msg)
	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}


func NewGame() *Game {
	g := &Game{
		player: &Player{
			x: 				startPlayerX,
			y: 				startPlayerY,
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
		},
		pause: 				false,
		pausePressed: 		false,
		pauseTime: 			0,
		soundPressed:		false,
		soundTime:			0,

	}

	return g
}

func (g *Game) Exit() error {
	//TODO: finish sounds and music
	return nil
}
