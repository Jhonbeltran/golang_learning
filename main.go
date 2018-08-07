package main //Forma de crear modulos en go

import "fmt"

func main()  {
  var name string
  fmt.Println("Hola Mundo")
  fmt.Print("Ingresa tu nombre: ")
  fmt.Scanf("%s", &name)
  fmt.Println("Has ingresado como: ",name)
}
