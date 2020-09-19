package main

import (
	"fmt"
	"time"
)

// Exemplo de comunicacao entre channels

func main() {
	channel := make(chan int)

	// coroutine para input
	go func() {
		for i := 0; i < 10; i++ {

			// faz um input no channel
			// enquanto nao liberar o canal não vai atribuir um novo valor
			channel <- i
		}
	}()

	// coroutine para processar o channel
	go func() {
		for {
			// libera o channel usando o println
			fmt.Println(<-channel)
		}
	}()

	time.Sleep(time.Second)
}
