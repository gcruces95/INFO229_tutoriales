package main

import(
	"fmt"
)

func main(){
	var edad int
	fmt.Println("Ingresa tu edad: ")
	//fmt.Scanln(&edad)
	//fmt.Println("Mi edad es: ", edad)
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es: %d\n", edad)
}