package main

import (
	"fmt"
)

type Animal struct {
	Name                     string
	FirstAppearenceTimestamp int32
	Characteristics          []string
}

type Dog struct {
	Animal
	CanBark bool
}

func main() {
	milo := Dog{}
	milo.Name = "Milo"
	milo.FirstAppearenceTimestamp = 1234
	milo.Characteristics = []string{
		"sweet", "quiet", "funny",
	}
	milo.CanBark = true

	secondDog := milo
	secondDog.Name = "Filomena"

	fmt.Println(milo)
	fmt.Println(secondDog)
}
