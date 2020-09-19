package main

import "fmt"

// Foo Constante visivel para outros pacotes só por causa da maiuscula
const Foo string = "foo"

// bar Constante usada apenas nesse modulo
const bar string = "bar"

func criaMultiplasConstantes() (int, int, int) {
	const (
		r int = 10
		g int = 10
		b int = 10
	)

	return r, g, b
}

func criaConstanteB() int {
	const b = 22
	fmt.Printf("Constante B %v  %T \n", b, b)
	return b
}

func main() {
	var r, g, b = criaMultiplasConstantes()
	criaConstanteB()

	fmt.Printf("Constante RGB(%v, %v, %v) \n", r, g, b)
	fmt.Printf("Constante global %v  %T \n", Foo, Foo)
	fmt.Printf("Constante do pacote %v  %T \n", bar, bar)
}
