package service

import (
	"fmt"
)

var RomanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) (int, error) {
	sum := 0
	greatest := 0
	letterBefore := ""
	repeatingLetter := 0
	for i := len(s) - 1; i >= 0; i-- {
		letter := s[i]
		if string(letter) == letterBefore {
			repeatingLetter++
		} else {
			repeatingLetter = 0
			letterBefore = string(letter)
		}
		if repeatingLetter > 2 {
			return 0, fmt.Errorf("not a roman letter")
		}
		num := RomanNumerals[string(letter)]
		if num >= greatest {
			greatest = num
			sum += num
			continue
		}
		sum -= num
	}
	if sum <= 0 {
		return 0, fmt.Errorf("not a roman letter")
	}
	return sum, nil
}
