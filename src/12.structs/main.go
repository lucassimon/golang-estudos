package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 7.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

// Vehicle Conceito de interface é forcar que todos tenham a mesma assinatura
// padronizar também diversos adaptadores
type Vehicle interface {
	start() string
}

// Car Estrutura do carro e trabalhando com Tags
type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"-"`
	Color string `json:"color"`
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s, Year: %d, Color: %s", c.Name, c.Year, c.Color)
}

func (c Car) start() string {
	return "The Car:" + c.Name + " has been started"
}

// SuperCar é uma heranca para carro
type SuperCar struct {
	Car
	CanFly bool
}

// MotorCycle é um veiculo
type MotorCycle struct {
	Name string `json:"name"`
}

func (mc MotorCycle) start() string {
	return "The MotorCycle:" + mc.Name + " has been started"
}

// funcao generica que recebe um parametro do tipo veiculo
func startVehicle(v Vehicle) string {
	return v.start()
}

// interfaces vazias para criar tipos dinamicos
type Names []interface{}

//	Na hora de chamar esse metodo eu preciso passar o ponteiro
//
// que foi criado la funcao main
func (n *Names) load() {
	*n = Names{
		"Foo",
		"Bar",
		"Ping",
		"Pong",
		1,
		1.2222,
	}
}

func (n Names) print() {
	fmt.Println(n)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2.precoComDesconto())

	pedido := pedido{
		userID: 1,
		itens: []item{
			{produtoID: 1, qtde: 2, preco: 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())

	fmt.Println("--- type nota")
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

	fmt.Println("--- struct car")

	car1 := Car{"Corolla", 2020, "White"}
	car2 := Car{"BMW 320", 2020, "Black"}
	fmt.Println(car1.Name, car1.Year, car1.Color)
	fmt.Println(car1.info())
	fmt.Println(car2.Name, car2.Year, car2.Color)
	fmt.Println(car2.info())

	fmt.Println("--- struct supercar")
	fusca := Car{"Fusca", 1900, "White"}
	supercar := SuperCar{Car: fusca, CanFly: true}
	fmt.Println(supercar.info())

	fmt.Println("--- des/serializando para json")
	ferrari := Car{"Ferrari", 2020, "red"}
	result, _ := json.Marshal(ferrari)
	fmt.Println(result)
	fmt.Println(string(result))

	var gol Car
	data := []byte(`{"name": "Gol", "year": 2018, "color": "Yellow"}`)

	// apontar para o endereço de memória criado acima
	json.Unmarshal(data, &gol)
	fmt.Println(gol.Name, gol.Year, gol.Color)

	fmt.Println("--- trabalhando com interfaces")
	fmt.Println(startVehicle(gol))

	fazer := MotorCycle{"Fazer"}
	fmt.Println(startVehicle(fazer))

	fmt.Println("--- // interfaces vazias para criar tipos")
	var names Names
	names.load()
	names.print()
}
