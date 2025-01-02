package main

import (
	"fmt"
	"reflect"
)

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

	array3 := [5]string{"P1", "P2", "P3"}
	fmt.Println("Array 3:", array3)
	array3[4] = "P5"
	fmt.Println("Array 3:", array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array 4:", array4)

	slice1 := []int{1, 2, 3}
	fmt.Println("Slice 1:", slice1)

	fmt.Println("Slice", reflect.TypeOf(slice1))
	fmt.Println("Array", reflect.TypeOf(array1))

	slice1 = append(slice1, 4)
	fmt.Println("Slice 1:", slice1)

	slice2 := array3[1:2]
	fmt.Println("Slice 2:", slice2)

	array3[1] = "P2-a"
	fmt.Println("Slice 2:", slice2)

}
