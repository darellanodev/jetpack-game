package main

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var(
	mplusNormalFont font.Face
	mplusSmallFont font.Face
	tttf *sfnt.Font
)


func LoadFonts() {

	var err error

	fontBytes := loadStaticResource(assets, "assets/fonts/pressstart2p.ttf")

	tttf, err = opentype.Parse(fontBytes)

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
	mplusSmallFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}
