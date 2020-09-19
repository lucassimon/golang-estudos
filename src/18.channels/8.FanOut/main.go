package main

import (
	"fmt"
)

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
		close(channel)
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
		close(channel)
	}()

	return channel
}

func main() {
	channel := generate(20, 10)

	//  distribui
	d1 := div(channel)
	d2 := div(channel)

	fmt.Println(<-d1)
	fmt.Println(<-d2)
}
