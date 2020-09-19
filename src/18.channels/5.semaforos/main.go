package main

import (
	"fmt"
)

// flag para dizer que uma operacao foi concluida
// o semaforo é exatamente isso utilizando um channel

func main() {
	channel := make(chan int)
	//  semaforo
	ok := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		// terminou o processamento
		ok <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		// terminou o processamento
		ok <- true
	}()

	// corrotina para liberar
	go func() {
		// primeiro true do semaforo independente de quem processar primeiro
		<-ok
		// segundo true do semaforo independente de quem processar primeiro
		<-ok
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}
