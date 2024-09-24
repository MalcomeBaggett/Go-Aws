package main

import (
	"fmt"
	"slices"
)

func main() {
	// arrays
	animals := [2]string{}
	animals[0] = "dog"
	animals[1] = "cat"

	moreAnimals := []string{
		"walrus",
		"penguin",
	}
	for i := 0; i < len(animals); i++ {
		fmt.Printf("This is one of my animals %v\n", animals[i])
	}
	moreAnimals = append(moreAnimals, "wolf")
	for _, animal := range moreAnimals {
		fmt.Printf("This is a %v\n", animal)
	}
	moreAnimals = slices.Delete(moreAnimals, len(moreAnimals)-1, len(moreAnimals))
	fmt.Println(moreAnimals)

	// while loop
	for i := 0; i < 5; {
		fmt.Println(i)
		i++
	}
}
