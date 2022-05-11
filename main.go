package main

import (
	"log"

	"github.com/test-prospace/build"
	"github.com/test-prospace/usecase"
)

func main() {

	inputText := usecase.ReadInput()

	intergalacticToRomanInputArray, metalValueInputArray, questions := usecase.CategorizeInput(inputText)

	intergalacticToRomanMap, err := build.BuildIntergalacticTointergalacticToRomanMap(intergalacticToRomanInputArray)
	if err != nil {
		log.Fatal(err)
	}

	metalValueMap, err := build.BuildMetalValueMap(metalValueInputArray, intergalacticToRomanMap)
	if err != nil {
		log.Fatal(err)
	}

	result, err := build.BuildResult(questions, intergalacticToRomanMap, metalValueMap)
	if err != nil {
		log.Fatal(err)
	}

	usecase.PrintResult(result)
}
