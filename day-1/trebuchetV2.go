package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Elemento struct {
	Posicion int
	Valor    string
}

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
		var numberAndOcurrence = []Elemento{}
		word := scanner.Text()
		var numberSlice = []string{}
		var wordToDelete string = word
		fmt.Println("Attempting Word: ", word)
		for i := 0; i < len(possibleNumbers); i++ {
			if strings.Contains(wordToDelete, possibleNumbers[i]) {
				numberAndOcurrence = append(numberAndOcurrence, Elemento{strings.Index(word, possibleNumbers[i]), possibleNumbers[i]})
				fmt.Println("Found ", possibleNumbers[i], " at ", strings.Index(word, possibleNumbers[i]))
				wordToDelete = strings.Replace(wordToDelete, possibleNumbers[i], "", 1)
				fmt.Println("Word: ", wordToDelete)
				i = 0
			}
		}

		iter := 0

		for iter < len(word) {
			for i := 0; i < len(numberAndOcurrence); i++ {
				if iter == numberAndOcurrence[i].Posicion {
					var numberToInclude string
					if isInt(numberAndOcurrence[i].Valor) {
						numberToInclude = numberAndOcurrence[i].Valor
					} else {
						numberToInclude = fromStringToNumberString[numberAndOcurrence[i].Valor]
					}
					numberSlice = append(numberSlice, numberToInclude)
				}
			}
			iter++
		}

		fmt.Println("Word: ", word)
		fmt.Println("Number slice: ", numberSlice)
		fmt.Println("Sum: ", calculateStringConcat(numberSlice))

		totalSum += calculateStringConcat(numberSlice)

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
