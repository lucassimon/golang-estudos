package main

import (
	"fmt"
	"os"

	databases "github.com/lucassimon/golang-estudos/src/20.factory/databases"
)

func main() {
	var t int
	fmt.Println("Informe um tipo de conexao")
	_, err := fmt.Scanf(&t)
	if err != nil {
		fmt.Println("Erro ao ler a opção", err)
		os.Exit(1)
	}

	conn := databases.Factory(t)

	if conn == nil {
		fmt.Println("Não foi possivel conectar no database escolhido")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Println("Erro ao ler a opção", err)
		os.Exit(1)
	}

	err = conn.Close()
	if err != nil {
		fmt.Println("Erro ao fechar a conexão")
	}

}
