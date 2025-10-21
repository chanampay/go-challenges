package main

import "fmt"

func main() {
	// TODO: declarar name (string) y age (int), asignarles valores
	// y luego imprimir:
	// Hola, me llamo <name> y tengo <age> años.
	//1) Presentación mínima
	// Objetivo: declarar variables primitivas y concatenar en impresiones.
	// Qué imprimir: Hola, me llamo <nombre> y tengo <edad> años.

	// Restricciones: usar var o :=; sin funciones propias.

	// Starter:
	var name string = "clemenza"
	var age int = 68
	fmt.Println("hola, me llamo " + name)
	fmt.Printf("hola, me llamo %v y tengo %v años\n pwpiro\n", name, age)

}
