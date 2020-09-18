package carros

import "fmt"

// Car Estrutura do carro e trabalhando com Tags
type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s, Color: %s", c.Name, c.Color)
}

func (c Car) start() string {
	return "The Car:" + c.Name + " has been started"
}
