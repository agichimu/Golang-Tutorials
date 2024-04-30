package main

import "fmt"

func goconstants() {

	// const series string = "Fast & Furious"

	const (
		_ = iota
		Economy
		Compact
		Standard
		Fullsize
		Luxury
	)
	fmt.Println("Economy", Economy)
	fmt.Println("Compact", Compact)
	fmt.Println("Standard", Standard)
	fmt.Println("Fulsize", Fullsize)
	fmt.Println("Luxury", Luxury)
}
