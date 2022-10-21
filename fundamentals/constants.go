package main

import (
  "fmt"
)

const (
  a = iota
  b
  c
)

const (
  _ = iota + 2
  b2
  c2
)

const (
  _ = iota
  KB = 1 << (10 * iota)
  MB
  GB
  TB
  PB
)

func main()  {
  const myConst = 43
  fmt.Println(myConst)
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println("=================")
  fmt.Println(b2)
  fmt.Println(c2)
  fmt.Println("=================")
  fmt.Println(MB)
  fmt.Println(PB)
}
