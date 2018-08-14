package ownStruct

type Cv struct {
	Name    string
	Age     int
	Studies []string
}

func Structures() Cv {
	curriculum := Cv{Name: "Jhon", Age: 21, Studies: []string{"Machine Learnig", "Python"}}
	// curriculum := new(Cv)
	// curriculum.Name = "Jhon"
	// curriculum.Age = 21
	// curriculum.Studies = []string{"Machine Learnig", "Python"}
	return curriculum
}
