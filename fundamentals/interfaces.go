package main

import (
  "fmt"
)

func main() {
  var w Writer = ConsoleWriter{}
  w.Write([]byte("This is some message"))

  number := IntCounter(5)
  var counter Incrementer = &number

  fmt.Println(counter.Increment())
}

type Writer interface {
  Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
  return fmt.Println(string(data))
}


type Incrementer interface {
  Increment() int
}

type IntCounter int

func (counter* IntCounter) Increment() int {
  *counter++
  return int(*counter)
} 
