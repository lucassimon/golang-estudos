package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runProcessWithTime(name string, count int) {
	for x := 0; x < count; x++ {
		fmt.Println(name, "->", x)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}

func runComCoroutines() {
	//  iniciando go couroutines
	go runProcessWithTime("coroutine 1", 10)
	go runProcessWithTime("coroutine 2", 10)

	var s string
	fmt.Scanln(&s)
}
