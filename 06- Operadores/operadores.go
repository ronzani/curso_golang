package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 5 / 2
	multiplicacao := 5 * 2
	resto := 5 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// ATRIBUIÇÃO
	var v1 string = "Variavel 1"
	v2 := "Variavel 2"

	fmt.Println(v1, v2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	// OPERADORES UNÁRIOS
	fmt.Println("--------------")
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	fmt.Println(numero)
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	fmt.Println(numero)
	numero /= 10
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
