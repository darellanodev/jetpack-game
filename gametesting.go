package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Gametesting struct {
	player1 *Playertesting
	player2 *Playertesting
	enemy1  *Enemytesting
	enemy2  *Enemytesting
	dot1  	*Dottesting
	dot2  	*Dottesting
	fuel   	*Fuel
}

func (gt *Gametesting) Update() error {

	return nil
}

func (gt *Gametesting) Init() error {

	LoadSprites()
	LoadFonts()
	LoadSounds()

	return nil
}


func (gt *Gametesting) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(sprites["background"], op)

	gt.player1.Draw(screen)
	gt.player2.Draw(screen)
	gt.enemy1.Draw(screen)
	gt.enemy2.Draw(screen)
	gt.fuel.Draw(screen)
	gt.dot1.Draw(screen)
	gt.dot2.Draw(screen)

	text.Draw(screen, "this is a sample", mplusNormalFont, 10, 80, color.White)

	msg := fmt.Sprintf("player 1 (posX:%d posY:%d), fuelX:%d fuelY:%d", gt.player1.x, gt.player1.y, gt.fuel.x, gt.fuel.y)
	
	ebitenutil.DebugPrint(screen, msg)
	
}

func (gt *Gametesting) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}


func NewGametesting() *Gametesting {
	gt := &Gametesting{
		player1: &Playertesting{
			x: startPlayerX,
			y: startPlayerY,
		},
		player2: &Playertesting{
			x: startPlayerX + playerWidth * unit,
			y: startPlayerY + playerHeight * unit,
		},
		enemy1: &Enemytesting{
			x:     startEnemyX,
			y:     startEnemyY,
			up:    true,
			down:  false,
			left:  false,
			right: true,
		},
		enemy2: &Enemytesting{
			x:     startEnemyX + enemyWidth * unit,
			y:     startEnemyY + enemyHeight * unit,
			up:    true,
			down:  false,
			left:  false,
			right: true,
		},
		fuel: &Fuel{
			x: startFuelX,
			y: startFuelY,
			snaps: false,
		},
		dot1: &Dottesting{
			x: startEnemyX + ((enemyWidth / 2) * unit),
			y: startEnemyY + ((enemyHeight / 2) * unit),
		},
		dot2: &Dottesting{
			x: startPlayerX + ((playerWidth / 2) * unit),
			y: startPlayerY + ((playerHeight / 2) * unit),
		},

	}

	return gt
}

func (gt *Gametesting) Exit() error {
	//TODO: finish sounds and music
	return nil
}
