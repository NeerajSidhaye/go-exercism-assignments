package luhn

import (
	"strconv"
	"strings"
)

//Valid :
func Valid(input string) bool {

	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 2 {
		return false
	}

	doubleIndex := 0
	if len(input) %2 !=0 {
		doubleIndex = 1
	}

	sum :=0
	for i, val := range input {

		digit, err := strconv.Atoi(string(val))
		if err!=nil {
			return false
		}

		if doubleIndex == i {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9 
			}
			doubleIndex = doubleIndex + 2
		}

		sum = sum + digit
	}

	return (sum%10)==0

}

