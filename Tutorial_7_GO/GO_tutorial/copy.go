package main

import "fmt"

func main() {
	arreglo := []int{1, 2, 3, 4}
	fmt.Println(arreglo)
	// Se genera un arreglo vacío de longitud 4
	copia_arreglo := make([]int, 4)
	fmt.Println(copia_arreglo)
	// Usaremos la funcion copy, que sirve para copiar un arreglo
	// estructura -> copy(destino, fuente)
	// Copiamos el contenido del arreglo "arreglo", en el arreglo "copia_arreglo"
	copy(copia_arreglo, arreglo)
	fmt.Println(copia_arreglo)
}
