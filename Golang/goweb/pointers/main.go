package main

import "fmt"

func main() {

	var v int = 19
	// El anperson se usa para hacer referencia al espacio en memoria
	Increment(&v)
	fmt.Println("The value the la variable v is : ", v)
}

func Increment(v *int) {
	*v++
}
