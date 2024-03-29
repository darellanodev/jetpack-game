package main

import (
	"slices"
	"testing"
)

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
	t.Error("wanted an error but didn't get one")
	}
}
func TestLevel(t *testing.T) {

	t.Run("if level is empty then is an invalid level", func(t *testing.T) {
		level := ""
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level does not have 7 lines then is an invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has more than 7 lines then is an invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111" + CRLF + 
				 "111111"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in second row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "a00000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in third row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "-a0000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in fourth row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-a000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in last row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "11111a"
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("level is valid", func(t *testing.T) {
		level := "fire everywhere" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "212112"
		
		want := true
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("level is valid with a floor of lava with drops (3)", func(t *testing.T) {
		level := "fire everywhere" + CRLF +
				 "000000" + CRLF +
				 "-00000" + CRLF +
				 "0-0000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "113111"
		
		want := true
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("level files are incorrectly named", func(t *testing.T) {
		levelFiles := []string{"level5.txt","level6.txt"}
		
		err := verifyLevelNames(levelFiles)
		assertError(t, err)
	})

	t.Run("level files are incorrectly named, other example", func(t *testing.T) {
		levelFiles := []string{"level1.txt","level3.txt"}
		
		err := verifyLevelNames(levelFiles)
		assertError(t, err)
	})

	t.Run("level files are correctly named", func(t *testing.T) {
		levelFiles := []string{"level1.txt","level2.txt"}
		
		err := verifyLevelNames(levelFiles)
		
		if err != nil {
			t.Errorf("got %v want nil", err)
		}
	})

	t.Run("get level names", func(t *testing.T) {
		entries := []string{
			"2024/01/21 20:57:56 entry: -r--r--r-- 63 0001-01-01 00:00:00 level1.txt",
			"2024/01/21 20:57:56 entry: -r--r--r-- 63 0001-01-01 00:00:00 level2.txt",
		}
		
		want := []string{
			"level1.txt",
			"level2.txt",
		}
		got := getLevelFiles(entries)

		if !slices.Equal(want, got) {
			t.Errorf("got %v want %v", got, want)
		}

	})




}