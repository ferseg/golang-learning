package main

import (
  "fmt"
)

func main()  {
  // Array
  var arr = [...]int{1, 2, 3}
  fmt.Println(arr)


  arrCopy := arr
  arrCopy[1] = 5

  fmt.Println(arr)
  fmt.Println(arrCopy)

  arrRef := &arr
  arrRef[2] = 4

  fmt.Println(arr)
  fmt.Println(arrRef)

  // Slice
  thisIsASlice := []int{1,2,3}
  sliceRef := thisIsASlice
  sliceRef[0] = 10

  fmt.Println(thisIsASlice)
  fmt.Println(sliceRef)

}
