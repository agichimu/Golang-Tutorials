package main

import "fmt"

var carInventory = map[string]int{}

func addToInventory(carName string, amount int) {
	carInventory[carName] = amount
	fmt.Println("Added " + carName + " to the inventory")
}
func goFunctions() {
	addToInventory("Audi", 10)
	addToInventory("BMW", 100)
	addToInventory("Mercedes", 100)
	addToInventory("Bugatti", 10)
	addToInventory("RollsRoyce", 5)
	addToInventory("toyota", 10000000)
}
