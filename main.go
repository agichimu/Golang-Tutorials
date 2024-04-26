package main

import (
	"fmt"
	"strings"
	_ "strings"
)

func main() {
	String()

	/*Variables are used to Store data*/

	var nameOne string = "Alexander"
	var agency string = "fast track!"
	var totalDevelopers int = 29
	carBrand := "Porsche."
	startingPrice := 50.55 /*Float*/

	/*
		Agency = name of the variable
		String is a datatype
	*/

	/*
	   x := 4
	   y := 5
	   fmt.Println(x * y)
	*/

	fmt.Println("Hello", nameOne)
	fmt.Println(agency)
	fmt.Println("My dream car is a", carBrand)
	fmt.Println("My team consists of", totalDevelopers, "Developers.")
	fmt.Printf("Our Prices range From %v", startingPrice)

	/*
		%v holds place of a variable in a formatted string
	*/
}

func String() {
	/*
		Go strings are immutable which means when created they can not be modified
	*/
	stg := "Hello, World!"
	length := len(stg)
	fmt.Println("Length =", length)
	fmt.Println(strings.ToUpper(stg)) //converts to uppercase
	fmt.Println(strings.ToLower(stg)) //converts to lowercase
}
