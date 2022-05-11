package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func PrintResult(result []string) {
	if len(result) > 0 {
		fmt.Println("Result : ")
		for _, temp := range result {
			fmt.Println(temp)
		}
	} else {
		log.Fatalf("no question asked")
	}
}

func ReadInput() []string {
	reader := bufio.NewReader(os.Stdin)
	inputText := make([]string, 0)
	shouldStop := false
	fmt.Println("Please input lines of text (type 'quit' to finish input): ")
	for !shouldStop {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.ToLower(text) == "quit" || text == "" {
			shouldStop = true
		} else {
			inputText = append(inputText, text)
		}
	}
	return inputText
}

func CategorizeInput(inputText []string) ([]string, []string, []string) {
	hasNumber := regexp.MustCompile("[0-9]+")
	intergalacticToRomanInputArray := make([]string, 0)
	metalValueInputArray := make([]string, 0)
	questions := make([]string, 0)
	for _, text := range inputText {
		count := hasNumber.FindAllString(text, -1)
		if len(count) > 0 {
			metalValueInputArray = append(metalValueInputArray, text)
		} else if strings.Count(text, " ") == 2 {
			intergalacticToRomanInputArray = append(intergalacticToRomanInputArray, text)
		} else {
			questions = append(questions, text)
		}
	}
	return intergalacticToRomanInputArray, metalValueInputArray, questions
}
