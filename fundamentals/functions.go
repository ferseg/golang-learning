package main

import (
	"fmt"
)

func main() {

	result := sum(5, 6, 87, 2, 1)
	fmt.Println(result)

	divideResult, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(divideResult)
  
  g := greeter {
    greeting: "Hey",
    name: "YOU",
  }

  g.greet()
}

func sum(values ...int) (result int) {
	for _, value := range values {
		result += value
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// Methods
type greeter struct {
	greeting string
	name     string
}

func (g* greeter) greet() {
  fmt.Println(g.greeting, g.name)
}
