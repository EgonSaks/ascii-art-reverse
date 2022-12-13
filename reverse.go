package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

var readFlags = flag.String("reverse", "example.txt", "read file from flag")

// Reverse function calls and input checks
func reverse(args []string) {
	checkForAudit()
	fonts := "fonts/standard.txt"
	const usage = "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>"
	if !strings.Contains(*readFlags, "--reverse=") && len(args) == 1 {
		fmt.Println(args[0])
		os.Exit(0)
	}
	if len(args) > 0 {
		fmt.Println(usage)
		return
	}
	input, err := os.ReadFile("test/" + *readFlags)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
	}
	matrix := strings.Split(string(input), "\n")

	spaces := findSpace(matrix)
	userInput := splitUserInput(matrix, spaces)
	userInputMap := userInputMapping(userInput)
	asciiGraphic := getASCIIgraphicFont(fonts)

	output := mapUserInputWithASCIIgraphicFont(userInputMap, asciiGraphic)
	fmt.Println(output)
}

// Find a space (empty column) from user input
func findSpace(matrix []string) []int {
	var emptyColumns []int
	count := 0
	for column := 0; column < len(matrix[0]); column++ {
		for row := 0; row < len(matrix)-1; row++ {

			if matrix[row][column] == 32 {
				count++
			} else {
				count = 0
				break
			}
			if count == len(matrix)-1 {
				emptyColumns = append(emptyColumns, column)
				count = 0
			}

		}
	}

	// Check for extra spaces and convert them accordingly
	count = 5
	var indexToRem []int
	for i := range emptyColumns {
		if count == 0 {
			count = 5
			continue
		}
		if i > 0 {
			if emptyColumns[i] == (emptyColumns[i-1])+1 {
				indexToRem = append(indexToRem, i)
				count -= 1
			}
		}
	}
	for i := len(indexToRem) - 1; i >= 0; i-- {
		emptyColumns = removeIndex(emptyColumns, indexToRem[i])
	}
	return emptyColumns
}

// Remove index
func removeIndex(s []int, index int) []int {
	// check if the index is valid
	if index < 0 || index >= len(s) {
		return s
	}
	// remove the element at the specified index
	return append(s[:index], s[index+1:]...)
}

// Check if input is formatted correctly
func checkForAudit() {
	if strings.Contains(os.Args[1], "--") && !strings.Contains(os.Args[1], "=") {
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
		os.Exit(0)
	}
}

// Split user input based on space (empty column)
func splitUserInput(matrix []string, emptyColumns []int) string {
	var result string
	result = "\n"
	start := 0
	end := 0
	for _, column := range emptyColumns {
		if end < len(matrix[0]) {
			end = column

			for _, characters := range matrix {

				if len(characters) > 0 {
					columns := characters[start:end]
					result = result + columns + " "
				}
				result = result + "\n"

			}
			start = end + 1

		}
	}
	return result
}

// Map user input for search
func userInputMapping(result string) map[int][]string {
	strSlice := strings.Split(result, "\n")
	graphicInput := make(map[int][]string)
	j := 0
	for _, ch := range strSlice {
		if ch == "" {
			j++
		} else {
			graphicInput[j] = append(graphicInput[j], ch)
		}
	}
	return graphicInput
}

// Get ASCII graphic fonts
func getASCIIgraphicFont(fonts string) map[int][]string {
	readFile, err := os.ReadFile(fonts)
	if err != nil {
		fmt.Printf("Could not read the content in the file due to %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	ascii := make(map[int][]string)
	i := 31

	for _, ch := range slice {
		if ch == "" {
			i++
		} else {
			ascii[i] = append(ascii[i], ch)
		}
	}

	return ascii
}

// Match user input with ASCII graphic fonts for ascii and return an output
func mapUserInputWithASCIIgraphicFont(graphicInput, ascii map[int][]string) string {
	var keys []int
	for k := range graphicInput {
		keys = append(keys, k)
	}

	// Sort the keys in the slice
	sort.Ints(keys)

	// Initialize the output string
	var output string
	var sliceOfBytes []byte
	// Match user input with ASCII graphic fonts for ascii
	for _, value := range keys {
		graphicValue := graphicInput[value]
		for asciiKey, asciiValue := range ascii {
			if reflect.DeepEqual(asciiValue, graphicValue) {
				sliceOfBytes = append(sliceOfBytes, byte(asciiKey))
			}
		}
		output = string(sliceOfBytes)
	}
	return output
}
