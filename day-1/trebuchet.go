package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	for scanner.Scan() {
		numberSlice := []string{}
		word := scanner.Text()
		for i := 0; i < len(word); i++ {
			if isInt(string(word[i])) {
				numberSlice = append(numberSlice, string(word[i]))
			}
		}
		totalSum += calculateStringConcat(numberSlice)
		fmt.Println("Word: ", word)
		fmt.Println("Number slice: ", numberSlice)
		fmt.Println("Sum: ", calculateStringConcat(numberSlice))
	}

	fmt.Println("Total sum: ", totalSum)
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func calculateStringConcat(numberSlice []string) int {
	var numberString string
	numberString = numberSlice[0] + numberSlice[len(numberSlice) - 1]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return 0
	}
	return number
}
