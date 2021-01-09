package main

import "fmt"

func main() {
	//Ciclo for estandar
	//la variable i se inicia en 0, se ejecutara mientras i sea menor a 10, cada pasada por el ciclo el valor de i aumenta en 1.
	for i := 0; i < 10; i++ {
		//imprime el valor de i.
		fmt.Println(i)
	}
	fmt.Println("")
	//
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	fmt.Println("")
	//Ciclo for que simula a while.
	g := 0
	for {
		fmt.Println(g)
		g++
		if g >= 10 {
			//break se usa pasa cortar la ejecucion del ciclo.
			break
		}
	}
}
