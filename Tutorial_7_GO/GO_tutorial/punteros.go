package main

import "fmt"

func main() {
	/*
		1.	Es una direccion de memoria
		2.	En lugar del valor, tenemos la direccion en la que esta el valor
		3.	X, Y => as123d => 5		--	las variables X e Y apuntan a una direccion de memoria y el valor almacenado en esa direccion de memoria es 5.
		4.	X => as123d => 6		--	X cambia el valor almacenado en la direccion de memoria de 5 a 6.
		5.	El valor de las variables X e Y será 6.
		6.	*T	=>	*int, *string, *Struct
		7.	valor zero => nil
	*/

	var x, y *int
	entero := 5
	// & se usa para acceder a la direccion de memoria.
	x = &entero
	y = &entero
	// Se cambia el valor almacenado en la direccion de memoria de 5 a 6.
	*x = 6
	// Imprime las direcciones de memoria.
	fmt.Println(x)
	fmt.Println(y)
	// Imprime el valor almacenado en la direccion de memoria.
	fmt.Println(*x)
	fmt.Println(*y)
}
