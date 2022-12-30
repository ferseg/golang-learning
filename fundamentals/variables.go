package main

import (
	"fmt"
)

func main() {
	var i = 42
	fmt.Println(i)
  var ii int = 50
  iii := 10
  var iiFloat = float32(ii)
  fmt.Println(ii)
  fmt.Println(iiFloat + 1.1)
  fmt.Println(iii)


  // Strings
  s := "This is a string"
  b := []byte(s)
  fmt.Println(s)
  fmt.Printf("%v", b)
}
