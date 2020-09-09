package main

import "fmt"

var y int = 20

// runs with go run *.go

func main() {
	fmt.Printf("Primeiro escopo %v \n", y)
	printY()
	fmt.Printf("Escopo do pacote %v \n", z)
	printZ()
}

func printY() {
	fmt.Printf("%v \n", y)
}
