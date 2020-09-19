package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateMsg(s string) chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - Value %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()

	return channel
}

func funil(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			// <- channel1 libera o canal1 para o channel
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()

	//  tenho 1 channel somente com os resultados de todos os channels
	return channel
}

func main() {

	x := funil(generateMsg("hello world"), generateMsg("foo bar"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}

	fmt.Println("Finished")
}
