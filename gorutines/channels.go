package gorutines

import "fmt"

func Channels() {
	//Creacion de canal
	ch := make(chan string)
	//Funcion anónima ejecutada dentro de un gorutine
	go func() {
		defer close(ch)
		ch <- "Este es un texto que recibe el canal"
	}()
	fmt.Println(<-ch)

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
	}()

	//Podemos iterar sobre los contenidos de un chan
	for n := range ch1 {
		fmt.Printf("Ch1: %d\n", n)
	}

	//Canales con buffer (limite)
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Printf("Ch2: %d\n", <-ch2)
	fmt.Printf("Ch2: %d\n", <-ch2)
	//Luego de cada impresión liberamos un espacio en el Canal
	ch2 <- 3
	fmt.Printf("Ch2: %d\n", <-ch2)

}
