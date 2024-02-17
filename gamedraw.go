package main

import (
	"image/color"
	"strconv"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g * Game) getWidthText(textLine string) int {
	width := 0
	
	if len(textLine) < 10 {
		width = appWidth / 2 - 100
	}
	if len(textLine) > 10 && len(textLine) < 17 {
		width = appWidth / 3
	}
	if len(textLine) > 17 && len(textLine) < 30 {
		width = appWidth / 4
	}
	if len(textLine) > 30 && len(textLine) < 40 {
		width = appWidth / 8
	}
	return width
}

func (g *Game) drawVerticalTexts(screen *ebiten.Image, textLines []string) {
	posY := 100
	heigth := appHeight / 2 - 150
	for _, textLine := range textLines {
		text.Draw(screen, textLine, lib.MplusNormalFont, g.getWidthText(textLine), heigth + posY, color.White)
		posY += 50
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.status == GameStatusPreloadingGame {

		
		g.preloadingProgressBar.Draw(screen)
		
		
		displayText := []string{
			"Loading",
			"Please wait ...",
		}

		g.drawVerticalTexts(screen, displayText)		

	} else if g.status == GameStatusMainMenu {
		g.mainmenu.Draw(screen)
	} else {


		backgroundSpriteName := "background" + strconv.Itoa(g.level.number)
		lib.DrawNormalImage(screen, sprites[backgroundSpriteName], 0, 0)
	
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
	
		if g.explosion.Creating && g.showExplosionTime < 50 {
			g.explosion.Draw(screen)
		}
	
		g.rocket.Draw(screen)
	
		for _, platform := range g.platforms {
			platform.Draw(screen)
		}
	
		//draw first lava floors (because then normal floors will be in front of lava floors and it will look better)
		for _, floor := range g.floors {
			if floor.FloorType == objects.FloorLava {
				floor.Draw(screen, g.count)
			}
		}
		//then draw normal floors
		for _, floor := range g.floors {
			if floor.FloorType == objects.FloorNormal {
				floor.Draw(screen, g.count)
			}
		}
	
		if g.status != GameStatusInit {
			g.hud.Draw(screen)
		}
	
		if g.status == GameStatusGameComplete {
	
			displayText := []string{
				"Game Complete!",
				"Thanks for playing, this game is",
				"in an early stage of development",
				"More stuff coming soon.",
				"Press ENTER to play again.",
			}
	
			g.drawVerticalTexts(screen, displayText)
	
		}
		if g.status == GameStatusGameOver {
	
			displayText := []string{
				"Game Over",
				"Press ENTER to play again.",
			}
	
			g.drawVerticalTexts(screen, displayText)		
		}
	
		if g.status == GameStatusPaused {
	
			displayText := []string{
				"Paused",
				"Press P to continue",
			}
	
			g.drawVerticalTexts(screen, displayText)		
	
		}
	
		if g.soundTextTime > 0 && g.status == GameStatusPlaying {
			soundSwichText := "SOUND OFF"
			
			if soundEnabled {
				soundSwichText = "SOUND ON"
			}
	
			displayText := []string{
				soundSwichText,
				"Press S to ON / OFF",
			}
	
			g.drawVerticalTexts(screen, displayText)
		}
	
		if g.status == GameStatusTravelingToLevel {
	
			displayText := []string{
				"Traveling to the",
				"next level...",
			}
	
			g.drawVerticalTexts(screen, displayText)	
		}



	}
	
	

	// msg := fmt.Sprintf("posX:%d posY:%d, fuelX:%d fuelY:%d, enemyX:%d enemyY:%d", g.player.x, g.player.y, g.fuel.x, g.fuel.y, g.enemy.x, g.enemy.y)
	// msg := fmt.Sprintf("debugMsg:%s soundEnabled: %v lives: %v",g.debugMsg, soundEnabled, g.player.lives)
	// ebitenutil.DebugPrint(screen, msg)

}