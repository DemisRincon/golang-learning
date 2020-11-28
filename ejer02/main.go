package main

import (
	"fmt"
	"strconv"
)

//asignacion explicita
var numer int
var texto string
var status bool = true

func main() {
	//asignacion con valores en una linea
	num4, num5, num6, texto, text2, status := 2, 5, 4, "este es mi texto", "texto2", false
	//conversion de tipos
	var moneda int64 = 0
	num4 = int(moneda)
	//convertir ints a texto
	texto = fmt.Sprintf("%d", moneda)
	text2 = strconv.Itoa(int(moneda))
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(texto)
	fmt.Println(text2)
	fmt.Println(status)
	showStatus()

}

func showStatus() {
	fmt.Println(status)
}
