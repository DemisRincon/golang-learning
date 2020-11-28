package main

import (
	"fmt"
)

func main() {
	//ciclo for
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	// for untill
	numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("ingresenumero")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}
	//
	var k int
	for k < 10 {
		fmt.Printf("\n valor de i: %d again is %d", k, k)
		if k == 5 {
			fmt.Printf("multiplicamos por 2 \n")
			k = k * 2
			continue
		}
		fmt.Println(" paso por aqui")
		k++
	}
}
