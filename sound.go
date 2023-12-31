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
	f, _ := filesystem.Open(path)
	s.f = f
	d, _ := wav.DecodeWithoutResampling(f)
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

	audioContext := audio.NewContext(sampleRate)

	for _, soundName := range []string{
		"fuel_pick",
		"start",
		"die",
		"traveling",
		"rocket_fuel_drop",
		"rocket_move",
	} {
		sounds[soundName], _ = NewSoundFromFile(assets, audioContext, "assets/sounds/" + soundName + ".wav")
	}

}