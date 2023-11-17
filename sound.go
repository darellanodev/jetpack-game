package main

import (
	"io/fs"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Sound struct {
	f      fs.File
	player *audio.Player
}


const (
	sampleRate   = 48000
)

var soundEnabled bool = initialSoundEnabled

var sounds map[string]*Sound

func NewSoundFromFile(audioContext *audio.Context, path string) (*Sound, error) {
	s := &Sound{}

	f, _ := os.Open(path)

	s.f = f

	// Decode wav-formatted data and retrieve decoded PCM stream.
	d, _ := wav.DecodeWithoutResampling(f)

	// Create an audio.Player that has one stream.
	player, _ := audioContext.NewPlayer(d)

	s.player = player

	return s, nil
}

func (s *Sound) Play() error {

	if !soundEnabled {
		return nil
	}
	if err := s.player.Rewind(); err != nil {
		return err
	}

	s.player.Play()

	return nil
}

func LoadSounds() {

	sounds = make(map[string]*Sound)
	// Initialize audio context.
	audioContext := audio.NewContext(sampleRate)

	for _, soundName := range []string{
		"fuel_pick",
		"start",
		"die",
	} {
		sounds[soundName], _ = NewSoundFromFile(audioContext, "assets/sounds/" + soundName + ".wav")
	}

}