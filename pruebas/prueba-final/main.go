package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type direction int

const (
	up direction = iota
	right
	left
	down
)

var (
	fruit  = "☼"
	player = "☺"
)

var (
	errGameOver = errors.New("GAME OVER")
)

func main() {

	var nColumnas, nFilas, points, tiempo int

	d := right

	nColumnas = 10
	nFilas = 10

	a := generarArena(nColumnas, nFilas)

	xPlayer, yPlayer := generarPlayer(a, nColumnas)

	xFruit, yFruit := ubicacionFruit(nColumnas, nFilas)

	generarFruit(a, nColumnas, xFruit, yFruit)

	clearScreen()

	mostrarArena(a, tiempo, points)

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

			generarFruit(a, nColumnas, xFruit, yFruit)

			die := playerMove(a, d, nColumnas, nFilas, points, &xPlayer, &yPlayer, &xFruit, &yFruit)

			if die == true {
				fmt.Println("GAME OVER!!!")
				return
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

func generarFruit(a [][]string, nColumnas, x, y int) {

	n := make([]string, nColumnas)

	n[x] = fruit

	a[y] = n

}

func generarPlayer(a [][]string, nColumnas int) (x, y int) {

	x = 0
	y = 0

	n := make([]string, nColumnas)

	n[x] = player

	a[y] = n

	return
}

func playerMove(a [][]string, d direction, nColumnas, nFilas, points int, xPlayer, yPlayer, xFruit, yFruit *int) (die bool) {

	switch d {
	case up:

		n := make([]string, nColumnas)

		if *yPlayer-1 < 0 {

			die = true
			return
		}

		*yPlayer = *yPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case right:

		n := make([]string, nColumnas)

		if *xPlayer+1 >= len(n) {

			die = true
			return
		}

		*xPlayer = *xPlayer + 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case left:

		n := make([]string, nColumnas)

		if *xPlayer-1 < 0 {

			die = true
			return
		}

		*xPlayer = *xPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case down:

		n := make([]string, nColumnas)

		if *yPlayer+1 >= len(n) {

			die = true
			return
		}

		*yPlayer = *yPlayer + 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	}

	return
}

func checkIFExistFruit(old, new []string) {

	for i := range old {
		if old[i] != "" {
			new[i] = fruit
			return
		}
	}

}

func reubication(a [][]string, n []string, nColumnas, nFilas, points int, xPlayer, yPlayer, xFruit, yFruit *int) {

	if n[*xPlayer] != "" {

		n[*xPlayer] = player

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
