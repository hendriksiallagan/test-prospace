package build

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/test-prospace/service"
)

func BuildIntergalacticTointergalacticToRomanMap(intergalacticToRomanInputArray []string) (map[string]string, error) {
	intergalacticToRomanMap := make(map[string]string)
	for _, roman := range intergalacticToRomanInputArray {
		temp := strings.Split(roman, " ")
		intergalacticToRomanMap[temp[0]] = temp[2]
	}
	for _, tmp := range intergalacticToRomanMap {
		found := false
		for romanChar := range service.RomanNumerals {
			if tmp == romanChar {
				found = true
			}
		}
		if !found {
			return intergalacticToRomanMap, fmt.Errorf("%s is not a roman character", tmp)
		}
	}
	return intergalacticToRomanMap, nil
}

func BuildMetalValueMap(metalValueInputArray []string, intergalacticToRomanMap map[string]string) (map[string]int, error) {
	metalValueMap := make(map[string]int)
	for _, intergalactic := range metalValueInputArray {
		romanValues := ""
		temp := strings.Split(intergalactic, " ")
		for i, tmp := range temp {
			if x, ok := intergalacticToRomanMap[tmp]; ok {
				romanValues += x
			}
			if strings.ToLower(tmp) == "is" {
				rom, err := service.RomanToInt(romanValues)
				if err != nil {
					return metalValueMap, err
				}
				if len(temp) == i {
					return metalValueMap, fmt.Errorf("invalid format")
				}
				val, err := strconv.Atoi(temp[i+1])
				if err != nil {
					return metalValueMap, err
				}
				metalValueMap[temp[i-1]] = val / rom
				continue
			}
		}
	}
	return metalValueMap, nil
}

func BuildResult(questions []string, intergalacticToRomanMap map[string]string, metalValueMap map[string]int) ([]string, error) {
	result := make([]string, 0)
	for _, question := range questions {
		var tempResult string
		foundRomanValues := false
		temp := strings.Split(question, " ")
		romanValues := ""
		for i, tmp := range temp {
			if x, ok := intergalacticToRomanMap[tmp]; ok {
				tempResult += tmp + " "
				romanValues += x
				foundRomanValues = true
			}
			if x, ok := metalValueMap[tmp]; ok {
				rom, err := service.RomanToInt(romanValues)
				if err != nil {
					return result, err
				}
				tempResult += fmt.Sprintf("%s is %d Credits", tmp, x*rom)
				break
			}
			if i == len(temp)-1 {
				if foundRomanValues {
					rom, err := service.RomanToInt(romanValues)
					if err != nil {
						return result, err
					}
					tempResult += fmt.Sprintf("is %d", rom)
				} else {
					tempResult = "I have no idea what you are talking about"
				}
			}
		}
		result = append(result, tempResult)
	}
	return result, nil
}
