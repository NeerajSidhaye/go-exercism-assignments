// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	byteRemark :=  []byte(remark)
	fmt.Println(byteRemark)
	
	yelling, _ := regexp.Compile(`^[^a-z]*[A-Z][^a-z]*$`)
	question,_ := regexp.Compile(`\?\s*$`)
	silence, _:= regexp.Compile(`^\s*$`)

	if yelling.Match(byteRemark) {
		if question.Match(byteRemark){
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if question.Match(byteRemark){
		return "Sure."
	}

	if silence.Match(byteRemark) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
