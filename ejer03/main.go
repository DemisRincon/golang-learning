package main

import "fmt"

var estado bool

func main() {
	estado = true
	// se puede asignar valor en la misma linea y luego oregunrtar
	if estado = false; estado == true {
		fmt.Println(estado)
	} else if estado == false {
		fmt.Println("es falso")
	}
	// se puede asignar valor en la misma linea y luego oregunrtar
	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("mayor a 5")
	}

}
