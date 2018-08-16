package gorutines

import (
	"fmt"
	"time"
)

func Rutines() {
	forGo(300)
	forGo(400)
	time.Sleep(10000 * time.Millisecond)
}

func printGo(index int) {
	fmt.Println("Go Rutine -----> ", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go printGo(i)
	}
}
