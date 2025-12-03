package main

import (
	"fmt"
)

type standart struct {
	Name string
	Age  int
}

func initStandart(name string, age int) standart {
	return standart{
		Name: name,
		Age:  age,
	}
}

func (s standart) Speak() {
	fmt.Println("My name is", s.Name, "and I am", s.Age, "years old.")
}

func main() {
	a := initStandart("Alice", 30)
	a.Speak()
}
