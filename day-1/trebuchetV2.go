package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting trebuchet calibrations...")
	fmt.Println("Opening file...")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading file...")

	scanner := bufio.NewScanner(file)

	totalSum := 0

	var possibleNumbers = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	var fromStringToNumberString = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for scanner.Scan() {
		word := scanner.Text()
		first := findFirst(word, possibleNumbers)
		last := findLast(word, possibleNumbers)
		if len(first) > 1 {
			first = fromStringToNumberString[first]
		}
		if len(last) > 1 {
			last = fromStringToNumberString[last]
		}

		totalSum += calculateStringConcat([]string{first, last})

		fmt.Println("Word: ", word)
		fmt.Println("First: ", first)
		fmt.Println("Last: ", last)
		fmt.Println("Sum: ", calculateStringConcat([]string{first, last}))

	}

	fmt.Println("Total sum: ", totalSum)
}

func findFirst(str string, substrs []string) string {
	firstIndex := len(str)
	value := ""
	for _, substr := range substrs {
		if pos := strings.Index(str, substr); pos > -1 && pos < firstIndex {
			firstIndex = pos
			value = substr
		}
	}
	return value
}

func findLast(str string, substrs []string) string {
	lastIndex := -1
	value := ""
	for _, substr := range substrs {
		if pos := strings.LastIndex(str, substr); pos > lastIndex {
			lastIndex = pos
			value = substr
		}
	}
	return value
}

func calculateStringConcat(numberSlice []string) int {
	var numberString string
	numberString = numberSlice[0] + numberSlice[len(numberSlice)-1]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return 0
	}
	return number
}
