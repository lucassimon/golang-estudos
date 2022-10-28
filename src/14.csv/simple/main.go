package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Contato estrutura de dados para a planilha
type Contato struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func readAll() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Printf("erro ao abrir o arquivo %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'

	reader.FieldsPerRecord = 3

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Printf("erro ao ler a informacao do arquivo %v", err)

	}

	fmt.Println(rawData)

}

func linhaALinha() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Printf("erro ao abrir o arquivo %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'

	reader.FieldsPerRecord = 3

	var rawData [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("erro ao ler a linha %v", err)
		}

		rawData = append(rawData, record)
	}

	fmt.Println(rawData)
}

func mapeandoParaStruct() {
	f, err := os.Open("data.csv")
	if err != nil {
		log.Printf("erro ao abrir o arquivo %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'

	reader.FieldsPerRecord = 3

	var contatos []Contato
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("erro ao ler a linha %v", err)
		}

		c := Contato{
			Nome:      record[0],
			Sobrenome: record[1],
		}

		if record[2] == "" {
			log.Printf("erro a idade nao pode ser vazia")
		}

		idade, err := strconv.Atoi(record[2])
		if err != nil {
			log.Printf("erro ao tratar a idade do contato %v", err)
			continue
		}
		c.Idade = idade

		contatos = append(contatos, c)
	}

	fmt.Println(contatos)
}

func main() {
	readAll()
	linhaALinha()
	mapeandoParaStruct()
}
