package main

import "fmt"

func main() {
	// Se genera un arreglo de largo 10 y tipo de variable int.
	// Por defecto, cada valor del arrelo es 0.
	var arreglo1 [10]int
	fmt.Println(arreglo1)
	// Se genera un arreglo de largo 3 y tipo de variable int.
	// Los valores del arreglo son 1,2,3.
	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)
	// Se genera un arreglo de largo 4 y tipo de variable int.
	// Los valores del arreglo son 1,2,3 y por defecto la ultima posicion tiene el valor 0.
	arreglo3 := [4]int{1, 2, 3}
	// Se le asigna el valor 4 a la posicion 4 de arrelo3.
	arreglo3[3] = 4
	fmt.Println(arreglo3)
	//len se usa para saber el largo del arreglo
	fmt.Println(len(arreglo1))
}
