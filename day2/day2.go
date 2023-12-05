package day2

import (
	"fmt"

	utils "github.com/nicolasbalao/AdventOfCode2023/Utils"
)

func SolveP1() {
	fmt.Println("Solving day 2 part 1")

	inputFile := utils.FileScanner("day2/input.txt")

	for inputFile.Scan() {
		line := inputFile.Text()

		fmt.Println(line)
	}
}

func SolveP2() {

}
