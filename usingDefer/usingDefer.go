package usingDefer

import "fmt"

//El defer se ejecuta cuando la ejecución de la función que lo llama termina
func Start() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func StakingDefeders() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}

	fmt.Println("done")
}
