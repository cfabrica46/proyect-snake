package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	fruta  = "☼"
	player = "☻"
)

func main() {

	var nColumnas, points, tiempo, nFilas int

	nColumnas = 10
	nFilas = 10

	a := generarArena(nColumnas, nFilas)

	tick := time.Tick(time.Second * 1)

	x, y := ubicacionFruta(nColumnas, nFilas)

	generarFruta(a, nColumnas, x, y)

	c := make(chan bool)

	go comer(c)

	for {
		select {
		case <-tick:

			select {
			case <-c:

				a = generarArena(nColumnas, nFilas)

				x, y = ubicacionFruta(nColumnas, nFilas)

				generarFruta(a, nColumnas, x, y)

				points++

			default:

			}
			clearScreen()

			mostrarArena(a, tiempo, points)

			tiempo++

		}

	}

}

func generarArena(nColumnas, nFilas int) (matriz [][]string) {

	slice := make([]string, nFilas)

	for i := 0; i < nColumnas; i++ {

		matriz = append(matriz, slice)
	}

	return
}

func mostrarArena(a [][]string, tiempo, points int) {

	fmt.Println()

	puntaje := fmt.Sprintf("Points: %d", points)
	t := fmt.Sprintf("Time: %d", tiempo)

	fmt.Println(puntaje)
	fmt.Println(t)

	fmt.Println()
	fmt.Println()

	for i := range a {
		for index := range a[i] {
			if a[i][index] == "" {
				fmt.Print("■  ")
			} else {
				fmt.Printf("%s  ", a[i][index])
			}
		}
		fmt.Println()
	}
}

func ubicacionFruta(nColumnas, nFilas int) (x, y int) {

	rand.Seed(time.Now().UnixNano())

	x = rand.Intn(nColumnas)

	y = rand.Intn(nFilas)

	return
}

func generarFruta(a [][]string, nColumnas, x, y int) {

	n := make([]string, nColumnas)

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
