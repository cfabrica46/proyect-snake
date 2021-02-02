package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cfabrica46/snake/db"
)

func ingresar(databases *sql.DB, user db.User) (err error) {

	var eleccionMenu int
	var salir bool

	clearScreen()

	fmt.Printf("Bienvenido %v tu ID es: %v\n", user.Username, user.ID)

	for salir == false {
		fmt.Println("¿Qué Desea Hacer?")
		fmt.Println("1.	Jugar")
		fmt.Println("2.	Mostrar Tus Puntos")
		fmt.Println("0.	Salir")

		fmt.Scan(&eleccionMenu)

		switch eleccionMenu {
		case 1:
			clearScreen()
			err := play(databases, user)

			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("GAME OVER!!!")

		case 2:
			clearScreen()
			scores, err := db.GetScoresWithUserID(databases, user.ID)

			if err != nil {
				fmt.Println(err.Error())
			}

			if len(scores) == 0 {
				fmt.Println(db.ErrNotScores)
			}

			mostrarScores(scores)

			time.Sleep(time.Second * 3)

			clearScreen()

		case 0:
			return
		default:
			fmt.Println("Opcion invalida")
		}
	}
	return
}

func registrar(databases *sql.DB) (user *db.User, err error) {

	var usernameScan, passwordScan string

	fmt.Println("Ingrese su username")
	fmt.Scan(&usernameScan)
	fmt.Println("Ingrese su password")
	fmt.Scan(&passwordScan)

	check, err := db.CheckIfUserAlreadyExist(databases, usernameScan)

	if err != nil {
		return
	}

	if check == true {
		err = db.ErrUserExist
		return
	}

	err = db.InsertUser(databases, usernameScan, passwordScan)
	if err != nil {
		return
	}

	user, err = db.GetUser(databases, usernameScan, passwordScan)

	if err != nil {
		if err == sql.ErrNoRows {
			if user == nil {
				err = errUsernamePasswordIncorrect
				return
			}
			return
		}
		return
	}

	fmt.Println("nuevo usuario", usernameScan)

	return
}

func mostrarScores(scores []db.Score) {

	fmt.Println("N°\tScore\tDate")
	fmt.Println()

	for i := range scores {

		fmt.Printf("%d.\t%v\t%v\n", i+1, scores[i].Score, scores[i].Date)

	}

}
