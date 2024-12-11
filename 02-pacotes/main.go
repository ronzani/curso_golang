package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	// erro := checkmail.ValidateFormat("robson@email.com")
	erro := checkmail.ValidateFormat("robson.com")
	fmt.Println(erro)
}
