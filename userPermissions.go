package main

import "fmt"

func permissions() {

	userRole := "admin"

	//fmt.Println("User Role:", userRole)

	switch userRole {

	case "admin":
		//fmt.Println("- Full System Access")
		fallthrough
	case "editor":
		//fmt.Println("- can publish content")
		fallthrough
	case "viewer":
		//fmt.Println("- can view content")
	default:
		//fmt.Println("No Access")
	}

	/*Calling functions*/
	looping()

}

func looping() {

	/*
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				continue
			}
			fmt.Println("Odd Number:", i)
		}
	*/

	i := 0
	for {
		fmt.Println("loop:", i)
		i++
		if i < 6 {
			break
		}
	}
}
