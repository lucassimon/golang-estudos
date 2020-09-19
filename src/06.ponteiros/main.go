package main

import "fmt"

func enderecoB() int {
	var a = 22
	fmt.Printf("Valor atribuido a variavel A %v  %T \n", a, a)
	fmt.Printf("Endereco de memoria da variavel A %v  %T \n", &a, a)

	return a
}

func referenciarVariaveis() int {
	var b = 42
	var c = &b
	fmt.Println("--- referenciarVariaveis")
	fmt.Printf("Valor atribuido a variavel B %v  %T \n", b, b)
	fmt.Printf("Endereco de memoria da variavel B %v  %T \n", &b, b)

	fmt.Println("---")
	fmt.Printf("Endereco de memoria da variavel C %v  %T \n", c, c)
	fmt.Printf("Endereco de memoria da variavel C %v  %T \n", *c, c)
	return b
}

func apontamentoUm() *int {
	var b = 62
	var c = &b

	fmt.Println("--- apontamentoUm")
	fmt.Printf("Valor atribuido a variavel B %v  %T \n", b, b)
	fmt.Printf("Endereco de memoria da variavel B %v  %T \n", &b, b)

	fmt.Println("---")
	fmt.Printf("Endereco de memoria da variavel C %v  %T \n", c, c)
	fmt.Printf("Endereco de memoria da variavel C %v  %T \n", *c, c)

	*c = 82

	fmt.Println("---")
	fmt.Printf("Altera o apontamento da variavel C para o novo valor %v  %T \n", c, c)
	fmt.Printf("E também a variavel B tem seu valor alterado %v  %T \n", b, b)

	return c
}

func outraFormaDeApontamento() *int {
	var b = 102
	var z *int = &b

	fmt.Println("--- outraFormaDeApontamento")
	fmt.Printf("Valor atribuido a variavel B %v  %T \n", b, b)
	fmt.Printf("Endereco de memoria da variavel B %v  %T \n", &b, b)

	fmt.Println("---")
	fmt.Printf("Endereco de memoria da variavel Z %v  %T \n", z, z)
	fmt.Printf("Endereco de memoria da variavel Z %v  %T \n", *z, z)

	return z
}

func naoAlteraValorNaMemoria(a int) int {
	a = a + 10
	return a
}

func alterandoValorNaMemoria(a *int) int {
	//  aqui estou alterando o valor contido dado o endereco de memoria para 100
	*a = 100
	return *a
}

func main() {
	enderecoB()
	referenciarVariaveis()
	apontamentoUm()
	outraFormaDeApontamento()

	fmt.Println("---")
	var foo = 10
	fmt.Printf("Novo valor gerado em apos chamar naoAlteraOValorNaMemoria %v \n", naoAlteraValorNaMemoria(foo))
	fmt.Printf("A variavel foo continua com 10 == %v \n", foo)

	fmt.Println("---")
	var bar = 20
	fmt.Printf("A variavel bar inicia com 20 == %v \n", bar)
	fmt.Printf("Novo valor gerado em apos chamar alterandoValorNaMemoria %v \n", alterandoValorNaMemoria(&bar))
	fmt.Printf("A variavel bar inicia com 100 == %v \n", bar)

}
