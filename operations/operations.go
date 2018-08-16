package operations

import (
	"errors"
	"fmt"
)

//Some stats (mean actually)
func Operations() {
	mean_values := mean()
	fmt.Println("mean:", mean_values)
}

func SumPanic() {
	number, err := sum("50", 50)

	if err != nil {
		panic(err)
	}
	fmt.Println("La suma es de ", number)
}

//mean recibe 5 elementos para encontrar la media
func mean() int {
	var nums [5]int
	var element int
	fmt.Println("Bienvenidos")
	fmt.Println("Ingresa el conjunto de datos")
	for fill := range nums {
		fmt.Printf("Indice %d: ", fill)
		fmt.Scanf("%d", &element)
		nums[fill] = element
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	mean_values := sum / len(nums)
	return mean_values
}

//Suma dos valores ** Manejo de errores **
func sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		return 0, errors.New("El valor A es un string")
	}
	switch b.(type) {
	case string:
		return 0, errors.New("El valor B es un string")
	}
	return a.(int) + b.(int), nil
}
