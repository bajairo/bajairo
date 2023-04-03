package main

import (
	"fmt"
)

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras hay que declararlos como el la linea 18 para que no diga qu no se uso
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//zero values go asigna un valor de cero pos default en otros idiomas null o algo asi int enteros float es decimales string textos bool es boleanos

	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	//ejercisio Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//suma como es la primera vez que se ejecuta el result se pone : pero ya desúes no hay necesidad

	result := x + y
	fmt.Println("suma:", result)

	// Resta
	result = x - y
	fmt.Println("resta", result)

	//Multiplicación
	result = x * y
	fmt.Println("multiplicación", result)

	// División
	result = y / x
	fmt.Println("División", result)

	//Modulo o residuo que es el remanente, esta operacion es muy comun para hacer los algoritmos
	//cuando se quiera saber si el numero es par o impar en el ejemplo es cero porque es una divicion excacta

	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental es uno de los operdores adicionales usados para ciclos iterativos y es solo añadir un
	//1 a una variable y es colcar 2 veces el signo + como ya lo tenemos definido como 10 le agrega un 1 seria 11

	x++
	fmt.Println("incremental:", x)

	// Decremental es el contrario colocamos 2 veces el -
	x--
	fmt.Println("decremental", x)

	// reto
	// rectangulo, trapesio y de un circulo

	const baseRectangulo = 10
	const alturaRectangulo = 5
	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("Area rectangulo", areaRectangulo)

	const BaseTrapecio = 20
	const basemenorTrapecio = 10
	const alturaTrapecio = 15
	areaTrapecio := basemenorTrapecio * BaseTrapecio * alturaTrapecio
	fmt.Println("Area Trapecio", areaTrapecio)

	const pi3 float64 = 3.14
	const radio = 10
	areaCirculo := pi3 * radio * radio
	fmt.Println("Area Circulo", areaCirculo)

	// valores primitivos
	//numeros enteros depende de el os (32 o 64bits)int8/16/32/64:
	//si no se declara el numero el lo toma por defecto del sistema operativo si es 32 o 64
	//indica el tipo de dato “int” con signo, y además el tamaño del dato.
	//uint8/16/32/64: indica el tipo de dato “int” pero sin signo(solo positivos) y además el tamaño máximo del dato
	//int / uint : toma el tamaño de bits en el que está basado el procesador (con signo / sin signo)
	//float32: 32 bits con signo
	//float64: 64 bits con signo
	//string: se deben declarar con comillas dobles “”
	//bool: true o false
	//complex64: números complejos (real e imaginario float32)
	//complex128: números complejos (real e imaginario float64)
	//declaracion de variables
	// cuando se coloca := quiere decir que la variable no a sido declarada nunca
	//cuando se pone = es porque ya se delaclaro antes
	//FUNCION fmt es la que se usa para imprimir
	//Println indica que imprima en lineas separadas agragando un espacio ejemplo
	hellomessage := "Hello"
	wordlmessage := "world"
	fmt.Println(hellomessage, wordlmessage)
	//Printf agrega una funcion extra al  string que se incerta como valor de entrada
	//en el ejemplo colocamos %s para indicar que el valor que va es un string %d es la forma que le vamos a
	//decir que va a recibir un entero, la forma como se le agrega el salto de linea es con un \ para escapar
	// y la n para indicar es es un salto de linea se agrega una coma, despues del string y se agrega
	// las 2 variables variables, si no sabemos que tipo de dato va si es un string o int colocamos
	//un %v en los dos

	// print f
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//sprintf genera un string pero no lo imprime en consola solo lo guarda como string, con fmt le damos println
	// lo que hace es con ese comando todo lo que resulte generado en el string lo va a guardar a una nueva variable
	// que en ele ejemplo lo llamamos message y el message lo imprimimos directamente y vamos a tener la misma linea
	//no agregamos el \n para que no nos de un salto de linea adicional y quede en limpio
	// aqui unas adicionales https://pkg.go.dev/fmt

	//Springf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//tipo datos informacion muy util con este tipo podemos saber que tipo de datos que pertenece una variable la
	//forma de decir que nos impima la variable es con %T en el siguiente espacio agregamos el hellomessage segun
	//el ejemplo y le agregamos un salto de linea \n

	//tipo dato
	fmt.Printf("hellomessage: %T\n", hellomessage)
	fmt.Printf("cusos: %T\n", cursos)

	// uso de funciones y funciones anonimas se usa cuando queremos cambairs ciertos variables o parametros ayuda a
	//llevar codigo repetitivo a un codigo mas corto en el ejemplo solo queremos cambiar el 1 , 2 , 3
	//la forma de declararla es colocando func y el nombre de la función y la fomra es colocar el primer nombre pero
	//si se va a concatenar se coloca la siguiente en mayuscula se pone () y {}

	fmt.Println("hola mundo")
	fmt.Println("hola mundo 2")
	fmt.Println("hola mundo 3")
}
