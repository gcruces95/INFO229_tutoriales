# Tutorial #7 : GOlang

Este tutorial es inspirado del tutorial de [CodigoFacilito](https://codigofacilito.com/cursos/go). Introduce los conceptos básicos del lenguaje de programación [GO](https://golang.org/).

 1. Introducción 
 2. Hola mundo
 3. Variables
 4. Conversión de tipos
 5. Operadores
 6. Impresión y lectura
 7. Condicionales ciclos
 8. Funciones
 9. Arreglos
 10. Slides


## 1. Introducción
### 1.1 ¿Qué es GO?

Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico como Python y con el rendimiento de C o C++. Ha sido desarrollado por Google​ y sus diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson. Es un lenguaje de programación Open Source simple, eficiente y robusto, perteneciente a Google.

### 1.2 ¿Quienes usan GO?

Google
YouTube
Timehop
Uber
Twitch
Twitter
Mozilla
Muchas otras.

### 1.3 ¿Por qué usar GO?

Sintaxis clara y estandarizada
El compilador forza las buenas prácticas y es muy rápido.
soporta concurrencia y es fácil de manejar.
Fácil levantamiento de servidor web
Alto performance
fácil de aprender y subir a producción.

### 1.4 ¿Para que usar GO?

CommandLine Tools
Infraestructura de Backend
Herramientas internas de software
Servicios web
Herramientas de automatización

### 1.5 Instalación

Puede descargar haciendo click [Aquí](https://golang.org/dl/)


## 2. Primer Código
### 2.1 Hello World

Codigo base, debe ir si o si el package main
```
package main

func main() {

}
```

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.
```
package main

//
import "fmt"

func main() {
	fmt.Println("Hola mundo")
}
```
**NOTA:** Para obtener más información sobre libreria "fmt", haga click [Aquí](https://golang.org/pkg/fmt/)

### 2.2 Compiladores

Genera un archivo compilado en codio binario, el cual hace que el codigo se pueda ejecutar en cualquier sistema operativo.
```
go buil nombreArchivo.go
```
Compila el codigo, pero no genera archivo, y ejecuta directamente en consola.
```
go run nombreArchivo.go
```
## 3. Declaración de variables

var nombre_variable tipo_dato
```
package main

import "fmt"

func main() {
	//var nombre_variable tipo_dato
	var x int
	x = 23
	y := 50 //El operador := reconoce el tipo de variable y asigna el valor
	fmt.Println(x, y)
}
```

## 4. Conversión de tipos
```
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
```

## 5. Entrada y salida de datos
```
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
```
## 6. Condiciones
```
package main

import "fmt"

func main() {
	x := 10
	y := 5
	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d es menor que %d\n", x, y)
	} else {
		fmt.Printf("Los números son iguales\n")
  }
}
```
## 7. Ciclos
