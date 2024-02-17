package objects

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type Mainmenu struct {
	img *ebiten.Image
}


func NewMainmenu(imgMainmenu *ebiten.Image) *Mainmenu {
	
	return &Mainmenu{
		img: imgMainmenu,
	}
}

func (m *Mainmenu) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, m.img, 0, 0)
	
}
