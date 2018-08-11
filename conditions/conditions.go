package main

import "fmt"

func main() {
	iffunc()
	forfunc()
}

func iffunc() {
	var number int
	fmt.Println("********** IF **********")
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

func forfunc() {
	fmt.Println("********** FOR **********")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Print("\n")

	fmt.Println("*** For Similar While ***")
	fmt.Println("condición definida c > 0")
	c := 100
	for c > 0 {
		c -= 10
		fmt.Println("Valor de c = ", c)
	}

	fmt.Println("*** For Similar While ***")
	fmt.Println("Condición definida internamente")
	s := 1000
	for {
		s -= 10
		fmt.Printf("%d, ", s)
		//Condición para terminar
		if s == 0 {
			fmt.Println("\nTermina el for 'infinito' break")
			break
		}
	}
}
