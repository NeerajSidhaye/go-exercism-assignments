package scrabble

import "strings"

// return score of a given word based on the predefined char values.
func Score(word string) int {
  
  val := make(map[string]int)
  val["AEIOULNRST"] = 1
  val["DG"] = 2
  val["BCMP"]= 3
  val["FHVWY"] = 4
  val["k"] = 5
  val["JX"] = 8
  val["QZ"] = 10

  input := []rune(word)
  res:=0
  for _,i := range input{   
      ct := 1
      r:=0
    for k,v := range val {
      
      if(strings.ContainsAny(strings.ToLower(k), strings.ToLower(string(i)))){
        r = ct * v 
      }
    }
     res = r + res
  }

  return res

}
