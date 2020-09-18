package main

import (
	"fmt"

	"github.com/lucassimon/golang-estudos/src/13.pacotes/car"
)

func main() {
	car := car.Car{"gol", "teste"}
	fmt.Println(car.start())
}
