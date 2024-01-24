package main

import "fmt"

type Car struct {
	Name  string
	Model string
	price string
}

func (c Car) toString() (result string) {
	result = "The Name del Car is: " + c.Name
	return
}
func main() {

	dataCar := Car{

		Model: "2007",
		price: "1000",
	}

	fmt.Println(dataCar.toString())
}
