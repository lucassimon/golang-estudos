package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func gerarSenha(senha string) []byte {

	senhaEcriptada, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		fmt.Println(err)
	}

	return senhaEcriptada
}

func compararSenha(senhaEncryptada []byte, senha string) error {
	return bcrypt.CompareHashAndPassword(senhaEncryptada, []byte(senha))
}

func main() {
	senha := "Teste123"

	senhaEcriptada := gerarSenha(senha)

	if err := compararSenha(senhaEcriptada, senha); err != nil {
		fmt.Println("Senha errada")
	} else {
		fmt.Println("Senha correta")
	}

	if err := compararSenha(senhaEcriptada, "WrongPassword"); err != nil {
		fmt.Println("Senha errada")
	}

}
