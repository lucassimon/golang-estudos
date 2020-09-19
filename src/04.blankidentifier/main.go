package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// retorna uma resposta e um error
	res, error := http.Get("https://google.com")

	if error != nil {
		fmt.Printf("%s", error)
	}

	// ignora o retorno da resposta de erro
	body, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s", body)
}
