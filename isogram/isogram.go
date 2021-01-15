package isogram

import "strings"

//IsIsogram :  finds if an input string is an isogram ( non-repeating char of word ) word or not.
func IsIsogram(s string) bool {

	if len(s) == 0 {
		return true
	}

	i := 0
	matchCount := 0
	isoGramFlag := true

	s = strings.ToLower(s)

	s = strings.ReplaceAll(s, "-", "")

	s = strings.ReplaceAll(s, " ", "")

	for ; i < len(s); i++ {

		j := 0
		for ; j < len(s); j++ {

			if s[i] == s[j] {
				matchCount++
			}
		}

		if matchCount > 1 { // found char more than once
			isoGramFlag = false
			break
		}
		matchCount = 0
	}
	return isoGramFlag
}
