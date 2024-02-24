package hud

import (
	_ "image/png"
	"math/rand"
)

type TypeWriter struct {
	displayText string
	randMaxTimeToWrite int
	countTimeToWrite int
}

func NewTypeWriter() *TypeWriter {
	return &TypeWriter {
		displayText: "",
		randMaxTimeToWrite: 5,
		countTimeToWrite: 0,
	}
}

func (t *TypeWriter) PutOneMoreCharacter(displayedText string, completeText string) string {

	displayedLength := len(displayedText) + 1
	result := completeText[:displayedLength]

	return result

}

func (t *TypeWriter) reinitRandomTimeToWriteACharacter() {
	t.randMaxTimeToWrite = rand.Intn(10) + 7
	t.countTimeToWrite = 0
}

func (t *TypeWriter) putNewCharacter(completeText string) {
	if len(t.displayText) < len(completeText) {
		t.displayText = t.PutOneMoreCharacter(t.displayText, completeText)
	}
}

func (t *TypeWriter) Write(completeText string) string {

	if t.countTimeToWrite < t.randMaxTimeToWrite {
		t.countTimeToWrite++
	} else {
		t.reinitRandomTimeToWriteACharacter()
		t.putNewCharacter(completeText)
	}

	return t.displayText

}
