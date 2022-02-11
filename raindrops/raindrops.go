package raindrops

import "strconv"

// converts a number into string with specific string, based on input number factors.
func Convert(number int) string {
	
	f := [3]int{3,5,7}
	s := [3]string{"Pling","Plang", "Plong"}
	rs:=""
	for i, j := range f {
		if ( number % j == 0){
			rs = rs + s[i]
		}
	}
	if len(rs) > 0 {
		return rs
	} else {
		return strconv.Itoa(number)
	}
}