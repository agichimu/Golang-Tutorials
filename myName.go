package main

import (
	"fmt"
	"strings"
)

func myName() {
	var name = "Alexander"
	/*
		agency is a variable name
		string is a variables datatype
	*/
	var agency = "Sky Movers"
	fmt.Println(name, "Welcome to", agency, "!!")
	/*
	   Calling Functions
	*/
	totalCars()
	prices()
	nameLength()
	stringEquality()
	tiny()
}

func totalCars() {
	var total = 50
	fmt.Println("Our Fleet Consists of", total, "cars")
}

func prices() {
	//Sample of Short Variable Declaration
	startingPrice := 29
	fmt.Println("Our starting Price is:", startingPrice)
}

func nameLength() {
	str := "Alexander"
	length := len(str)
	fmt.Println(length)
}

func stringEquality() {
	str1 := "Alexander"
	str2 := "alexander"
	strings.EqualFold(str1, str2)
	fmt.Println(strings.EqualFold(str1, str2))
}

func tiny() {
	var tinyUInt uint8 = 200
	fmt.Println(tinyUInt)
}
