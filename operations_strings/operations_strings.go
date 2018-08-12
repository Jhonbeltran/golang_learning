package operations_strings

import (
	"fmt"
	"strings"
)

func operations_strings() {
	var text = "hola mundo, hola platzi, Hola Go"
	fmt.Println(text)
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.ToTitle(text))
	//-1 para que cambien todas las ocurrencias
	fmt.Println(strings.Replace(text, "hola", "Hi", -1))

	text_array := strings.Split(text, " ")
	fmt.Println(text_array[0])
	fmt.Println(text_array, len(text_array))

}
