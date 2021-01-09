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

//	Se declara una funcion llamada nombre_completo() de tipo string.
//	usuario cumple el rol de identificador
func (usuario *User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

func main() {
	persona := new(User)
	persona.nombre = "Lalo"
	persona.apellido = "Lulu"
	persona.edad = 45
	fmt.Println(persona.nombre_completo())
}
