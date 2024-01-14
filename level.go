package main

import (
	"errors"
	"fmt"
	_ "image/png"
	"strings"
)

type Level struct {
	number				int
	title				string
	platformPlaces		[5]string
	floorPlaces			string
}

func NewLevel() *Level {
	
	return &Level{
		number: startingLevel,
		title:  "",
	}
}


func (f *Level) Next() {
	f.number++
	f.Load()
}

func (f *Level) Load() {

	levelTxt := string(loadStaticResource(assets, fmt.Sprintf("assets/levels/level%d.txt", f.number)))
	lines := strings.Split(levelTxt, CRLF)

	f.title 			= lines[0]
	f.platformPlaces[0] = lines[1]
	f.platformPlaces[1] = lines[2]
	f.platformPlaces[2] = lines[3]
	f.platformPlaces[3] = lines[4]
	f.platformPlaces[4] = lines[5]
	f.floorPlaces 		= lines[6]

}

func isLineValid(line string) bool {

	emptyLevelCharacterCount := strings.Count(line, emptyLevelCharacter)
	platformLevelCharacterCount := strings.Count(line, platformLevelCharacter)
	normalFloorLevelCharacterCount := strings.Count(line, normalFloorLevelCharacter)
	lavaFloorLevelCharacterCount := strings.Count(line, lavaFloorLevelCharacter)

	totalValidCharacters := emptyLevelCharacterCount + normalFloorLevelCharacterCount + lavaFloorLevelCharacterCount + platformLevelCharacterCount
	
	return totalValidCharacters == 6
}

func isLevelValid(level string) bool {

	lines := strings.Split(level, CRLF)
	if len(lines) == 0 {
		return false
	}

	if len(lines) != totalRowsTxt {
		return false
	}

	result := true
	for i := firstLevelRowTxt; i <= totalLevelRowsTxt; i++ {
		result = result && isLineValid(lines[i])
	}

	return result
}

func CheckLevels() error {

	for i := 1; i <= 2; i++ {
		levelPath := fmt.Sprintf("assets/levels/level%d.txt", i)
		level := string(loadStaticResource(assets, levelPath))

		if !isLevelValid(level) {
			msg := fmt.Sprintf("level %d (%v) has an invalid format", i, levelPath)
			return errors.New(msg)
		}

	}

	return nil
}

