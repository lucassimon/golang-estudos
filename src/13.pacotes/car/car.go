package car

import "fmt"

// Car Estrutura do carro e trabalhando com Tags
type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Info asdsd
func (c Car) Info() string {
	return fmt.Sprintf("Car: %s, Color: %s", c.Name, c.Color)
}

// Start asdasd
func (c Car) Start() string {
	return "The Car:" + c.Name + " has been started"
}
