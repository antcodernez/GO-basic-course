package main

import "fmt"

func main() {
	// DECLARACIÓN DE CONSTANTES
	const pi float64 = 3.1415
	// palabra reservada, nombre, float64 o float32
	// float32 Almacena números mas pequeños
	//float64 Puede Almacenar números mas grandes
	const pi2 = 3.14
	// No estoy estoy especificando el tipo de dato, es otra forma de declarar variables

	fmt.Println("El valor de pi 1 es: ", pi)
	fmt.Println("El valor de pi 2 es: ", pi2)
	// concatenamos con comas "," como en c++

	// DECLARACIÓN DE VARIABLES
	// Integer
	// forma 1 No le digo el tipo de dato, pero la creo y le asigno un valor y me parsea el tipo de dato
	base := 12
	// Los dos puntos indican que no ha sido declarada antes y cuando use ":=" me va a crear la variable y le va a asignar ese valor

	// forma 2 Declaramos, le decimos el tipo de dato y le asignamos un valor
	var altura int = 14
	// declaramos la varibale, le indicamos el tipo de variable que va a ser y le asignamos el valor

	// forma 3 Declaramos la variable y le decimos el tipo de dato, sin asignar un valor
	var area int

	// NOTA: GO NO COMPILA SI LAS VARIABLES NO SON USADAS
	fmt.Println(base, altura, area)

	// Zero values: Es el valor que va a tener por defecto una variable sino se le asigna un valor
	var a int     // 0
	var b float64 // 0
	var c string  // "" <-- nada
	var d bool    // false

	fmt.Println(a, b, c, d)

	// area cuadrado
	fmt.Println("Calculo del area de un cuadrado")
	base_1 := 10
	result := base_1 * base_1

	fmt.Println("El lado de tu cuadrado es:", base_1, "El area de tu cuadrado es: ", result)
}
