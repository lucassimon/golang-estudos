package main

import "fmt"

var y int = 20

// runs with go run *.go

func main() {
	fmt.Printf("Primeiro escopo %v \n", y)
	PrintY()
	fmt.Printf("Escopo do pacote %v \n", z)
	PrintZ()
}

// PrintY some Print
func PrintY() {
	fmt.Printf("%v \n", y)
}
