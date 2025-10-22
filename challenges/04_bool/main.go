package main

import "fmt"

func main() {
	// 	4) Estados booleanos

	// Objetivo: usar bool y mostrar su valor textual.
	// Qué imprimir: dos líneas:

	// Usuario activo: <true/false>

	// Tiene suscripción: <true/false>
	// TODO: declarar active y subscribed como bool
	// imprimir ambas líneas en orden

	var active bool = true
	var subscribed bool = false
	fmt.Println("usuario activo", active)
	fmt.Println("tiene subscripcion: ", subscribed)

}
