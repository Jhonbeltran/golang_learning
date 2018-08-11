package main

import "fmt"

func main() {
	var number int
	fmt.Print("Ingresa un número: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Println("Es Par")
	} else if number == 3 {
		fmt.Println("Es 3!!!")
	} else {
		fmt.Println("Es impar")
	}

	if animal := "Camello"; animal == "camello" {
		fmt.Println("Diferentes")
	} else {
		fmt.Println("Iguales")
	}
}
