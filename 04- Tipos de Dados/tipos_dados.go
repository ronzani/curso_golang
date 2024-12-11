package main

import (
	"errors"
	"fmt"
)

func main() {

	// numeros inteiros

	var n1 int8 = 100
	var n2 int16 = -10000
	var n3 int32 = 1000000000
	var n4 int64 = 100000000000000000
	var n5 int = 123456

	numero := 123

	fmt.Println(numero)
	fmt.Println(n1, n2, n3, n4, n5)

	var n6 uint = 200
	fmt.Println(n6)

	// alias
	//int32 = rune
	var n7 rune = 456
	// uint8 = byte
	var n8 byte = 78
	fmt.Println(n7, n8)

	// numeros reais
	var nr1 float32 = 123.45
	var nr2 float64 = 123467.890123
	nr3 := 456.789

	fmt.Println(nr1, nr2, nr3)

	// Strings
	var str1 string = "String-1"
	str2 := "string-2"
	fmt.Println(str1, str2)

	// boleano
	var b1 bool = true
	fmt.Println(b1)

	// valores "Zero"
	var numero1 int
	var texto string
	var b2 bool

	fmt.Println(numero1, texto, b2)

	// tipo erro
	var erro1 error
	var erro2 error = errors.New("Deu errado")

	fmt.Println(erro1, erro2)

}
