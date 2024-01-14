package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var(
	mplusNormalFont font.Face
	mplusSmallFont font.Face
	tttf *sfnt.Font
)


func LoadFonts() error {

	var err error
	var fontBytes []byte

	fontBytes, err = loadStaticResource(assets, "assets/fonts/pressstart2p.ttf")

	if err != nil {
		return err
	}

	tttf, err = opentype.Parse(fontBytes)

	if err != nil {
		return err
	}

	const dpi = 72

	mplusNormalFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return err
	}
	mplusSmallFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return err
	}

	return nil
}
