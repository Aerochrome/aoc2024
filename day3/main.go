package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := getInput()
	validInstructions := parseValidInstructions(input)

	total := 0
	apply := true

	for _, vi := range validInstructions {
		if len(vi) == 1 {
			if vi[0] == "do" {
				apply = true
			} else {
				apply = false
			}
			continue
		}

		if apply {
			total += calculateInstruction(vi)
		}
	}

	fmt.Printf("total: %d", total)
}

func calculateInstruction(instr []string) int {
	a, _ := strconv.Atoi(instr[1])
	b, _ := strconv.Atoi(instr[2])

	return a * b
}

func parseValidInstructions(str string) [][]string {
	result := make([][]string, 0)

	pattern := `(?:(mul)\(([0-9]+),([0-9]+)\))|(?:(don\'t|do)\(\))`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		if match[1] != "" {
			// is mul
			result = append(result, []string{match[1], match[2], match[3]})
			continue
		}

		result = append(result, []string{match[4]})
	}

	return result
}

func getInput() string {
	filePath := "day3/input.txt"

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
