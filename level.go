package main

import (
	_ "image/png"
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

	switch f.number {
	case 1:
		f.title = "the kind planet"
		f.platformPlaces[0] = "00000"
		f.platformPlaces[1] = "10000"
		f.platformPlaces[2] = "01000"
		f.platformPlaces[3] = "00000"
		f.platformPlaces[4] = "00000"
		f.floorPlaces = 	  "111111"
	case 2:
		f.title = "fire everywhere"
		f.platformPlaces[0] = "00000"
		f.platformPlaces[1] = "00001"
		f.platformPlaces[2] = "01000"
		f.platformPlaces[3] = "00000"
		f.platformPlaces[4] = "00000"
		f.floorPlaces = 	  "212112"

	}
}

