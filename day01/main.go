package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/john-u/aoc-2023/util"
)

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += findCalibrationValue(line)
	}

	return sum
}

func findCalibrationValue(line string) int {
	digits := strings.Map(func(char rune) rune {
		if unicode.IsLetter(char) {
			return -1
		}

		return char
	}, line)

	var first, _ = utf8.DecodeRuneInString(digits)
	var last, _ = utf8.DecodeLastRuneInString(digits)

	value, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
	if err != nil {
		log.Fatal(err)
	}

	return value
}

func main() {
	var input = util.ReadInput()

	fmt.Println("Part 1:", partOne(input))
}
