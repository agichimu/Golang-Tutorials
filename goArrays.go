package main

import "fmt"

func arraysInGo() {

	var carModels [5]string

	carModels[0] = "BMW,"
	carModels[1] = "Mercedes,"
	carModels[2] = "Audi,"
	carModels[3] = "Toyota"

	//fmt.Println("carModels:", carModels)

	//carFleet()
}

// 2d arrays

func carFleet() {
	var carModels [3][2]string
	carModels[0] = [2]string{"4 BMW available ", "5 Mercedes available"}
	carModels[1] = [2]string{"lexus", "volvo"}
	carModels[2] = [2]string{"peageout", "tesla"}
	fmt.Println(carModels)
}
