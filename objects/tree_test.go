package objects

import (
	"testing"
)
func TestTreePosition(t *testing.T) {

	t.Run("if the rand value is less than 40 then set the tree's posX to -1000 to place it off the screen", func(t *testing.T) {
		want := -1000 // position off the screen

		randValue := 20
		floorPosX := 100
		isRocketFloor := false
		got := CalculateTreePositionX(floorPosX, randValue, isRocketFloor)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if the rand value is greater than 40 then the tree has the same posX as the floor", func(t *testing.T) {
		want := 100 // position same as the floor

		randValue := 70
		floorPosX := 100
		isRocketFloor := false
		got := CalculateTreePositionX(floorPosX, randValue, isRocketFloor)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if is over a floor where the rocket lands then set the tree posX outside the screen", func(t *testing.T) {
		want := -1000 // position outside the screen

		randValue := 70
		floorPosX := 100
		isRocketFloor := true
		got := CalculateTreePositionX(floorPosX, randValue, isRocketFloor)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
