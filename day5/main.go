package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Type   string // bigger smaller
	Number int
}

type ParseResult struct {
	Rules       map[int][]Rule
	PageUpdates [][]int
}

func main() {
	input := getInput()
	validUpdates, invalidUpdates := findValidPageUpdates(input)
	middleNumSum := getMiddleNumSum(validUpdates)

	fixedInputs := validateInvalidUpdates(invalidUpdates, input.Rules)
	middleSumForFixed := getMiddleNumSum(fixedInputs)

	fmt.Printf("middle num sum: %d\n", middleNumSum)
	fmt.Printf("fixed middle num sum: %d", middleSumForFixed)
}

func validateInvalidUpdates(invalidUpdates [][]int, rules map[int][]Rule) [][]int {
	valid := make([][]int, 0, len(invalidUpdates))

	for _, update := range invalidUpdates {
		valid = append(valid, findValidOrder(update, rules))
	}

	return valid
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func newNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func findValidOrder(invalidUpdate []int, rules map[int][]Rule) []int {
	var rootNode *Node

	for _, num := range invalidUpdate {
		if rootNode == nil {
			rootNode = newNode(num)
			continue
		}

		rulesForNum, _ := rules[num]
		insertIntoNode(rootNode, num, rulesForNum)
	}

	return flattenNode(rootNode)
}

func insertIntoNode(node *Node, value int, rulesForValue []Rule) {
	for _, rule := range rulesForValue {
		if rule.Number != node.value {
			continue
		}

		if rule.Type == "bigger" {
			if node.right == nil {
				node.right = newNode(value)
				return
			}
			insertIntoNode(node.right, value, rulesForValue)
			return
		}
	}

	if node.left == nil {
		node.left = newNode(value)
		return
	}
	insertIntoNode(node.left, value, rulesForValue)
}

func flattenNode(node *Node) []int {
	if node == nil {
		return []int{}
	}

	flattened := make([]int, 0)

	flattened = append(flattened, flattenNode(node.left)...)
	flattened = append(flattened, node.value)
	flattened = append(flattened, flattenNode(node.right)...)

	return flattened
}

func getMiddleNumSum(validUpdates [][]int) int {
	sum := 0

	for _, u := range validUpdates {
		if len(u)%2 == 0 {
			log.Fatal("no middle number found")
		}

		sum += u[(len(u) / 2)]
	}

	return sum
}

func findValidPageUpdates(result ParseResult) ([][]int, [][]int) {
	validUpdates := make([][]int, 0, len(result.PageUpdates))
	invalidUpdates := make([][]int, 0, len(result.PageUpdates))

	for _, update := range result.PageUpdates {
		if isPageUpdateValid(update, result.Rules) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	return validUpdates, invalidUpdates
}

func isPageUpdateValid(update []int, rules map[int][]Rule) bool {
	beforeNums := make(map[int]struct{}, len(update))

	for _, num := range update {
		numRules, ok := rules[num]
		if !ok {
			beforeNums[num] = struct{}{}
			continue
		}

		for _, numRule := range numRules {
			if numRule.Type == "bigger" {
				if !slices.Contains(update, numRule.Number) {
					continue
				}

				if _, ok := beforeNums[numRule.Number]; !ok {
					return false
				}
			}
		}

		beforeNums[num] = struct{}{}
	}

	return true
}

func getInput() ParseResult {
	filePath := "day5/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make(map[int][]Rule)
	pageUpdates := make([][]int, 0)
	isRule := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			isRule = false
			continue
		}

		if isRule {
			split := strings.Split(line, "|")
			if len(split) != 2 {
				// Faulty line
				continue
			}

			left, _ := strconv.Atoi(split[0])
			right, _ := strconv.Atoi(split[1])

			if _, ok := rules[right]; !ok {
				rules[right] = make([]Rule, 0)
			}

			rules[right] = append(rules[right], Rule{
				Type:   "bigger",
				Number: left,
			})
			continue
		}

		split := strings.Split(line, ",")
		nums := make([]int, len(split))

		for i, numstr := range split {
			num, _ := strconv.Atoi(numstr)
			nums[i] = num
		}
		pageUpdates = append(pageUpdates, nums)
	}

	return ParseResult{
		Rules:       rules,
		PageUpdates: pageUpdates,
	}
}
