package main

import (
	"fmt"
)

// crear una funcion Greet que reciba un nombre y devuelva un saludo
// firma sugerida:
// func Greet(name string) string
func main() {
	fmt.Println(Greet("apolonia"))
	fmt.Println(Greet2("michael"))

}

func Greet(name string) string {
	var greetings string = "buen dia, " + name
	return greetings
}
func Greet2(name string) string {
	return "buen dia, " + name
}
