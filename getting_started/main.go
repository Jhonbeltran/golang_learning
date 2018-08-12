package main //Forma de crear modulos en go

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Hola Mundo")
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Has ingresado como: %s --- Bienvenido\n", name)
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)
}
