package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	faculdade string
	curso     string
}

func main() {
	fmt.Println("Herença")

	p1 := pessoa{"Robson", "Ronzani", 28, 180}
	fmt.Println(p1)

	e1 := estudante{p1, "UFT", "Computação"}
	fmt.Println(e1.nome, e1.curso, e1.faculdade)
}
