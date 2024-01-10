package main

import "github.com/hajimehoshi/ebiten/v2"

func drawNormalImage(screen *ebiten.Image, img *ebiten.Image, posX int, posY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(posX), float64(posY))
	screen.DrawImage(img, op)
}

func drawHorizontalFlippedImage(screen *ebiten.Image, img *ebiten.Image, imageWidth int, posX int, posY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(posX)+float64(imageWidth), float64(posY))
	screen.DrawImage(img, op)
}

