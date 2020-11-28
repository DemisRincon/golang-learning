package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	numero1, numero2, result, leyenda := 0, 0, 0, ""
	fmt.Println("ingrese numero 1")
	fmt.Scanln(&numero1)
	//En windows se detiene el programa en este punto pues el enter contiene \n\r
	//por lo que toma la\r como segundo valor y en linux solo toma \n
	fmt.Println("ingrese numero 2")
	fmt.Scanln(&numero2)
	fmt.Println("Descripcion: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	result = numero2 + numero1
	fmt.Println(leyenda, result)

}
