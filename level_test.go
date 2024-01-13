package main

import "testing"
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
				 "100000" + CRLF +
				 "010000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has more than 7 lines then is an invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "100000" + CRLF +
				 "010000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111" + CRLF + 
				 "111111" + CRLF
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in second row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "a00000" + CRLF +
				 "100000" + CRLF +
				 "010000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111" + CRLF
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in third row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "1a0000" + CRLF +
				 "010000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111" + CRLF
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("if level has an invalid character in fourth row then invalid level", func(t *testing.T) {
		level := "the kind planet" + CRLF +
				 "000000" + CRLF +
				 "100000" + CRLF +
				 "01a000" + CRLF +
				 "000000" + CRLF +
				 "000000" + CRLF + 
				 "111111" + CRLF
		
		want := false
		got := isLevelValid(level)
		
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})


}