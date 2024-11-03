package scenes

import (
	"image/color"
	_ "image/png"

	"github.com/darellanodev/jetpack-game/hud"
	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Planets struct {
	imgFirePlanet *ebiten.Image
	imgGreenPlanet *ebiten.Image
	imgCirclePlanetSelector *ebiten.Image
	imgBackground *ebiten.Image
	timeTraveling int
	y int
	rocket *objects.Rocket
	typeWriter *hud.TypeWriter
}

const (
	maxTimeTravelingToPlanet = 550
	timeToStartSizeReduction = 200
	initialFirePlanetY = -100
	initialGreenPlanetY = -300
	planetWidth = 200
	circlePlanetSelectorWidth = 220
)

func NewPlanets(planetsSprites []*ebiten.Image, rocketSprites []*ebiten.Image, planetsHudSprites []*ebiten.Image, planetsBackgrounds []*ebiten.Image) *Planets {
	
	return &Planets{
		imgFirePlanet: planetsSprites[0],
		imgGreenPlanet: planetsSprites[1],
		imgCirclePlanetSelector: planetsHudSprites[0],
		imgBackground: planetsBackgrounds[0],
		timeTraveling: 0,
		y: 0,
		rocket: objects.NewRocket(rocketSprites),
		typeWriter: hud.NewTypeWriter(),
	}
}


func (p *Planets) drawCirclePlanetSelector(screen *ebiten.Image) {

	rotation := float64(p.y * 2)
	offset := (circlePlanetSelectorWidth - planetWidth) / 2
	x := 300 - offset
	y := p.y - (planetWidth / 2) - offset
	
	lib.DrawRotateImage(screen, p.imgCirclePlanetSelector, x, y, rotation)
}

func (p *Planets) drawFirePlanet(screen *ebiten.Image) {

	rotation := float64(p.y / 4)
	y := initialFirePlanetY + p.y
	textToDisplay := ""

	if y > 100 {
		textToDisplay = p.typeWriter.Write("Approaching the fire planet")
	}

	text.Draw(screen, textToDisplay, lib.MplusNormalFont, 320, y - 20, color.White)
	lib.DrawRotateImage(screen, p.imgFirePlanet, 300, y, rotation)
}

func (p *Planets) drawGreenPlanet(screen *ebiten.Image) {

	rotation := float64(p.y / 8)
	y := initialGreenPlanetY + p.y / 2

	lib.DrawRotateImage(screen, p.imgGreenPlanet, 500, y, rotation)

}

func (p *Planets) drawBackground(screen *ebiten.Image) {

	y := initialGreenPlanetY + p.y / 4
	lib.DrawNormalImage(screen, p.imgBackground, 0, y)
}

func (p *Planets) Draw(screen *ebiten.Image, canSkip bool) {
	p.drawBackground(screen)
	p.drawCirclePlanetSelector(screen)
	p.drawFirePlanet(screen)
	p.drawGreenPlanet(screen)
	p.rocket.Draw(screen)
	if (canSkip) {
		text.Draw(screen, "press X to skip", lib.MplusNormalFont, 620, 740, color.White)
	}
}

func (p *Planets) Init() {

	p.rocket.SetFireAllways()
	
}

func (p *Planets) Update() {

	if p.timeTraveling < maxTimeTravelingToPlanet {
		p.timeTraveling++
		p.y += 1
		p.rocket.MoveTo(370, 700 - p.y / 5)
		if p.timeTraveling > timeToStartSizeReduction {
			p.rocket.ReduceScale()
		}
		
	}
	
}

func (p *Planets) IsTraveling() bool {
	return p.timeTraveling < maxTimeTravelingToPlanet
}

func (p *Planets) Skip() {
	p.timeTraveling = maxTimeTravelingToPlanet
}
