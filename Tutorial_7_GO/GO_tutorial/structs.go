package main

import "fmt"

/*
	1.	Se usa para generar nuestras propias estructuras con tipo de datos.
	2.	type nombreEstructura struct{
		atributos de la estructura
	}
*/

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	//	Se crea la variable tipo User especificando parametros yatributos, no importa el orden
	persona1 := User{nombre: "Pepo", apellido: "Perez", edad: 30}
	fmt.Println(persona1)
	//	Si no se especifica los parametros, se deben ingresar los atributos en orden.
	persona2 := User{33, "Lala", "Lila"}
	fmt.Println(persona2)
	//	Imprime una struct sin declararla en variable.
	fmt.Println(User{nombre: "Pepa", apellido: "Lopez", edad: 25})
	//	Otra forma de declarar
	persona3 := new(User)
	persona3.nombre = "Lalo"
	persona3.apellido = "Lulu"
	persona3.edad = 45
	fmt.Println(*persona3)
}
