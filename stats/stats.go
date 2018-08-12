package stats

import "fmt"

//Some stats (mean actually)
func Stats() {
	mean_values := mean()
	fmt.Println("mean:", mean_values)
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
