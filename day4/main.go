package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := getInput()
	task1 := task1(input)
	task2 := task2(input)

	fmt.Printf("task1: %v\n", task1)
	fmt.Printf("task2: %v\n", task2)
}

func task1(input []string) int {
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

	return occuring
}

type storeKey struct {
	x int
	y int
}

func task2(input []string) int {
	aStore := make(map[storeKey]int)

	for i, line := range input {
		for j, char := range line {
			if char != 'M' {
				continue
			}

			upPossible := false
			downPossible := false

			if i >= 2 {
				upPossible = true
			}

			if len(input)-1-i >= 2 {
				downPossible = true
			}

			if j >= 2 {
				// check i and then upper left
				if upPossible {
					byteSlice := make([]byte, 0, 3)

					hi := j
					vi := i

					for k := 0; k < 3; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi--
						vi--
					}

					if checkForwardsMas(string(byteSlice)) {
						increaseStore(aStore, j-1, i-1)
					}
				}

				// check i and then lower left
				if downPossible {
					byteSlice := make([]byte, 0, 3)

					hi := j
					vi := i

					for k := 0; k < 3; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi--
						vi++
					}

					if checkForwardsMas(string(byteSlice)) {
						increaseStore(aStore, j-1, i+1)
					}
				}
			}

			if len(line)-1-j >= 2 {
				// check i and then upper right
				if upPossible {
					byteSlice := make([]byte, 0, 3)

					hi := j
					vi := i

					for k := 0; k < 3; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi++
						vi--
					}

					if checkForwardsMas(string(byteSlice)) {
						increaseStore(aStore, j+1, i-1)
					}
				}

				// check i and then lower right
				if downPossible {
					byteSlice := make([]byte, 0, 3)

					hi := j
					vi := i

					for k := 0; k < 3; k++ {
						byteSlice = append(byteSlice, input[vi][hi])

						hi++
						vi++
					}

					if checkForwardsMas(string(byteSlice)) {
						increaseStore(aStore, j+1, i+1)
					}
				}
			}
		}
	}

	occurring := 0

	for _, el := range aStore {
		if el > 1 {
			occurring++
		}
	}

	return occurring
}

func increaseStore(store map[storeKey]int, x, y int) {
	key := storeKey{x, y}

	if _, ok := store[key]; !ok {
		store[key] = 1
		return
	}

	store[key] = store[key] + 1
}

func checkBackwards(str string) bool {
	return str == "SAMX"
}

func checkForwards(str string) bool {
	return str == "XMAS"
}

func checkForwardsMas(str string) bool {
	return str == "MAS"
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
