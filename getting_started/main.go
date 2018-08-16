package main //Forma de crear modulos en go

import (
	"fmt"

	"github.com/Jhonbeltran/gocurso/arrays"
	"github.com/Jhonbeltran/gocurso/basic_switch"
	"github.com/Jhonbeltran/gocurso/gorutines"
	"github.com/Jhonbeltran/gocurso/maps"
	"github.com/Jhonbeltran/gocurso/operations"
	"github.com/Jhonbeltran/gocurso/ownStruct"
	"github.com/Jhonbeltran/gocurso/pointers"
	"github.com/Jhonbeltran/gocurso/usingDefer"
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
	operations.Operations()
	arrays.Examples()
	basic_switch.Basic_switch()
	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetAuthor("Lloverá y yo veré"))
	blackmirror := ownStruct.CreateSerie()
	blackmirror.Name = "Black Mirror"
	blackmirror.CreadorDelPrograma = "Charlie Brooker"
	blackmirror.Generos = []string{"Ciencia ficción", "Ficción utópica y distópica", "Sátira", "Thriller psicológico", "Serie de antología"}
	blackmirror.Premios = []string{"Premio Peabody"}
	blackmirror.Autores = []string{"Charlie Brooker", "Jesse Armstrong", "Will Bridges"}
	playtest := ownStruct.CreateEpisode()
	playtest.Titulo = "Playtest"
	playtest.Duracion = 57
	playtest.Temporada = 3
	ownStruct.CallSubscribe(blackmirror)
	ownStruct.CallSubscribe(playtest)
	usingDefer.Start()
	usingDefer.StakingDefeders()
	//operations.SumPanic()
	pointers.PrintPointers()
	gorutines.Rutines()
}
