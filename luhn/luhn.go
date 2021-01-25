package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Valid : luhn algorithm
func Valid(input string) bool {

	// replacing any space from between the chars or leading & trailing spaces.
	input = strings.ReplaceAll(input, " ", "")

	// input with lengh 1 or less is not valid.
	if len(input) <= 1 {
		return false
	}

	regexp, _ := regexp.Compile("[0-9]+")
	matchInput := regexp.MatchString(input)

	if !matchInput {
		return false
	}

	//the int slice to hold the digits from input string.
	numSlice := make([]int, 0, len(input))

	for i, _ := range input {

		// converting to int from slice rune of char ( string is rune of chars )
		num, err := strconv.Atoi(input[i : i+1])
		if err != nil {
			fmt.Println("invalid input - not a number-", err, input[i:i+1])
			return false
		}
		// adding each int to slice of int
		numSlice = append(numSlice, num)

	}

	// int slice to hold the digits after doubling every 2nd digit from right
	secondNumHolderSlice := make([]int, 0, len(numSlice))

	numSliceSecondIndexholder := len(numSlice) - 2

	for numSliceSecondIndexholder >= 0 {
		secondValue := numSlice[numSliceSecondIndexholder]
		secondValue = secondValue * 2
		if secondValue > 9 {
			secondValue = secondValue - 9
		}
		secondNumHolderSlice = append(secondNumHolderSlice, secondValue)
		numSliceSecondIndexholder = numSliceSecondIndexholder - 2
	}

	// slice to hold original number at other place ( non-second digit )
	originalNumSlice := make([]int, 0, len(numSlice))

	nonSecondIndex := len(numSlice) - 1

	for nonSecondIndex >= 0 {
		originalNumSlice = append(originalNumSlice, numSlice[nonSecondIndex])
		nonSecondIndex = nonSecondIndex - 2
	}

	numSecondHolderTotal := 0
	for _, val := range secondNumHolderSlice {
		numSecondHolderTotal = numSecondHolderTotal + val
	}

	originalNumSliceTotal := 0
	for _, val := range originalNumSlice {
		originalNumSliceTotal = originalNumSliceTotal + val
	}
	finalTotal := numSecondHolderTotal + originalNumSliceTotal

	if (finalTotal % 10) == 0 {
		return true
	}
	return false
}
