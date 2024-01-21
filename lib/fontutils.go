package lib

import (
	"embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

var(
	MplusNormalFont font.Face
	MplusSmallFont font.Face
	tttf *sfnt.Font
)


func LoadFonts(filesystem embed.FS) error {

	var err error
	var fontBytes []byte

	fontBytes, err = LoadStaticResource(filesystem, "assets/fonts/pressstart2p.ttf")

	if err != nil {
		return err
	}

	tttf, err = opentype.Parse(fontBytes)

	if err != nil {
		return err
	}

	const dpi = 72

	MplusNormalFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return err
	}
	MplusSmallFont, err = opentype.NewFace(tttf, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return err
	}

	return nil
}
