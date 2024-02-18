package scenes

import (
	_ "image/png"

	"github.com/darellanodev/jetpack-game/lib"
	"github.com/darellanodev/jetpack-game/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

type Planets struct {
	imgFirePlanet *ebiten.Image
	imgGreenPlanet *ebiten.Image
	timeTraveling int
	y int
	rocket *objects.Rocket
}

const (
	maxTimeTravelingToPlanet = 550
	initialFirePlanetY = -100
	initialGreenPlanetY = -300
)

func NewPlanets(planetsSprites []*ebiten.Image, rocketSprites []*ebiten.Image) *Planets {
	
	return &Planets{
		imgFirePlanet: planetsSprites[0],
		imgGreenPlanet: planetsSprites[1],
		timeTraveling: 0,
		y: 0,
		rocket: objects.NewRocket(rocketSprites),
	}
}

func (p *Planets) Draw(screen *ebiten.Image) {

	lib.DrawNormalImage(screen, p.imgFirePlanet, 300, initialFirePlanetY + p.y)
	lib.DrawNormalImage(screen, p.imgGreenPlanet, 500, initialGreenPlanetY + p.y / 2)
	p.rocket.Draw(screen)
	
}

func (p *Planets) Init() {

	p.rocket.SetFireAllways()
	
}

func (p *Planets) Update() {

	if p.timeTraveling < maxTimeTravelingToPlanet {
		p.timeTraveling++
		p.y += 1
		p.rocket.MoveTo(370, 700 - p.y / 5)
	}
	
}

func (p *Planets) IsTraveling() bool {
	return p.timeTraveling < maxTimeTravelingToPlanet
}

