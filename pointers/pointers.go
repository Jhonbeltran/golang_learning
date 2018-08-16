package pointers

import "fmt"

func PrintPointers() {
	variable := 100
	var puntero *int
	puntero = &variable
	fmt.Println(variable, puntero)
	fmt.Println(variable, *puntero)
	fmt.Println(&variable, puntero)
}
