package main

import "fmt"

func repeticaoSimples() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}
}

func repeticaoSimulandoWhile() {
	var x int = 10
	fmt.Println("repeticaoSimulandoWhile", x)
	for x < 10 {
		fmt.Printf("%v \n", x)
		x++
	}
}

func repeticaoInfinita() {
	var x int = 10
	fmt.Println("repeticaoInfinita", x)
	for true {
		fmt.Printf("%v \n", x)
		x++
		if x == 15 {
			break
		}
	}

	// ou
	x = 10
	fmt.Println("repeticaoInfinita", x)
	for {
		fmt.Printf("%v \n", x)
		x++
		if x == 25 {
			break
		}
	}
}

func main() {
	repeticaoSimples()
	repeticaoSimulandoWhile()
	repeticaoInfinita()
}
