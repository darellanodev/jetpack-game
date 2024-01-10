package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func getSubImage (spriteSheet *ebiten.Image, frameWidth int, frameHeight int, spriteCount int, frameCount int, speed int ) *ebiten.Image {
	
	i := (spriteCount / speed) % frameCount
	sx, sy := (i * frameWidth), 0

	return spriteSheet.SubImage(image.Rect(sx, sy, sx + frameWidth, sy + frameHeight)).(*ebiten.Image)
}