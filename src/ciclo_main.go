package main

import "fmt"

func main() {

	// for condicional se le escribe una condicion para que vaya iterando for que es el kiwor para poder invocarlo
	//un i igual a 0 que indica que indice va a empezar en 0 va a ejecutarce hasta que i sea menos a 10 y en cada ciclo va a sumar un digito
	// tendriamos una secuanecia de 0 hasta el 9 ejemplos queremos que imprima el 10 le decimos que se <= 10
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//for wile hasta que una condicion se cumpla en este caso vamos a hacer un counter = a 0 y esta condicion se va
	//cumpliendo el se va a seguir ejecuntando la condicion es que counter sea menor a 10 importante si no le indicamos
	//un sumador de counter no va a terminar el ciclo for porque siempre va a permanecer 0 debemos indicarle un
	//incremental de counter plus (counter++) en el ejemplo counter es igal a 0 pero le damos un ++ al finalizar el codigo va a sumarle unoal codigo
	//tip para hacer un salto de linea se hace ("\n")
	//el ultimo ejemplo da hasta 9 porque lo condicionamos hasta menor de 10

	for i := 0; i <= 10; i++ {
		fmt.Println(i)

		fmt.Printf("\n")
	}
	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	//for forvever que va hasta la enternidad seria counterforever := 0 for { } seva a jecutar por siempre porque no le estamos idicando
	//hasta dondese detiene con un control c es el estandar para detener cualquier comando en una terminal de linux
	//debemos tener en cuanta que hay que tener una manera que dentro del bloque del codigo pueda escaparce o salirce de ese ciclo for
	//porque si fuera a produccion ese ciclo va a estar consumiendo servidor y no lo va a poder parar hasta que se tumbe el codigo
	//counterForever :=0
	//for{
	//	fmt.Println(counterforever)
	//	counterForever++
	// no se ejecuta para no saturar la maquina
	// tarea hacer un cilco for  donde el i sea un numero determinado iniciando en 0 y que pueda ir decreciendo
	//operadores logicos y de comparacion
	// ver esta guia https://devhints.io/go
	// Son operadores que nos permiten hacer una comparación de condiciones y en caso de cumplirse como
	//sino ejecutarán un código determinado. Si se cumple es VERDADERO/TRUE y si no se cumple son
	//FALSO/FALSE.
	//Operadores de comparación
	//Son aquellos que retornan TRUE o FALSE en caso de cumplirse o no una expresión. Son los siguientes
	//valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
	//valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
	//valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
	//valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
	//valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
	//valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.
	//Operadores lógicos
	//Son aquellos que retorna TRUE o FALSE si cumplen o no una condición utilizando puertas lógicas.
	//Operador AND:
	//Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como
	//TRUE. En Go, se utiliza este símbolo &&.
	//Ejemplo1: 1>0 && 2 >0 Esto retornará TRUE porque tanto la primera como la segunda condición son
	//verdaderas.
	//Ejemplo2: 2<0 && 1 > 0 Esto retornará FALSE porque una de las condiciones no es verdadera.
	//Operador OR:
	//Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En
	//Go, se representa con el símbolo ||.
	//Ejemplo: 2<0 || 1 > 0 Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.
	//Operador NOT:
	//Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:
	//myBool :=  true
	//fmt.Println(!myBool) // Esto retornará false
	//El condicional if

}
