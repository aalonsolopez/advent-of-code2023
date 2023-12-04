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
		var numberToOcurrence = map[string]int{}
		numberSlice := []string{}
		word := scanner.Text()
		for i := 0; i < len(possibleNumbers); i++ {
			if strings.Contains(word, possibleNumbers[i]) {
				numberToOcurrence[possibleNumbers[i]] = strings.Index(word, possibleNumbers[i])
			}
		}
		for len(numberToOcurrence) > 0 {
			var minIndex int
			var minNumber string
			for key, value := range numberToOcurrence {
				if value < minIndex {
					minIndex = value
					minNumber = key
				}
			}
			numberSlice = append(numberSlice, fromStringToNumberString[minNumber])
			delete(numberToOcurrence, minNumber)
		}

		fmt.Println("Word: ", word)
		fmt.Println("Number slice: ", numberSlice)
	}

	fmt.Println("Total sum: ", totalSum)
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
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
