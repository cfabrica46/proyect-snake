package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type casilla []string
type fila []casilla
type arena []fila

var (
	fruta  = []string{"☼"}
	player = []string{"☻"}
)

var a arena

var points, tiempo int

func main() {

	var nColumnas, nFilas int

	nColumnas = 10
	nFilas = 10

	a = generarArena(nColumnas, nFilas)

	tick := time.Tick(time.Second * 1)

	x, y := ubicacionFruta(nColumnas, nFilas)

	generarFruta(nColumnas, x, y)

	c := make(chan bool)

	go comer(c)

	for {
		select {
		case <-tick:

			select {
			case <-c:

				a = generarArena(nColumnas, nFilas)

				x, y = ubicacionFruta(nColumnas, nFilas)

				generarFruta(nColumnas, x, y)

				points++

			default:

			}
			clearScreen()

			mostrarArena()

			tiempo++

		}

	}

}

func generarArena(nColumnas, nFilas int) (a arena) {
	var f fila
	var c casilla

	for i := 0; i < nColumnas; i++ {

		f = append(f, c)
	}

	for i := 0; i < nFilas; i++ {
		a = append(a, f)
	}

	return
}

func mostrarArena() {

	fmt.Println()

	puntaje := fmt.Sprintf("Points: %d", points)
	t := fmt.Sprintf("Time: %d", tiempo)

	fmt.Println(puntaje)
	fmt.Println(t)

	fmt.Println()
	fmt.Println()

	for i := range a {
		fmt.Println(a[i])
	}
}

func ubicacionFruta(nColumnas, nFilas int) (x, y int) {

	rand.Seed(time.Now().UnixNano())

	x = rand.Intn(nColumnas)

	y = rand.Intn(nFilas)

	return
}

func generarFruta(nColumnas, x, y int) {

	n := make(fila, nColumnas)

	n[x] = fruta

	a[y] = n

}

func clearScreen() {

	v := runtime.GOOS

	switch v {
	case "linux":

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

	case "windows":

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func comer(c chan bool) {
	for {
		fmt.Scanln()

		c <- true
	}
}
