package main

import (
	"fmt"

	"github.com/lucassimon/golang-estudos/src/13.pacotes/car"
)

func main() {
	gol := car.Car{Name: "Gol Bolinha", Color: "Green"}
	fmt.Println(gol.Start())
}
