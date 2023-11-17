package main

import (
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)


var mplusNormalFont font.Face
var mplusHudFont font.Face
var tttf *sfnt.Font

func LoadFonts() {

	fontBytes, err := os.ReadFile("assets/fonts/pressstart2p.ttf")
	if err != nil {
		panic(err)
	}
	tttf, _ = opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	mplusNormalFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusHudFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}
