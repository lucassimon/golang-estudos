package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var m sync.Mutex

func runProcessWithTime(name string, count int) {
	for x := 0; x < count; x++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)

		// Travo o processo
		m.Lock()
		// incremento a variavel
		result++
		m.Unlock()
		//  libera o processo
		fmt.Println(name, "->", x, "partial result", result)
	}
}

func runComCoroutines() {
	//  iniciando go couroutines
	go runProcessWithTime("coroutine 1", 10)
	go runProcessWithTime("coroutine 2", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result", result)
}

// MUTEX
// setar uma chave processo preso para ninguém mexer

func main() {
	runComCoroutines()
}
