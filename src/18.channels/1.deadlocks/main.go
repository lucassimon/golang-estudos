package main

import (
	"fmt"
)

func channelComDeadlock() {
	// cria o channel
	channel := make(chan int)
	// inputa um valor para o channel
	channel <- 10
	// libera o channel
	fmt.Println(<-channel)
}

func channelSemDeadlock() {
	// cria o channel
	channel := make(chan int)

	// TODO VALOR atribuido ao channel tem que ser executado em uma go coroutine

	go func() {
		// inputa um valor para o channel
		channel <- 10
	}()

	// libera o channel
	fmt.Println(<-channel)
}

func main() {
	channelComDeadlock()
	channelSemDeadlock()
}
