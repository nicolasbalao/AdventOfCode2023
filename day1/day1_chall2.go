package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func SolveChall2() {

	fmt.Println("Solving day 1 chall 2")

	inputFile, err := os.Open("day1/input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(inputFile)

	// Golang regexp doesn't work with lookahead assertion
	regex := regexp.MustCompile("(?=(one|two|three|four|five|six|seven|eight|nine|\\d))")
	numbersMap := map[string]string{
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

	total := 0
	line := 0

	for fileScanner.Scan() {
		line += 1

		var digits []string

		matches := regex.FindAllString(fileScanner.Text(), -1)

		fmt.Println(fileScanner.Text())
		fmt.Println(matches)
		fmt.Println(line)
		digits = append(digits, matches[0], matches[len(matches)-1])

		for i, digit := range digits {
			_, err := strconv.Atoi(digit)

			if err != nil {
				digits[i] = numbersMap[digit]
			}
		}

		numberString := digits[0] + digits[1]

		number, _ := strconv.Atoi(numberString)

		total += number
	}

	fmt.Printf("Total: %v\n", total)

}
