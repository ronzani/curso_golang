package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo de structs")

	var u1 usuario
	fmt.Println(u1)

	u1.nome = "Robson"
	u1.idade = 28
	fmt.Println(u1)

	endereco1 := endereco{"Rua dos bobos", 0}

	u2 := usuario{"Pedro", 12, endereco1}
	fmt.Println(u2)

	u3 := usuario{nome: "Alice"}
	fmt.Println(u3)
}
