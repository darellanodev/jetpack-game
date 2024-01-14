package main

import (
	"embed"
	"io/fs"

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

func NewSoundFromFile(filesystem embed.FS, audioContext *audio.Context, path string) (*Sound, error) {

	s := &Sound{}

	f, err := filesystem.Open(path)

	if err != nil {
		return nil, err
	}

	s.f = f
	d, err := wav.DecodeWithoutResampling(f)

	if err != nil {
		return nil, err
	}

	player, err := audioContext.NewPlayer(d)

	if err != nil {
		return nil, err
	}

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

func LoadSounds() error {

	var err error

	sounds = make(map[string]*Sound)

	audioContext := audio.NewContext(sampleRate)

	for _, soundName := range []string{
		"fuel_pick",
		"start",
		"die",
		"traveling",
		"rocket_fuel_drop",
		"rocket_move",
	} {
		sounds[soundName], err = NewSoundFromFile(assets, audioContext, "assets/sounds/" + soundName + ".wav")

		if err != nil {
			return err
		}
	}

	return nil

}