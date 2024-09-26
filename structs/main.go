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
	*b = 9
	// The above code will derefernce to the original variable of a and turn it to 9 because b was pointing to the same memory adress as a
	fmt.Println(b)
	myPerson.changeName("mr x")
	anotherNewPerson := NewPerson("M", 100)
	fmt.Println(myPerson.Name)
	fmt.Println(anotherNewPerson)

	mySlicor := []int{
		1, 2, 3,
	}

	for _, value := range mySlicor {
		fmt.Println((value))
	}
}
