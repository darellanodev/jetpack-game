package hud

import (
	_ "image/png"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type TypeWriter struct {
	x int
	y int
	text string
	displayText string
	currentCharacter int
	randMaxTimeToWrite int
	countTimeToWrite int
}

func NewTypeWriter(x int, y int, text string) *TypeWriter {
	return &TypeWriter {
		x: x,
		y: y,
		text: text,
		displayText: "",
		currentCharacter: 0,
		randMaxTimeToWrite: 5,
		countTimeToWrite: 0,
	}
}

func (t *TypeWriter) PutOneMoreCharacter(displayedText string, completeText string) string {

	displayedLength := len(displayedText) + 1
	result := completeText[:displayedLength]

	return result

}

func (t *TypeWriter) Update() {

	if t.countTimeToWrite < t.randMaxTimeToWrite {
		t.countTimeToWrite++
	} else {
		t.randMaxTimeToWrite = rand.Intn(10) + 5
		if len(t.displayText) < len(t.text) {
			t.displayText = t.PutOneMoreCharacter(t.displayText, t.text)
		}
	}

}

func (t *TypeWriter) Draw(screen *ebiten.Image) {

}