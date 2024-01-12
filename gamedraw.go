package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {

	backgroundSpriteName := "background" + strconv.Itoa(g.level.number)
	drawNormalImage(screen, sprites[backgroundSpriteName], 0, 0)

	for _, blinkingStar := range g.blinkingStars {
		blinkingStar.Draw(screen, g.count)
	}

	if g.status == GameStatusPlaying || g.status == GameStatusPaused {
		g.player.Draw(screen, g.count)
	}

	if g.status != GameStatusTravelingToLevel && g.status != GameStatusFinishingLevel {
		g.enemy.Draw(screen)
		g.fuel.Draw(screen)
	}

	if g.showSmokeTime < maxTimeToShowSmoke {
		g.smoke.Draw(screen)
	}

	if g.explosion.creating && g.showExplosionTime < 50 {
		g.explosion.Draw(screen)
	}

	g.rocket.Draw(screen)

	for _, platform := range g.platforms {
		platform.Draw(screen)
	}

	//draw first lava floors (because then normal floors will be in front of lava floors and it will look better)
	for _, floor := range g.floors {
		if floor.floorType == FloorLava {
			floor.Draw(screen, g.count)
		}
	}
	//then draw normal floors
	for _, floor := range g.floors {
		if floor.floorType == FloorNormal {
			floor.Draw(screen, g.count)
		}
	}

	if g.status != GameStatusInit {
		g.hud.Draw(screen)
	}

	if g.status == GameStatusGameComplete {
		text.Draw(screen, "Game Complete!", mplusNormalFont, appWidth/3, appHeight/2, color.White)
		text.Draw(screen, "Thanks for playing, this game is", mplusNormalFont, appWidth/8, appHeight/2+offsetSecondTextLine, color.White)
		text.Draw(screen, "in an early stage of development", mplusNormalFont, appWidth/8, appHeight/2+offsetThirdTextLine, color.White)
		text.Draw(screen, "More stuff coming soon", mplusNormalFont, appWidth/4, appHeight/2+offsetFourthTextLine, color.White)
	}
	if g.status == GameStatusGameOver {
		text.Draw(screen, "Game Over", mplusNormalFont, appWidth/3, appHeight/2, color.White)
	}

	if g.status == GameStatusPaused {
		text.Draw(screen, "Paused", mplusNormalFont, appWidth/3, appHeight/2, color.White)
		text.Draw(screen, "Press P to continue", mplusNormalFont, appWidth/5, appHeight/2+offsetSecondTextLine, color.White)
	}

	if g.soundTextTime > 0 && g.status == GameStatusPlaying {
		if soundEnabled {
			text.Draw(screen, "Sound ON", mplusNormalFont, appWidth/3, appHeight/2, color.White)
		} else {
			text.Draw(screen, "Sound OFF", mplusNormalFont, appWidth/3, appHeight/2, color.White)
		}
		text.Draw(screen, "Press S to ON/OFF", mplusNormalFont, appWidth/3, appHeight/2+offsetSecondTextLine, color.White)
	}

	if g.status == GameStatusTravelingToLevel {
		text.Draw(screen, "Traveling to the", mplusNormalFont, appWidth/3, appHeight/2, color.White)
		text.Draw(screen, "next level...", mplusNormalFont, appWidth/3, appHeight/2+offsetSecondTextLine, color.White)
	}

	// msg := fmt.Sprintf("posX:%d posY:%d, fuelX:%d fuelY:%d, enemyX:%d enemyY:%d", g.player.x, g.player.y, g.fuel.x, g.fuel.y, g.enemy.x, g.enemy.y)
	// msg := fmt.Sprintf("debugMsg:%s soundEnabled: %v lives: %v",g.debugMsg, soundEnabled, g.player.lives)
	// ebitenutil.DebugPrint(screen, msg)

}