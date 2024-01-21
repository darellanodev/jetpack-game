package main

import (
	"errors"
	"fmt"
	_ "image/png"
	"slices"
	"strings"

	"github.com/darellanodev/jetpack-game/lib"
)

var (
	levels []string
	totalGameLevels int
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

func (f *Level) Load()  {

	level := levels[f.number-1]

	levelLines := strings.Split(level, CRLF)

	f.title 			= levelLines[0]
	f.platformPlaces[0] = levelLines[1]
	f.platformPlaces[1] = levelLines[2]
	f.platformPlaces[2] = levelLines[3]
	f.platformPlaces[3] = levelLines[4]
	f.platformPlaces[4] = levelLines[5]
	f.floorPlaces 		= levelLines[6]

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

func getLevelFiles() ([]string, error) {

	var levelFiles []string
	entries, err := assets.ReadDir("assets/levels")

	if err != nil {
		return nil, err
	}
	for _, entry := range entries{
		if strings.Contains(entry.Name(), ".txt") {
			levelFiles = append(levelFiles, entry.Name())
		}
	}

	return levelFiles, nil
}

func verifyLevelNames(levelFiles []string) error {

		totalLevels := len(levelFiles)
		
		var searchFile string
		var isCorrect bool
		for i := 1; i <= totalLevels; i++ {
			searchFile = fmt.Sprintf("level%d.txt", i)
			isCorrect = slices.Contains(levelFiles, searchFile)

			if !isCorrect {
				return errors.New("The file " + searchFile + " does not exists")
			}
		}

		return nil
}

func LoadLevels() error {

	levelFiles, err := getLevelFiles()
	totalGameLevels = len(levelFiles)

	if totalGameLevels == 0 {
		return errors.New("there are no levels")
	}
	
	if err != nil {
		return err
	}

	err = verifyLevelNames(levelFiles)

	if err != nil {
		return err
	}

	for i := 1; i <= totalGameLevels; i++ {
		levelPath := fmt.Sprintf("assets/levels/level%d.txt", i)
		levelBytes, err := lib.LoadStaticResource(assets, levelPath)

		level := string(levelBytes)

		if err != nil {
			return err
		}

		if !isLevelValid(level) {
			msg := fmt.Sprintf("level %d (%v) has an invalid format", i, levelPath)
			return errors.New(msg)
		}

		levels = append(levels, level)

	}

	return nil
}

