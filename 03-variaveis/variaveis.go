package main

import "fmt"

func main() {
	var v1 string = "Variavel 1"
	v2 := "Variavel 2"

	fmt.Println(v1)
	fmt.Println(v2)

	var (
		v3 string = "Variavel 3"
		v4 string = "Variavel 4"
	)
	fmt.Println(v3, v4)

	v5, v6 := "Variavel 5", "Variavel 6"
	fmt.Println(v5, v6)

	const c1 string = "Constante 1"
	fmt.Println(c1)

	v5, v6 = v6, v5
	fmt.Println(v5, v6)
}
