package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println("Array 1:", array1)
	array1[0] = 1
	array1[3] = 4
	fmt.Println("Array 1:", array1)

	var array2 [5]string
	array2[0] = "P1"
	fmt.Println("Array 2:", array2)

	array3 := [5]string{"P1"}
	fmt.Println("Array 3:", array3)
}
