package day1

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}


func Solve(){
	fmt.Println("solving function for day1");

	// Read the file
	inputFile, err := os.Open("day1/input_challenge1.txt")
	check(err)

	fileScanner := bufio.NewScanner(inputFile)

	regex := regexp.MustCompile("[0-9]")


	total := 0

	// Loop on each line of the file input
	for fileScanner.Scan(){

		var matches []string
		var stringNumber string
		var number int


		// FinAllStrign(text, maxMatchesReturned)
		// -1 for all matches
		matches = regex.FindAllString(fileScanner.Text(), -1)


		// Take the first and last string number for create a new string number
		stringNumber = matches[0] + matches[len(matches) - 1]
		number, err = strconv.Atoi(stringNumber)

		if(err != nil){
			number = 0
		}


		total += number
	}
	fmt.Printf("Total: %d \n", total)
}

