# Tutorial #7 : GOlang

Este tutorial es inspirado del tutorial de [CodigoFacilito](https://codigofacilito.com/cursos/go). Introduce los conceptos básicos del lenguaje de programación [GO](https://golang.org/).

 1. Introducción 
 2. Primer código
 3. Declaración de variables
 4. Conversión de tipos
 5. Entrada y salida de datos
 6. Condiciones
 7. Ciclos
 8. Arrelos y Slices
 9. Make y Append
 10. Copy
 11. Punteros
 12. Struct
 13. Otros


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

Para descargar e instalar Go haz click [Aquí](https://golang.org/dl/)


## 2. Primer Código
### 2.1 Hello World

Todos los programas en go tienen como package un "main" el cual debe declararse al inicio, luego se declara una función homónima dentro del programa. Si este no tiene el package bien declarado no será ejecutable, mientras que sin la presencia del "main" se generará un error de compilación.

```
package main

func main() {

}
```

El paquete fmt nos permitirá imprimir y además implementa formatos E/S con funciones análogas a las del lenguaje C tales como printf y scanf. El formato de los 'verbos' son derivados también de C pero son más simples.
```
package main

//
import "fmt"	//Se importa libreria fmt

func main() {
	fmt.Println("Hola mundo")	// Se usa para mostrar un mensaje en terminal
}
```
**NOTA:** Para obtener más información sobre libreria "fmt", haga click [Aquí](https://golang.org/pkg/fmt/)                                                                                 

**NOTA:** Al igual que otros lenguajes go necesita puntos y coma(;), pero el mismo compilador de go se encarga de colocarlos, por lo que no es necesario escribirlos.

### 2.2 Compiladores

Al compilar nuestro código se genera un archivo compilado en lenguaje máquina, el cual hace que el código se pueda ejecutar en cualquier sistema operativo. Para compilar el código y generar el archivo ejecutable se debe escribir en consola lo siguiente:
```
go build nombreArchivo.go
```
Otra opción es compilar el código y ejecutarlo directamente en la consola de la siguiente forma: (Al ejecutar en consola no se creará el ejecutable como con go build)
```
go run nombreArchivo.go
```
## 3. Declaración de variables

Las variables declaradas en go deben establecer un tipo de dato, el cual no debe cambiar en el futuro (Una variable tipo string no puede almacenar un número, pues el tipo ya estará definido).
Las variables se declaran de la siguiente manera:
```
var nombre_variable tipo_dato
```
Donde 'var' deberá escribirse siempre que declaremos una variable, mientras que nombre_variable y tipo_dato dependerán del tipo de variable con el que se trabaje.
```
package main

import "fmt"

func main() {
	//var nombre_variable tipo_dato
	var x int
	//Se le asigna una valor int a una variable declarada como int
	x = 23
	//El operador := reconoce el tipo de variable y asigna el valor, es una variante de var nombre_variable tipo_dato
	y := 50 
	fmt.Println(x, y)
}
```
**NOTA:** Go permite también declarar el tipo de variable a más de una variable al mismo tiempo, a modo de ejemplo podemos probar a declarar 3 enteros diferentes:
```
var x,y,z int

```
Otra forma de declarar variables es omitir 'var' junto con el tipo de dato utilizando ':=', el cual se encargará de autodefinir el tipo de dato de la variable que estamos declarando.
```
x:= 23

```

## 4. Conversión de tipos

Utilizaremos la librería "strconv", la cual nos permitirá convertir el tipo de dato de las variables declaradas a la hora de utilizarlas.

```
package main

import (
	"fmt"
	"strconv"	//Se importa libreria strconv
)

func main() {
	edad1 := 22
	edad2 := "33"
	//strconv.Itoa convierte una variable int a string.
	edad_str := strconv.Itoa(edad1)
	//strconv.Atoi convierte una variable string a int.
	//strconv.Atoi devuelve multiples valores, pero como nos interesa solo 1 ignoramos el resto usando "_".
	edad_int, _ := strconv.Atoi(edad2)
	fmt.Println("Mi edad es: " + edad_str)
	fmt.Println("Tu edad es: ", edad_int)
}
```

**NOTA:** Para obtener más información sobre libreria "strconv", haga click [Aquí](https://golang.org/pkg/strconv/) 

## 5. Entrada y salida de datos

Dentro de la sintaxis de Go para leer y mostrar datos, está presente el uso de verbos. Aqui usaremos la libreria "fmt" para la entrada y salida de datos.

Tambien se puede usar la libreria "bufio".

**NOTA:** Para obtener más información sobre libreria "bufio", haga click [Aquí](https://golang.org/pkg/bufio/) 

### 5.1 Verbos

Tambien nombrados como verbos de formateo, se usan para el manejo de datos tanto de entrada como de salida.

```
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
```

**NOTA:** Para obtener más información sobre libreria "fmt" y los **verbos** haga click [Aquí](https://golang.org/pkg/fmt/)

### 5.2 Input/Output



```
package main

import "fmt"

func main() {
	var edad int
	fmt.Printf("Ingresa tu edad: \n")
	// Se utiliza Scanf para leer desde consola.
	// Scanf, recibe 3 parametros, el primero debe llevar un verdo segun corresponda al tipo de dato que se esta recibiendo,
	// el segundo parametro debe colocarse un &variable, el cual indica donde se almacenara el dato ingresado.
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es: %d\n", edad)
}
```

## 6. Condiciones

Dentro de las condicionales que tenemos en Go, se encuentran:
```
* if
* else if
* else
```
El if es un condicional el cual se ejecuta cuando la condicion que se evalua es true.
Si la condicion es falsa, pasa a else if.
Si la condicion de else if es falsa, se pasa a else, el cual se ejecuta por descarte como ultima opción.
```
package main

import "fmt"

func main() {
	//Como la condicion es verdadera, se ejecutan las instrucciones dentro del if
	if condicion {
		fmt.Printf("Se cumple la condicion, por ende entra al if")
  	} else condicion {
		fmt.Printf("No se cumple la condicion del if, pero si la de else if")
	} else {
		fmt.Printf("No se cumplio ninguna condicion anterior")
	}
}
```

Ejemplo uso de condiciones.

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

Una caracteristica especial de GO, es que no existe la estrucctura de ciclo While. Solamente tenemos el ciclo For, el cual se forma de 3 parametros los cuales no son oblogatorios usar, y se pueden omitir para realizar variaciones de la estructura for.

Ya que no existe el ciclo while, este se puede simular usando un for y un break, como se puede ver en el ejemplo siguiente.

```
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
```

## 8. Arrelos y Slices

Dentro de Go existen los arreglos, los cuales son estructuras discretas, en las cuales se debe definir el tipo de variable y el largo del arreglo, para reservar el espacio en memoria que usará.

Como una variante a los arreglo se crearon los Slices, los cuales son arreglo pero con una estructura dinámica, no es necesaria que se defina el largo y se incluye el concepto de capacidad. La si un slice supera su capacidad, se reserva un nuevo espacio de memoria y pero este sigue enlazado al slice original.

```
package main

import "fmt"

func main() {
	// Se genera un arreglo de largo 10 y tipo de variable int.
	// Por defecto, cada valor del arrelo es 0.
	var arreglo1 [10]int
	fmt.Println(arreglo1)
	// Se genera un arreglo de largo 3 y tipo de variable int.
	// Los valores del arreglo son 1,2,3.
	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)
	// Se genera un arreglo de largo 4 y tipo de variable int.
	// Los valores del arreglo son 1,2,3 y por defecto la ultima posicion tiene el valor 0.
	arreglo3 := [4]int{1, 2, 3}
	// Se le asigna el valor 4 a la posicion 4 de arrelo3.
	arreglo3[3] = 4
	fmt.Println(arreglo3)
	//len se usa para saber el largo del arreglo
	fmt.Println(len(arreglo1))
}
```


```
package main

import "fmt"

func main() {
	// Se genera un Slice cuyos valores seran 1,2,3,4 de tipo int.
	// No es necesario declarar el largo como en los arreglos.
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	// A diferencia de un arreglo, si no de asignan valores al slice, su valor es nulo.
	var slice2 []int
	fmt.Println(slice2)
	// Se usa len para saber el largo del slice
	fmt.Println(len(slice1))
	//Generaremos un slice a partir de un arreglo
	arreglo := [3]int{1, 2, 3}
	// Se genera el slice a partir de las posiciones 1 al 2 del arreglo
	// arreglo[:2] se define como slicing
	slice3 := arreglo[:2]
	fmt.Println(slice3)
}

```

## 9. Make y Append

```
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
```

## 10. Copy

Copy se usa para copiar el contenido de un arreglo en otro arreglo.

```
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

```

## 11. Punteros

```
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
```

## 12. Struct

```
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

```

## 13. Otros

En conclusion, en este tutorial hemos visto lo básico de Golang, este tutorial seguirá en desarrollo.
