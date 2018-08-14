package ownStruct

import "fmt"

type Serie struct {
	Name               string
	CreadorDelPrograma string
	Géneros            []string
	Premios            []string
	Autores            []string
}

func Structures() *Serie {
	// curriculum := Serie{Name: "Jhon", Age: 21, Studies: []string{"Machine Learnig", "Python"}}
	blackmirror := new(Serie)
	blackmirror.Name = "Black Mirror"
	blackmirror.CreadorDelPrograma = "Charlie Brooker"
	blackmirror.Géneros = []string{"Ciencia ficción", "Ficción utópica y distópica", "Sátira", "Thriller psicológico", "Serie de antología"}
	blackmirror.Premios = []string{"Premio Peabody"}
	blackmirror.Autores = []string{"Charlie Brooker", "Jesse Armstrong", "Will Bridges"}

	return blackmirror
}

func (s Serie) Subscribe(name string) {
	fmt.Printf("* %s ahora está viendo %s \n", name, s.Name)
}
