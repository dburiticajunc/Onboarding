package main

type Preferences struct {
	Drinks []string
	Foods  []string
}

type Person struct {
	Name    string
	LasName string
	Likes   Preferences
}

type Operation struct {
	Name string
	Do   func()
}

func main() {

	// Instancia de una estructura, y optencion de sus atributos
	dataP := Person{
		Name:    "Daniel",
		LasName: "Buritica",
		Likes: Preferences{
			Drinks: []string{"te"},
			Foods:  []string{"egg"},
		},
	}

	dataPre := Preferences{
		Drinks: []string{"te", "coffe"},
		Foods:  []string{"egg", "white raise"},
	}
	dataP2 := Person{"daniel", "buritica", dataPre}
	println(dataP2.LasName)
	println(dataP.LasName)

	// Estructura que tiene dfinida una funcion como atributo

	dataO := Operation{
		Name: "Something",
		Do: func() {
			println("There is funtion in struct: omething")
		},
	}
	dataO.Do()

}
