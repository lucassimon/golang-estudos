package main

import "fmt"

func mapsSimples() {
	// indice é uma string e o valor é um inteiro
	mapa := make(map[string]int)
	mapa["a"] = 10
	mapa["b"] = 20
	fmt.Println(mapa)
}

func deletandoValorDoMap() {
	// indice é uma string e o valor é um inteiro
	mapa := make(map[string]int)
	mapa["a"] = 10
	mapa["b"] = 20
	mapa["c"] = 30

	fmt.Println(mapa)
	delete(mapa, "c")
	fmt.Println(mapa["c"])
}

func verificandoValorDoMap() {
	// indice é uma string e o valor é um inteiro
	mapa := make(map[string]int)
	mapa["a"] = 10
	mapa["b"] = 20
	mapa["c"] = 30

	fmt.Println(mapa)
	delete(mapa, "c")
	fmt.Println(mapa["c"])

	_, exists := mapa["c"]
	fmt.Println(exists)

	value, exists := mapa["b"]
	fmt.Println(exists, value)

	if value, exists := mapa["a"]; exists {
		fmt.Println(exists, value)
	} else {
		fmt.Println("não existe")
	}
}

func outraDeclaracaoDoMap() {
	// indice é uma string e o valor é um inteiro
	var mapa = map[string]int{"foo": 10, "bar": 20}
	fmt.Println(mapa)
}

func percorrendoMaps() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])
	delete(aprovados, 95135745682)
	fmt.Println(aprovados[95135745682])
}

func mapsAninhados() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}

func main() {

	// maps podem ser comparados com os dicts do python

	mapsSimples()
	deletandoValorDoMap()
	verificandoValorDoMap()
	outraDeclaracaoDoMap()
	percorrendoMaps()
	mapsAninhados()
}
