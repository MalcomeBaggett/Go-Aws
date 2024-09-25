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
func (p *Person) changeName(newName string) {
	p.Name = newName
}
func main() {
	myPerson := Person{
		Name: "Malcome",
	}
	a := 7
	b := &a
	fmt.Println(b)
	myPerson.changeName("mr x")
	anotherNewPerson := NewPerson("M", 100)
	fmt.Println(myPerson.Name)
	fmt.Println(anotherNewPerson)
}
