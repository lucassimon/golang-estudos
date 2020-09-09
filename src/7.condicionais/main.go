package main

import (
	"fmt"
	"time"
)

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

func condicionalSwitchExemploUm(valor string) {

	switch valor {
	case "Foo":
		fmt.Println("O valor é Foo")
	case "Bar":
		fmt.Println("O valor é Bar")
	case "Bob":
		fmt.Println("O vaor é Bob")
	default:
		fmt.Println("Valor inválido")
	}
}

func condicionalSwitchExemploDois() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}

// Recebe qualquer tipo
func condicionalSwitchExemploTres(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não sei"
	}
}

func main() {
	condicionalIfElseIf()
	condicionalBooleano()
	condicionalBooleanoEInicializacao()
	condicionalSwitchExemploUm("Foo")
	condicionalSwitchExemploUm("Bar")
	condicionalSwitchExemploUm("Teste")
	condicionalSwitchExemploDois()
	fmt.Println(condicionalSwitchExemploTres(2.3))
	fmt.Println(condicionalSwitchExemploTres(1))
	fmt.Println(condicionalSwitchExemploTres("Opa"))
	fmt.Println(condicionalSwitchExemploTres(func() {}))
	fmt.Println(condicionalSwitchExemploTres(time.Now()))
}
