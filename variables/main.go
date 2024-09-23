package main

import "fmt"

func main() {
	// var myName string = "Malcome"
	myName := "Malcome"
	myInt := 10
	myFloat := 10.0
	fmt.Printf("Hello my name is %s\nmy int is %d\nmy float is %.1f", myName, myInt, myFloat)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	fmt.Printf("my friends name is %v \nmy bool is %v\nmy num is %v", myFriendsName, myBool, myOtherInt)
}
