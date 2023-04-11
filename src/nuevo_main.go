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
//doblereturn la forma de agregarlo es en el mismo orden en el que esta arriba

func returnValue(a int) int {
	return a * 2
}

func doubleReturn()

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("value:", value)
}
