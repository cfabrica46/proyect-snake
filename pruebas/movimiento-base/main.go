package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type casilla []string
type fila []casilla
type arena []fila

type direction int

const (
	up direction = iota
	right
	left
	down
)

var (
	player = []string{"☻"}
)

var a arena

var points, tiempo int

var (
	errGameOver = errors.New("GAME OVER")
)

func main() {

	var nColumnas, nFilas int

	d := right

	nColumnas = 10
	nFilas = 10

	a = generarArena(nColumnas, nFilas)

	x, y := generarPlayer(nColumnas)

	clearScreen()

	mostrarArena()

	tick := time.Tick(time.Second * 1)

	c := make(chan string)

	go scanDirection(c)

	for {
		select {
		case <-tick:

			select {
			case election := <-c:

				convertElection(election, &d)

			default:

			}

			a = generarArena(nColumnas, nFilas)

			err := playerMove(d, nColumnas, &x, &y)

			if err != nil {
				if err == errGameOver {
					fmt.Println(err)
					return
				}

				log.Fatal(err)
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

func generarPlayer(nColumnas int) (x, y int) {

	x = 0
	y = 0

	n := make(fila, nColumnas)

	n[x] = player

	a[y] = n

	return
}

func playerMove(d direction, nColumnas int, x, y *int) (err error) {

	switch d {
	case up:

		n := make(fila, nColumnas)

		if *y-1 < 0 {

			err = errGameOver
			return
		}

		*y = *y - 1

		n[*x] = player
		a[*y] = n

	case right:

		n := make(fila, nColumnas)

		if *x+1 >= len(n) {

			err = errGameOver
			return
		}

		*x = *x + 1

		n[*x] = player
		a[*y] = n

	case left:

		n := make(fila, nColumnas)

		if *x-1 < 0 {

			err = errGameOver
			return
		}

		*x = *x - 1

		n[*x] = player
		a[*y] = n

	case down:

		n := make(fila, nColumnas)

		if *y+1 >= len(n) {

			err = errGameOver
			return
		}

		*y = *y + 1

		n[*x] = player
		a[*y+1] = n

	}

	return
}

func scanDirection(c chan string) {

	var respuesta string

	for {
		fmt.Scan(&respuesta)

		c <- respuesta
	}
}

func convertElection(election string, d *direction) {

	switch election {
	case "w":

		*d = up

	case "d":

		*d = right

	case "a":

		*d = left

	case "s":

		*d = down

	default:

		return

	}

}
