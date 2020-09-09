package main

import "fmt"

func condicionalIfElseIf() {

	var a int = 7

	if a > 10 {
		fmt.Println("é maior que 10")
	} else if a > 5 {
		fmt.Println("é maior que 5 e menor igual a 10")
	} else {
		fmt.Println("é menor que 5")
	}
}

func condicionalBooleano() {

	var a bool = true

	if a {
		fmt.Println("sempre vai cair aqui")

	} else {
		fmt.Println("nunca vai cair aqui")
	}
}

func condicionalBooleanoEInicializacao() {

	var a bool = true

	if foo := "foo"; a {
		fmt.Printf("sempre vai cair aqui e o valor de foo é atribuido em tempo real %v \n", foo)
	}
}

func main() {
	condicionalIfElseIf()
	condicionalBooleano()
	condicionalBooleanoEInicializacao()
}
