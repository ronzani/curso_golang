package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(12, 21)
	fmt.Println(soma)

	var f1 = func(texto string) {
		fmt.Println(texto)
	}
	f1("Texto do f1")

	calc1, calc2 := calculos(15, 12)
	fmt.Println(calc1, calc2)

	calc3, _ := calculos(11, 8)
	fmt.Println(calc3)
}
