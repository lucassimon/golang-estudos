package main

import (
	"fmt"
)

// Função inicialização do pacote
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")

	// funcoes em variaveis
	z := 0
	add := func() int {
		z += 2
		return z
	}
	fmt.Println(add())
	fmt.Println(add())
}
