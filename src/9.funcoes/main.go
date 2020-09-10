package main

import (
	"fmt"
	"math"
	"runtime/debug"
	"strconv"
)

func soma(a int) int {
	return a + a
}

func subtracao(a int) int {
	return a - a
}

func divisao(a int) int {
	return a / a
}

func multiplicacao(a int) int {
	return a * a
}

func modulo(a int) int {
	return a % 2
}

func exponenciacaoAoQuadrado(a float64) float64 {
	return math.Pow(a, 2)
}

func operadorAnd(a int, b int) {
	fmt.Println("E =>", a&b)
}

func operadorOr(a int, b int) {
	fmt.Println("OR =>", a|b)
}

func operadorXor(a int, b int) {
	fmt.Println("XOR =>", a^b)
}

func maiorQue(a int, b int) {
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
}

func menorQue(a int, b int) {
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
}

func conversaoFloat() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
}

func conversaoInt() {
	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)
}

func conversaoASCII() {

	// Cuidado, string converte codigo ASC em string

	fmt.Println(string(67))
}

func conversaoAto() {
	fmt.Println("Teste " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")

	fmt.Println(num - 122)
}

func conversaoBool() {
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func retornarDoisValores(a string, b string) (string, string) {
	return a, b
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

func funcaoDentroDeFuncao() func() int {
	// decorator
	x := 10
	return func() int {
		return x * x
	}
}

func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func utilizandoDefer(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

// Função inicialização do pacote
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
	soma(10)
	subtracao(5)
	divisao(15)
	multiplicacao(2)
	modulo(64)
	exponenciacaoAoQuadrado(4)
	operadorAnd(3, 2)
	operadorOr(3, 2)
	operadorXor(3, 2)
	maiorQue(3, 2)
	menorQue(3, 2)
	conversaoFloat()
	conversaoInt()
	conversaoASCII()
	conversaoAto()
	conversaoBool()
	fmt.Println(namedReturn("Foo"))
	fmt.Println(retornarDoisValores("Hello", "world"))
	fmt.Println(diversosInteiros(1, 2, 5, 10))

	// funcoes em variaveis
	z := 0
	add := func() int {
		z += 2
		return z
	}
	fmt.Println(add())
	fmt.Println(add())

	mult := funcaoDentroDeFuncao()
	fmt.Println(mult())

	f1()

	// closure
	imprimeX := closure()
	imprimeX()

	// recursividade
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(utilizandoDefer(6000))

}
