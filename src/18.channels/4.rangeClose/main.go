package main

import (
	"fmt"
	"sync"
)

// utilizando conceitos como waitGroup dentro dos channels
// imprimir valores através do range

func main() {
	channel := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	// roda em paralelo as 2 funcoes abaixo
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}

		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		// espera até o loop das duas funcoes terminarem
		waitGroup.Wait()
		// fecha o canal
		close(channel)
	}()

	// esvazia o channel por causa do for...range
	for number := range channel {
		fmt.Println(number)
	}

}
