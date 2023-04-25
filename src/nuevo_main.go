package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

// cuando se va a invocar una funcion concatenada se le da la primera en minuscula
//y la segunda en mayuscula ejemplo tripleArgument
//es un abuena practica que cuando se va a nombrar mas de una variable y tiene el mismo nobre solo se marca la ultima
// de las que son es decir que a y b son la misma variable ejemplo a, int b, int c, string se coloca a, b int, c string

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

// si necesitamos que la funcion retorne un valor, o retorene 2 o mas valores le decimos que nos de un
//doblereturn la forma de agregarlo es con un (a int) agregarlo dentro de un parentesis y decimos que van dos valores enteros
// y agregarlo en el mismo orden en el que esta arriba y en el segundo que retorne el * 2
//la forma de recibirlo es en el mismo orden
// si solo necesitamos un valor se le coloca _ porque go pide que se use lo que se delcara
// y descarta ese valor y solo toma el value 1 ejemplo
//value1, value2 := doubleReturn(2)
//fmt.Println("value1 y value2:", value1, value2)

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("value:", value)

	value1, _ := doubleReturn(2)
	fmt.Println("value1", value1)
}

//reto llevar las funciones aquellos algoritmos que  utilce para calcular las areas

//Docuemntacion y como implementar las de terceros y aprender a leer codigo go https://pkg.go.dev/  https://pkg.go.dev/std

// cilo o tareas repetitivas For condicional
