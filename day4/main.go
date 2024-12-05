package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := getInput()
	occuring := 0

	for i, line := range input {
		for j, char := range line {
			if char != 'X' {
				continue
			}

			upPossible := false
			downPossible := false

			if i >= 3 {
				upPossible = true

				byteSlice := make([]byte, 0, 4)
				for _, el := range input[i-3 : i+1] {
					byteSlice = append(byteSlice, el[j])
				}

				if checkBackwards(string(byteSlice)) {
					occuring++
				}

			}

			if len(input)-1-i >= 3 {
				downPossible = true

				byteSlice := make([]byte, 0, 4)
				for _, el := range input[i : i+4] {
					byteSlice = append(byteSlice, el[j])
				}

				if checkForwards(string(byteSlice)) {
					occuring++
				}
			}

			if j >= 3 {
				if checkBackwards(line[j-3 : j+1]) {
					occuring++
				}

				// check i and then upper left
				if upPossible {
					byteSlice := make([]byte, 0, 4)

					hi := j
					vi := i

					for k := 0; k < 4; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi--
						vi--
					}

					if checkForwards(string(byteSlice)) {
						occuring++
					}
				}

				// check i and then lower left
				if downPossible {
					byteSlice := make([]byte, 0, 4)

					hi := j
					vi := i

					for k := 0; k < 4; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi--
						vi++
					}

					if checkForwards(string(byteSlice)) {
						occuring++
					}
				}
			}

			if len(line)-1-j >= 3 {
				if checkForwards(line[j : j+4]) {
					occuring++
				}

				// check i and then upper right
				if upPossible {
					byteSlice := make([]byte, 0, 4)

					hi := j
					vi := i

					for k := 0; k < 4; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi++
						vi--
					}

					if checkForwards(string(byteSlice)) {
						occuring++
					}
				}

				// check i and then lower right
				if downPossible {
					byteSlice := make([]byte, 0, 4)

					hi := j
					vi := i

					for k := 0; k < 4; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi++
						vi++
					}

					if checkForwards(string(byteSlice)) {
						occuring++
					}
				}
			}
		}
	}

	fmt.Printf("occurs: %d", occuring)
}

func checkBackwards(str string) bool {
	return str == "SAMX"
}

func checkForwards(str string) bool {
	return str == "XMAS"
}

func getInput() []string {
	filePath := "day4/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result
}
