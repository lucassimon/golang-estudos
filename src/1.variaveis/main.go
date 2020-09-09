package main

import "fmt"

func atribuicaoShorthand() {
	a := 10
	fmt.Printf("%v %T \n", a, a)
}

func atribuicaoInteiro() int {
	var b int
	b = 22
	fmt.Printf("%v  %T \n", b, b)
	return b
}

func atribuicaoString() (string, string) {
	var c, d string = "hello", "world"
	fmt.Printf("%s %s %T %T \n", c, d, c, d)
	return c, d
}

func atribuicaoPontoFlutuante() float32 {
	var b float32
	b = -43.186798345560456
	fmt.Printf("%v %T \n", b, b)
	return b
}

func atribuicaoBooleano() bool {
	var b bool
	b = false
	fmt.Printf("%v %T\n", b, b)
	return b
}

func atribuicaoStringMultiplasLinhas() string {
	var c string = `Hello

	world

	from

	multiple lines
	`
	fmt.Printf("%s %T \n", c, c)
	return c
}

func main() {
	atribuicaoShorthand()
	atribuicaoInteiro()
	atribuicaoString()
	atribuicaoPontoFlutuante()
	atribuicaoBooleano()
	atribuicaoStringMultiplasLinhas()
}
