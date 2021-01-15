package isogram

import "strings"

//IsIsogram :  finds if an input string is an isogram ( non-repeating char of word ) word or not.
func IsIsogram(s string) bool {

	if len(s) == 0 {
		return true
	}

	i := 0
	matchCount := 0
	var isoGramFlag bool = true
	
	s = strings.ToLower(s)

	if strings.ContainsAny(s, "-") {
		s = strings.ReplaceAll(s, "-", "")
	}
	if strings.ContainsAny(s, " ") {
		s = strings.ReplaceAll(s, " ", "")
	}
	// by now, the input string will have all in lowercase and with NO hyphen or space characters in it.
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
