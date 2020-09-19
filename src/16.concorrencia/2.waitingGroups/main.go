package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func init() {

	fmt.Println("Numero de CPUS ", runtime.NumCPU())

	// Voce pode setar o numero de CPUS para executar

	//  vai rodar de forma concorrente em 1 cpu
	// runtime.GOMAXPROCS(1)

	//  vai rodar concorrente e em paralelo de acordo com o numero de cpus setado
	//  a partir da versão 1.5 ele pega o total de num de cpus
	runtime.GOMAXPROCS(3)
}

func runProcessWithTime(name string, count int) {
	for x := 0; x < count; x++ {
		fmt.Println(name, "->", x)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}

	// eu preciso deizer para o meu programa que o processo da coroutine acabou
	waitGroup.Done()
}

func runComCoroutines() {

	// estou adicionando 2 processos para serem executados
	waitGroup.Add(2)
	//  iniciando go couroutines
	go runProcessWithTime("coroutine 1", 10)
	go runProcessWithTime("coroutine 2", 10)

	// espera os 02 processos acabarem
	waitGroup.Wait()

}

func main() {
	runComCoroutines()
}
