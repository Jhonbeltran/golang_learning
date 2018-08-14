package main //Forma de crear modulos en go

import (
	"fmt"

	"github.com/Jhonbeltran/gocurso/arrays"
	"github.com/Jhonbeltran/gocurso/basic_switch"
	"github.com/Jhonbeltran/gocurso/maps"
	"github.com/Jhonbeltran/gocurso/stats"
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
	fmt.Print(a, b, c)
	fmt.Print("\n")
	stats.Stats()
	arrays.Examples()
	basic_switch.Basic_switch()
	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetAuthor("Lloverá y yo veré"))

}
