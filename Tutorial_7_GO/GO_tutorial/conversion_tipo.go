package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad1 := 22
	edad2 := "33"
	edad_str := strconv.Itoa(edad1)
	edad_int, _ := strconv.Atoi(edad2)
	fmt.Println("Mi edad es: " + edad_str)
	fmt.Println("Tu edad es: ", edad_int)
}
