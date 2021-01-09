package main

import "fmt"

func main() {
	// estructura -> make(tipo_dato, Longitud, Capacidad)
	// Si se omite la capacidad, este toma el valor de la longitud
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	// Se usa len para obtener el largo del slice
	fmt.Println(len(slice))
	// Se usa cap para obtener la capacidad del slice
	fmt.Println(cap(slice))
	// Se usa apend para agregar un elemento al slice
	slice = append(slice, 2)
	fmt.Println(slice)
}
