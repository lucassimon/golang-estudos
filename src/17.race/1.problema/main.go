package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func runProcessWithTime(name string, count int) {
	for x := 0; x < count; x++ {
		// p1
		z := result
		z++

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		fmt.Println(name, "->", x, "partial result", result)

		result = z
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

// pid 1 z = 1
// result = 1

// em outro momento o pid2 faz a mesma coisas
// pid 2 z = 1
// result = 1

// atribuindo sempre um valor de acordo com ultimo processo executado
// assim nunca vai ter o total do result

// os dois processos rodando simultaneamente estão interferindo no resultado final.

func main() {
	runComCoroutines()
}
