package main

import (
	"testing"

	"github.com/darellanodev/jetpack-game/hud"
)

func TestTypeWriter(t *testing.T) {

	t.Run("if displayed text is 'phr' and complete text is 'phrase' then result must be 'phra'", func(t *testing.T) {
		
		want := "phra"
		
		completePhrase := "phrase"
		
		typeWriter := hud.NewTypeWriter(10, 10, completePhrase)

		displayedPhrase := "phr"
		got := typeWriter.PutOneMoreCharacter(displayedPhrase, completePhrase)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if displayed text is '' and complete text is 'phrase' then result must be 'p'", func(t *testing.T) {
		
		want := "p"
		
		completePhrase := "phrase"
		
		typeWriter := hud.NewTypeWriter(10, 10, completePhrase)

		displayedPhrase := ""
		got := typeWriter.PutOneMoreCharacter(displayedPhrase, completePhrase)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if displayed text is 'phrase' and complete text is 'phrase' then result must be 'phrase'", func(t *testing.T) {
		
		want := "phra"

		completePhrase := "phrase"
		
		typeWriter := hud.NewTypeWriter(10, 10, completePhrase)

		displayedPhrase := "phr"
		got := typeWriter.PutOneMoreCharacter(displayedPhrase, completePhrase)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})


}