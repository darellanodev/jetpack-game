package main

import (
	"math/rand"
	"strconv"

	"github.com/darellanodev/jetpack-game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	
	if g.blackfader.IsActive() {
		g.blackfader.Update()
		if g.blackfader.IsMaxOpaque() {
			g.status = GameStatus(g.blackfader.GameStatus())
		}
	}

	if g.status == GameStatusResetGame {	
		g.level.Reinit()
		g.player.RestartLives()
		g.status = GameStatusInit
	}

	if g.status == GameStatusMainMenu {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.blackfader.Activate(int(GameStatusTravelingToPlanet))
		} else {
			return nil
		}
	}

	if g.status == GameStatusTravelingToPlanet {
		if g.planets.IsTraveling() {
			g.planets.Update()
			canSkip := !g.blackfader.IsActive() 
			if ebiten.IsKeyPressed(ebiten.KeyX) && canSkip{
				g.planets.Skip()
			}
			return nil
		} else {
			g.status = GameStatusInit
		}
	}

	if g.status == GameStatusInit {

		g.level.Next()
		g.placeLevelPlatforms()
		g.placeLevelFloors()
		g.placeLevelTrees()
		g.restartFuel()
		g.restartPlayer()
		g.rocket.RestartFuelItems()
		sounds["start"].Play()

		g.hud.SetMaxOxygenCapacity()
		g.hud.SetTitle(strconv.Itoa(g.level.number) + ": " + g.level.title)
		g.hud.SetLives(g.player.Lives)
		g.status = GameStatusLanding

		g.rocket.LandingSpeed = objects.RocketMaxSpeed
		g.smoke.MoveTo(g.rocket.GetX(), g.rocket.GetY())
		g.showSmokeTime = 0
		g.smoke.Creating = true
		g.explosion.Creating = false
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

		if g.rocket.GetY() < g.rocket.LandedY {

			g.rocket.Landing()

			g.smoke.MoveTo(g.rocket.GetX() + objects.RocketWidth/2, g.rocket.GetY() + objects.RocketHeight)
		} else {
			g.rocket.MoveTo(g.rocket.GetX(), g.rocket.LandedY)
			g.status = GameStatusPlaying
			g.smoke.Creating = false
			g.showSmokeTime = 0
		}
	}

	if g.status == GameStatusFinishingLevel {

		if g.rocket.GetY() > objects.StartRocketY {

			g.rocket.TakeOff()

			g.smoke.MoveTo(g.rocket.GetX() + objects.RocketWidth/2, g.rocket.GetY() + objects.RocketHeight)
		} else {

			if g.level.number == totalGameLevels {
				g.status = GameStatusGameComplete
			} else {

				g.rocket.MoveTo(g.rocket.GetX(), objects.StartRocketY)
				g.travelingTextTime = travelingTextMaxTime
				sounds["traveling"].Play()
				g.smoke.Creating = false
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

	if g.changeBlinkingStarsTime < objects.ChangeBlinkingStarsMaxTime {
		g.changeBlinkingStarsTime++
	} else {
		g.changeBlinkingStarsTime = 0
		for _, blinkingStar := range g.blinkingStars {
			if rand.Intn(100) < 20 {
				blinkingStar.MoveTo(rand.Intn(appWidth), rand.Intn(appHeight/3))
			} else {
				blinkingStar.MoveTo(50, 50) // do not show (behind the hud)
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

	if g.explosion.Creating && g.showExplosionTime < maxTimeToShowExplosion {
		g.showExplosionTime++

	}

	if g.showExplosionTime >= maxTimeToShowExplosion {
		g.explosion.Creating = false
		g.showExplosionTime = 0
	}

	if g.status == GameStatusPlaying {
		g.hud.Update()
	}

	if g.status == GameStatusGameOver || g.status == GameStatusGameComplete {
		
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			g.status = GameStatusResetGame
		}
		
		return nil
	}

	// check for collisions
	if g.status == GameStatusPlaying {

		var isCollidingPlayerWithEnemy, 
			isCollidingPlayerWithJumpingLavadrop,
			isCollidingPlayerWithLavaFloors bool

		// collision with enemy
		isCollidingPlayerWithEnemy = checkCollision(g.player, g.enemy)

		// collision with lava floors
		isCollidingPlayerWithLavaFloors = false
		for _, floor := range g.floors {
			if floor.IsLavaFloor() {

				isCollidingPlayerWithLavaFloor := checkCollision(g.player, floor)

				if isCollidingPlayerWithLavaFloor {
					isCollidingPlayerWithLavaFloors = true
				}

				if floor.FloorType == objects.FloorLavaWithDrops && !isCollidingPlayerWithJumpingLavadrop {
					// collision with jumping lava drop
					isCollidingPlayerWithJumpingLavadrop = checkCollision(g.player, floor.Lavadrop)
				}

			}
		}
		
		

		// collision with fuel
		isCollidingPlayerWithFuel := false
		if !g.fuel.Snaps {

			isCollidingPlayerWithFuel = checkCollision(g.player, g.fuel)
		}

		// collision with rocket when the player has the fuel
		if g.fuel.Snaps {

			isCollidingPlayerAndFuelWithRocket := checkCollision(g.player, g.rocket)

			if isCollidingPlayerAndFuelWithRocket {
				g.putFuelIntoRocket()
				isCollidingPlayerAndFuelWithRocket = false
			}
		}

		collidesWithSomethingBad := isCollidingPlayerWithEnemy || isCollidingPlayerWithLavaFloors || isCollidingPlayerWithJumpingLavadrop	

		if (collidesWithSomethingBad) && (g.player.InmuneToDamageTime == 0) {
			sounds["die"].Play()
			g.player.LostLive()
			g.player.InmuneToDamageTime = 200

			g.hud.SetLives(g.player.Lives)

			g.explosion.MoveTo(g.player.GetX(), g.player.GetY())
			g.explosion.Creating = true

			if g.player.Lives == 0 {
				g.status = GameStatusGameOver
			}
			g.restartLevel()
			return nil
		}

		if isCollidingPlayerWithFuel && !g.fuel.Snaps {
			g.fuel.Snaps = true
			isCollidingPlayerWithFuel = false
			g.player.HasFuel = true
			sounds["fuel_pick"].Play()
		}

	}

	if g.player.InmuneToDamageTime > 0 {
		g.player.InmuneToDamageTime--
	}

	if g.fuel.Snaps {
		g.fuel.MoveTo(g.player.HandsPosition())
	}

	return nil
}