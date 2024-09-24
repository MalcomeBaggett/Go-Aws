package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	myPerson := Person{
		Name: "Malcome",
	}
	anotherNewPerson := NewPerson("M", 100)
	fmt.Println(myPerson)
	fmt.Println(anotherNewPerson)
}
