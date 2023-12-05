package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var limitCubesPerColor = map[string]string{
	"red":   "12",
	"green": "13",
	"blue":  "14",
}

var sumValidGames = 0

func main() {
	fmt.Println("Starting game analyzer...")
	fmt.Println("Opening file...")
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading file...")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := adaptString(scanner.Text())

		wordSliced := strings.Split(word, ":")
		gameString := wordSliced[0]
		numberChar := getNumber(gameString)
		gameNumber, err := strconv.Atoi(numberChar)

		if err != nil {
			fmt.Println("Error converting string to int: ", err)
			return
		}

		if gameIsValid(wordSliced[1]) {
			sumValidGames += gameNumber
		}
	}

	fmt.Println("Sum of valid games: ", sumValidGames)
}

func getNumber(gameString string) string {
	numberChar := ""
	for i := 0; i < len(gameString); i++ {
		if isInt(string(gameString[i])) {
			numberChar += string(gameString[i])
		}
	}
	return numberChar
}

func adaptString(word string) string {
	word = strings.ReplaceAll(word, " ", "")
	return word
}

func gameIsValid(word string) bool {
	cubeSubset := strings.Split(word, ";")
	for subset := range cubeSubset {
		pick := strings.Split(cubeSubset[subset], ",")
		for element := range pick {
			colorNumber := getNumber(pick[element])
			color := strings.TrimPrefix(pick[element], colorNumber)
			if !isValidColor(color, colorNumber) {
				return false
			}
		}
	}
	return true
}

func isValidColor(color string, number string) bool {
	if limitCubesPerColor[color] == "" {
		return false
	}

	numberInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return false
	}

	limit, err := strconv.Atoi(limitCubesPerColor[color])
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
		return false
	}

	if numberInt > limit {
		return false
	}
	return true
}

func isInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
