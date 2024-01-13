package main

import (
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

	characterZeroCount := strings.Count(line, "0")
	characterOneCount := strings.Count(line, "1")

	totalValidCharacters := characterZeroCount + characterOneCount
	
	return totalValidCharacters == 6
}

func isLevelValid(level string) bool {

	if (len(level) == 0) {
		return false
	}

	if strings.Count(level, CRLF) != 7 {
		return false
	}

	lines := strings.Split(level, CRLF)

	result := isLineValid(lines[1]) && isLineValid(lines[2]) && isLineValid(lines[3])

	return result
	
}

