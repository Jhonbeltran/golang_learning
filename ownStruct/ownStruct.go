package ownStruct

import "fmt"

type Episode struct {
	Titulo    string
	Temporada int
	Duracion  int
}
type Serie struct {
	Name               string
	CreadorDelPrograma string
	Generos            []string
	Premios            []string
	Autores            []string
}

func CreateSerie() *Serie {
	// curriculum := Serie{Name: "Jhon", Age: 21, Studies: []string{"Machine Learnig", "Python"}}
	serie := new(Serie)
	return serie
}

func CreateEpisode() *Episode {
	episode := new(Episode)
	return episode
}

//Para poder subscribirse a episodios y series
type Netflix interface {
	Subscribe(name string)
}

func CallSubscribe(n Netflix) {
	n.Subscribe("Jhon")
}

func (s Serie) Subscribe(name string) {
	fmt.Printf("* %s ahora está viendo %s \n", name, s.Name)
}

func (e Episode) Subscribe(name string) {
	fmt.Printf("* %s ahora está viendo %s \n", name, e.Titulo)
}
