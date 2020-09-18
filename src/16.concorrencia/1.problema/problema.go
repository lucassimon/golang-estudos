package main

import (
	"fmt"
)

func runProcess(name string, count int) {
	for x := 0; x < count; x++ {
		fmt.Println(name, "->", x)
	}
}

func runSequencial() {
	runProcess("sequencial 1", 10)
	runProcess("sequencial 2", 10)
}
