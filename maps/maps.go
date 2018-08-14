package maps

//GetMap Retorna un map que es como un dict en python map[llave]valor
func GetMap() (map[string]int, map[string]string) {
	mapTest := make(map[string]int)
	mapTest["peso"] = 65
	mapTest["altura"] = 178

	otherMapTest := map[string]string{
		"Google":  "Bogotá",
		"LML2018": "Bogotá",
		"OWASP":   "Manizales",
	}

	return mapTest, otherMapTest
}
