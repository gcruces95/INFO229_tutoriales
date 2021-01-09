package main

import "fmt"

func main() {
	// Se genera un Slice cuyos valores seran 1,2,3,4 de tipo int.
	// No es necesario declarar el largo como en los arreglos.
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	// A diferencia de un arreglo, si no de asignan valores al slice, su valor es nulo.
	var slice2 []int
	fmt.Println(slice2)
	// Se usa len para saber el largo del slice
	fmt.Println(len(slice1))
	//Generaremos un slice a partir de un arreglo
	arreglo := [3]int{1, 2, 3}
	// Se genera el slice a partir de las posiciones 1 al 2 del arreglo
	// arreglo[:2] se define como slicing
	slice3 := arreglo[:2]
	fmt.Println(slice3)
}
