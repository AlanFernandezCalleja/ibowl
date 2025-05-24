package main

import (
	"fmt"
)

func main() {
	message := saludo("Mundo")
	fmt.Println(message)

	for i := 1; i <= 5; i++ {
		fmt.Printf("Contando: %d\n", i)
	}
}

func saludo(nombre string) string {
	return "Hola, " + nombre
}
