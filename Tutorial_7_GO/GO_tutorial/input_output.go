package main

import (
	"fmt"
)

func main() {
	var edad int
	fmt.Printf("Ingresa tu edad: \n")
	//fmt.Scanln(&edad)
	//fmt.Println("Mi edad es: ", edad)
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es: %d\n", edad)
}
