package main

import (
	"fmt"
)

// passar channels por funcoes
// pipeline uma coisa consumindo a outra

// funcao variatica
// funcao que vai gerar numeros
func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()

	return channel
}

func div(input chan int) chan int {
	channel := make(chan int)
	go func() {
		//  libera o channel
		for number := range input {
			channel <- number / 2
		}
	}()

	return channel
}

func main() {
	numbers := generate(2, 4, 6, 8)
	result := div(numbers)

	// libera o resultado de 2
	fmt.Println(<-result)
	// libera o resultado de 4
	fmt.Println(<-result)
	// libera o resultado de 6
	fmt.Println(<-result)
	// libera o resultado de 8
	fmt.Println(<-result)
}
