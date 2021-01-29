package main

import "fmt"

type casilla []int
type fila []casilla
type arena []fila

func main() {

	var c casilla
	var f fila
	var a arena

	for i := 0; i < 10; i++ {
		f = append(f, c)
	}

	for i := 0; i < 10; i++ {
		a = append(a, f)
	}

	for i := range a {
		fmt.Println(a[i])
	}

}
