package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cfabrica46/proyect-snake/databases"
)

func play(db *sql.DB, user databases.User) (err error) {
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

			die := playerMove(a, d, nColumnas, nFilas, &points, &xPlayer, &yPlayer, &xFruit, &yFruit)

			if die == true {

				fmt.Println("GAME OVER!!!")

				scores, err := databases.GetScoresWithUserID(db, user.ID)

				if err != nil {
					if err == databases.ErrNotScores {

					} else {
						return err
					}
				}
				check := checkBestScore(points, scores)

				if check == true {
					fmt.Println("CONGRATULATIONS YOU OBTAIN A NEW RENCORD!!!!")
				}

				err = databases.InsertNewScore(db, user.ID, points)

				return err
			}

			clearScreen()

			mostrarArena(a, tiempo, points)

			tiempo++

		}

	}
}

func generarArena(nColumnas, nFilas int) (matriz [][]rune) {

	slice := make([]rune, nFilas)

	for i := 0; i < nColumnas; i++ {

		matriz = append(matriz, slice)
	}

	return
}

func mostrarArena(a [][]rune, tiempo, points int) {

	fmt.Println()

	puntaje := fmt.Sprintf("Points: %d", points)
	t := fmt.Sprintf("Time: %d", tiempo)

	fmt.Println(puntaje)
	fmt.Println(t)

	fmt.Println()
	fmt.Println()

	for i := range a {
		for index := range a[i] {
			if a[i][index] == 0 {
				fmt.Print("■  ")
			} else {

				fmt.Printf("%v  ", string(a[i][index]))
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

func generarFruit(a [][]rune, nColumnas, x, y int) {

	n := make([]rune, nColumnas)

	n[x] = fruit

	a[y] = n

}

func generarPlayer(a [][]rune, nColumnas int) (x, y int) {

	x = 0
	y = 0

	n := make([]rune, nColumnas)

	n[x] = player

	a[y] = n

	return
}

func playerMove(a [][]rune, d direction, nColumnas, nFilas int, points, xPlayer, yPlayer, xFruit, yFruit *int) (die bool) {

	switch d {
	case up:

		n := make([]rune, nColumnas)

		if *yPlayer-1 < 0 {

			die = true
			return
		}

		*yPlayer = *yPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case right:

		n := make([]rune, nColumnas)

		if *xPlayer+1 >= len(n) {

			die = true
			return
		}

		*xPlayer = *xPlayer + 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case left:

		n := make([]rune, nColumnas)

		if *xPlayer-1 < 0 {

			die = true
			return
		}

		*xPlayer = *xPlayer - 1

		checkIFExistFruit(a[*yPlayer], n)

		reubication(a, n, nColumnas, nFilas, points, xPlayer, yPlayer, xFruit, yFruit)

	case down:

		n := make([]rune, nColumnas)

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

func checkIFExistFruit(old, new []rune) {

	for i := range old {
		if old[i] != 0 {
			new[i] = fruit
			return
		}
	}

}

func reubication(a [][]rune, n []rune, nColumnas, nFilas int, points, xPlayer, yPlayer, xFruit, yFruit *int) {

	if n[*xPlayer] != 0 {

		n[*xPlayer] = player

		*points++

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

func checkBestScore(newScore int, scores []databases.Score) (check bool) {

	for i := range scores {
		if scores[i].Score < newScore {
			check = true
			return
		}
	}

	return
}
