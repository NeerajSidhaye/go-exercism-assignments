package hamming

import (
	"fmt"
)

//Distance : calculating Haming Distance
func Distance(a, b string) (int, error) {

	
	if len(a) == len(b) {  

		var hamingDistanceCount int

		runeA := []rune(a)
		runeB := []rune(b)

		for i := range runeA {

			if runeA[i] != runeB[i] {
				hamingDistanceCount ++
			}
		}
		return hamingDistanceCount, nil
	} 
		 return 0, &argError{0, true}
	

}

func (e *argError) Error() string {
	return fmt.Sprintf("%-d %t", e.count, e.er)
}

type argError struct {
	count int 
	er bool
}
