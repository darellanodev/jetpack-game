package lib

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

func getImageCenter(img *ebiten.Image) (float64, float64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	return halfW, halfH 
}

func DrawNormalImage(screen *ebiten.Image, img *ebiten.Image, posX int, posY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(posX), float64(posY))
	screen.DrawImage(img, op)
}

func DrawRotateImage(screen *ebiten.Image, img *ebiten.Image, posX int, posY int, degrees float64) {
	halfW, halfH := getImageCenter(img)
	
	op := &colorm.DrawImageOptions{}
	cm := colorm.ColorM{}	

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(degrees * math.Pi / 180.0)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(float64(posX), float64(posY))

	colorm.DrawImage(screen, img, cm, op)
}

func DrawHorizontalFlippedImage(screen *ebiten.Image, img *ebiten.Image, imageWidth int, posX int, posY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(posX)+float64(imageWidth), float64(posY))
	screen.DrawImage(img, op)
}

func GetSubImage (spriteSheet *ebiten.Image, frameWidth int, frameHeight int, spriteCount int, frameCount int, speed int ) *ebiten.Image {
	
	i := (spriteCount / speed) % frameCount
	sx, sy := (i * frameWidth), 0

	return spriteSheet.SubImage(image.Rect(sx, sy, sx + frameWidth, sy + frameHeight)).(*ebiten.Image)
}