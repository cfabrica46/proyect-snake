package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
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
	fruit  = []string{"☼"}
	player = []string{"☻"}
)

var (
	errGameOver = errors.New("GAME OVER")
)

var a arena

var points, tiempo int

func main() {

	var nColumnas, nFilas int

	d := right

	nColumnas = 10
	nFilas = 10

	a = generarArena(nColumnas, nFilas)

	xPlayer, yPlayer := generarPlayer(nColumnas)

	xFruit, yFruit := ubicacionFruit(nColumnas, nFilas)

	generarFruit(nColumnas, xFruit, yFruit)

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

			generarFruit(nColumnas, xFruit, yFruit)

			err := playerMove(d, nColumnas, nFilas, &xPlayer, &yPlayer, &xFruit, &yFruit)

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
		for index := range a[i] {
			if a[i][index] == nil {
				fmt.Print("■\t")
			} else {
				fmt.Printf("%v\t", a[i][index])
			}
		}
		fmt.Println()
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

func ubicacionFruit(nColumnas, nFilas int) (x, y int) {

	rand.Seed(time.Now().UnixNano())

	x = rand.Intn(nColumnas)

	y = rand.Intn(nFilas)

	return
}

func generarFruit(nColumnas, x, y int) {

	n := make(fila, nColumnas)

	n[x] = fruit

	a[y] = n

}

func generarPlayer(nColumnas int) (x, y int) {

	x = 0
	y = 0

	n := make(fila, nColumnas)

	n[x] = player

	a[y] = n

	return
}

func playerMove(d direction, nColumnas, nFilas int, xPlayer, yPlayer, xFruit, yFruit *int) (err error) {

	switch d {
	case up:

		n := make(fila, nColumnas)

		if *yPlayer-1 < 0 {

			err = errGameOver
			return
		}

		*yPlayer = *yPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(n, nColumnas, nFilas, xPlayer, yPlayer, xFruit, yFruit)

	case right:

		n := make(fila, nColumnas)

		if *xPlayer+1 >= len(n) {

			err = errGameOver
			return
		}

		*xPlayer = *xPlayer + 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(n, nColumnas, nFilas, xPlayer, yPlayer, xFruit, yFruit)

	case left:

		n := make(fila, nColumnas)

		if *xPlayer-1 < 0 {

			err = errGameOver
			return
		}

		*xPlayer = *xPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(n, nColumnas, nFilas, xPlayer, yPlayer, xFruit, yFruit)

	case down:

		n := make(fila, nColumnas)

		if *yPlayer+1 >= len(n) {

			err = errGameOver
			return
		}

		*yPlayer = *yPlayer + 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(n, nColumnas, nFilas, xPlayer, yPlayer, xFruit, yFruit)

	}

	return
}

func checkIFExistFruit(old, new fila) {

	for i := range old {
		if old[i] != nil {
			new[i] = fruit
			return
		}
	}

}

func reubication(n fila, nColumnas, nFilas int, xPlayer, yPlayer, xFruit, yFruit *int) {

	if n[*xPlayer] != nil {

		n[*xPlayer] = casilla{player[0]}

		points++

		*xFruit, *yFruit = ubicacionFruit(nColumnas, nFilas)

	} else {
		n[*xPlayer] = player

	}

	a[*yPlayer] = n

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
