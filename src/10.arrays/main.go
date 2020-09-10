package main

import "fmt"

func inicializandoArrays() {
	fmt.Println("---- inicializandoArrays")
	// array predefinido com 10 posicoes
	x := [5]int{5, 10, 20, 17, 8}
	fmt.Println(x)
}

func arraySimples() {
	fmt.Println("---- arraySimples")

	// array predefinido com 10 posicoes
	var x [10]int
	fmt.Println(x)
	fmt.Println(len(x))

	// atribuicao em arrays
	x[0] = 10
	x[1] = 11
	x[2] = 12

	fmt.Println(x)

	// imprimir o indice 0
	fmt.Println(x[2])
}

func mediaDeNotas() {
	fmt.Println("---- mediaDeNotas")
	// homogênea (mesmo tipo) e estática (fixo)
	var notas [3]float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	// notas[3] = 10
	fmt.Println(notas)
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}

func percorrendoArrays() {
	fmt.Println("---- percorrendoArrays")
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}

func sliceParte1() {
	fmt.Println("---- sliceParte1")

	// cria na memoria um array com 5 posicioes iniciais
	slice := make([]int, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))

	// adiciona novas posições
	slice = append(slice, 10, 20, 13, 40)
	fmt.Println(slice)
}

func sliceParte2() {
	fmt.Println("---- sliceParte2")
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)
	fmt.Println(len(s1), len(s2))

	s1[0] = 7
	fmt.Println(s1, s2)
	fmt.Println(len(s1), len(s2))
}

func sliceParte2ComCapacidade() {
	fmt.Println("---- sliceParte2ComCapacidade")

	// cria na memoria um array com 2 posicioes iniciais e capacidade de 5
	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[1] = 10
	// vai gerar erro por causa que são 2 posições iniciais mesmo tendo capacidade de 5
	// slice[2] = 20

	// porem voce pode adicionar novas posicioes com append
	slice = append(slice, 10, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Se eu ultrapassar a minha capacidade ele cria um novo slice com novas capacidades
	slice = append(slice, 70)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func sliceParte3FazemApontamentos() {
	fmt.Println("---- sliceParte3ComCapacidade")

	// cria na memoria um array com 2 posicioes iniciais e capacidade de 5
	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 10, 2, 5, 40)

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("len: ", len(slice), "cap: ", cap(slice))
	}

	// lembrando que slices apontam para enderecos de memoria
	// ou seja nao criou uma cópia
	//
	sliceTest := slice
	slice[0] = 11
	fmt.Println(slice)
	fmt.Println(sliceTest)
}

func sliceParte4ApontamentoNoOriginal() {
	fmt.Println("---- sliceParte3ComCapacidade")

	// cria na memoria um array com 2 posicioes iniciais e capacidade de 5
	slice := make([]int, 2, 5)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 10, 2, 5, 40)

	//  apontei para o slice original
	sliceTest := slice

	// fiz um novo apontanmento para um outro slice com maior capacidade
	slice = append(slice, 1, 2, 3, 4, 5)
	slice[0] = 10

	// eles estão independentes agora
	fmt.Println(slice)
	fmt.Println(sliceTest)
}

func outraFormaDeCriarSlice() {
	// basta deixar a tipagem sem numero fixo
	sliceString := []string{"Hello", "World", "Much", "Better", "Forever"}
	fmt.Println(sliceString)
	fmt.Println(len(sliceString))

	// Andar dentro do slice
	fmt.Println(sliceString[0])

	// do comeco ATE a quantidade 2
	fmt.Println(sliceString[:2])

	// a partir da primeira elemento até a quantidade 2
	fmt.Println(sliceString[1:2])

	// a partir da segunda elemento até a 4
	fmt.Println(sliceString[2:4])

	// a partir da segunda elemento até o final
	fmt.Println(sliceString[2:])

}

func copiandoSlices() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	// array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}

func main() {
	inicializandoArrays()
	arraySimples()
	mediaDeNotas()
	percorrendoArrays()
	sliceParte1()
	sliceParte2()
	sliceParte2ComCapacidade()
	sliceParte3FazemApontamentos()
	sliceParte4ApontamentoNoOriginal()
	outraFormaDeCriarSlice()
	copiandoSlices()
}
