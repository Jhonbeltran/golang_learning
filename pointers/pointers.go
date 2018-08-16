package pointers

import "fmt"

func PrintPointers() {
	variable := 100
	var puntero *int
	puntero = &variable
	fmt.Println("Sin Modificar")
	fmt.Println(variable, puntero)
	fmt.Println(variable, *puntero)
	fmt.Println(&variable, puntero)
	*puntero = 300
	fmt.Println("Modificado")
	fmt.Println(variable, puntero)
	fmt.Println(variable, *puntero)
	fmt.Println(&variable, puntero)
}
