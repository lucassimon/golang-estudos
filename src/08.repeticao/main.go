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

func diversosInteiros(x ...int) int {
	// variadic parametros
	var total int = 0
	for index, valor := range x {
		total += valor
		fmt.Printf("indice %v, total = %v \n", index, total)
	}
	return total
}

func repeticaoHierarquica() {

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("horas: ", horas)

		for minutos := 0; minutos < 60; minutos++ {
			fmt.Println("minutos: ", minutos)
		}
		fmt.Println()
	}
}

func main() {
	repeticaoSimples()
	repeticaoSimulandoWhile()
	repeticaoInfinita()
	diversosInteiros(1, 2, 3, 4, 5, 6)
}
