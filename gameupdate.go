package main

import (
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {

	if g.status == GameStatusInit {

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
		g.smoke.MoveTo(g.rocket.x, g.rocket.y)
		g.showSmokeTime = 0
		g.smoke.creating = true
		g.explosion.creating = false
		g.showExplosionTime = 0
		g.changeBlinkingStarsTime = 0

	}

	if ebiten.IsKeyPressed(ebiten.KeyP) && (g.status == GameStatusPlaying || g.status == GameStatusPaused) {
		if g.pauseTime == 0 {

			if g.status == GameStatusPlaying {
				g.status = GameStatusPaused
			} else {
				g.status = GameStatusPlaying
			}
			g.pausePressed = true
			g.pauseTime = 20
		}
	}

	if g.pausePressed && g.pauseTime > 0 {
		g.pauseTime--
	}

	if g.status == GameStatusPaused {
		return nil
	}

	if g.status == GameStatusLanding {

		if g.rocket.y < g.rocket.landedY {
			g.rocket.MoveTo(g.rocket.x, g.rocket.y+(2)*int(g.rocket.landingSpeed))
			g.rocket.landingSpeed -= rocketAcceleration
			g.smoke.MoveTo(g.rocket.x+rocketWidth/2, g.rocket.y+rocketHeight)
		} else {
			g.rocket.MoveTo(g.rocket.x, g.rocket.landedY)
			g.status = GameStatusPlaying
			g.smoke.creating = false
			g.showSmokeTime = 0
		}
	}

	if g.status == GameStatusFinishingLevel {

		if g.rocket.y > startRocketY {
			g.rocket.MoveTo(g.rocket.x, g.rocket.y-(2)*int(g.rocket.landingSpeed))
			g.rocket.landingSpeed += rocketAcceleration
			g.smoke.MoveTo(g.rocket.x+rocketWidth/2, g.rocket.y+rocketHeight)
		} else {

			if g.level.number == totalGameLevels {
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

	if g.status == GameStatusTravelingToLevel {
		g.travelingTextTime--
		if g.travelingTextTime == 0 {
			g.travelingTextTime = travelingTextMaxTime
			g.status = GameStatusInit
		}

	}

	g.count++

	if g.status == GameStatusPlaying {

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
		if g.soundTime == 0 {
			soundEnabled = !soundEnabled
			g.soundPressed = true
			g.soundTime = 20
			g.soundTextTime = 200
		}

	}

	if g.soundTextTime > 0 {
		g.soundTextTime--
	}

	if g.soundPressed && g.soundTime > 0 {
		g.soundTime--
	}

	g.player.Update()
	g.enemy.Update()
	g.fuel.Update()

	if g.changeBlinkingStarsTime < changeBlinkingStarsMaxTime {
		g.changeBlinkingStarsTime++
	} else {
		g.changeBlinkingStarsTime = 0
		for _, blinkingStar := range g.blinkingStars {
			if rand.Intn(100) < 20 {
				blinkingStar.MoveTo(rand.Intn(appWidth), rand.Intn(appHeight/3))
			} else {
				blinkingStar.MoveTo(50, 50) // dont show (behind the hud)
			}
		}
	}

	for _, floor := range g.floors {
		floor.Update()
	}

	g.smoke.UpdateExpanded()
	g.explosion.UpdateExpanded()

	if g.showSmokeTime < maxTimeToShowSmoke {
		g.showSmokeTime++

	}

	if g.explosion.creating && g.showExplosionTime < maxTimeToShowExplosion {
		g.showExplosionTime++

	}

	if g.showExplosionTime >= maxTimeToShowExplosion {
		g.explosion.creating = false
		g.showExplosionTime = 0
	}

	if g.status == GameStatusPlaying {
		g.hud.Update()
	}

	if g.status == GameStatusGameOver || g.status == GameStatusGameComplete {
		return nil
	}

	// check for collisions
	if g.status == GameStatusPlaying {

		// collision with enemy
		isCollidingPlayerWithEnemy := checkCollision(g.player, g.enemy)

		// collision with lava floors
		isCollidingPlayerWithLavaFloors := false
		for _, floor := range g.floors {
			if floor.floorType == FloorLava {

				isCollidingPlayerWithLavaFloor := checkCollision(g.player, floor)

				if isCollidingPlayerWithLavaFloor {
					isCollidingPlayerWithLavaFloors = true
				}
			}
		}

		// collision with fuel
		isCollidingPlayerWithFuel := false
		if !g.fuel.snaps {

			isCollidingPlayerWithFuel = checkCollision(g.player, g.fuel)
		}

		// collision with rocket when the player has the fuel
		if g.fuel.snaps {

			isCollidingPlayerAndFuelWithRocket := checkCollision(g.player, g.rocket)

			if isCollidingPlayerAndFuelWithRocket {
				g.putFuelIntoRocket()
				isCollidingPlayerAndFuelWithRocket = false
			}
		}

		if (isCollidingPlayerWithEnemy || isCollidingPlayerWithLavaFloors) && g.player.inmuneToDamageTime == 0 {
			sounds["die"].Play()
			g.player.LostLive()
			g.player.inmuneToDamageTime = 200

			g.hud.setLives(g.player.lives)

			g.explosion.MoveTo(g.player.x, g.player.y)
			g.explosion.creating = true

			if g.player.lives == 0 {
				g.status = GameStatusGameOver
			}
			g.restartGame()
			return nil
		}

		if isCollidingPlayerWithFuel && !g.fuel.snaps {
			g.fuel.snaps = true
			isCollidingPlayerWithFuel = false
			g.player.hasFuel = true
			sounds["fuel_pick"].Play()
		}

	}

	if g.player.inmuneToDamageTime > 0 {
		g.player.inmuneToDamageTime--
	}

	if g.fuel.snaps {
		g.fuel.MoveTo(g.player.HandsPosition())
	}

	return nil
}