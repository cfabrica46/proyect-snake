package main

import (
	"fmt"
)

func main() {

	var matriz [][]rune

	slice := make([]rune, 10)

	for i := 0; i < 10; i++ {

		matriz = append(matriz, slice)
	}

	for i := range matriz {
		fmt.Println(matriz[i])
	}
}
