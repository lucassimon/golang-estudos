package main

import (
	"fmt"
)

// channel unbuffered

func main() {

	msg := make(chan string)

	// rodando em paralelo e concorrente
	go func() {
		// enviando um texto para o channel
		msg <- "hello world"
	}()

	// libera o channel atribuindo o valor dele para a variavel result
	result := <-msg

	fmt.Println(result)
}
