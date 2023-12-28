package main

func main() {
	exercise01()
	exercise02()
	exercise03()
}

func exercise01() {
	var Name string
	Name = "Daniel"
	println("___Ejercicio 1___")
	println(Name)
	println("_________________")
}

func exercise02() {

	var temperature string
	var humidity int8
	var pressure int8
	temperature = "32°"
	humidity = 12
	pressure = 11
	println("___Ejercicio 2___")
	println("Temperature: " + temperature)
	print("Humidity: ")
	println(humidity)
	print("Pressure: ")
	println(pressure)
	println("_________________")
}

func exercise03() {
	var firstName string
	var lastName string
	var age int8
	lastName = "6"
	var driver_license = true
	var height int
	childsNumber := 2
	println("___Ejercicio 2___")
	println(firstName)
	println(lastName)
	println(age)
	println(driver_license)
	println(height)
	println(childsNumber)
	println("_________________")
}
