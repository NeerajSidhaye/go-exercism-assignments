package hamming

import (
	"errors"
)

//Distance : calculating Haming Distance
func Distance(a, b string) (int, error) {

	if len(a) == len(b) {

		var hamingDistanceCount int

		i := 0
		for ; i < len(a); i++ {

			if a[i] != b[i] {
				hamingDistanceCount++
			}
		}
		return hamingDistanceCount, nil
	}
	return 0, errors.New("input strings not equal")
}