package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/cfabrica46/proyect-snake/databases"

	_ "github.com/mattn/go-sqlite3"
)

type direction int

const (
	up direction = iota
	right
	left
	down
)

var (
	fruit  = rune(9788)
	player = rune(9786)
)

var (
	errUsernamePasswordIncorrect = errors.New("Username y/o Password incorrectos")
)

func main() {

	var ingreso, usernameScan, passwordScan string

	db, err := databases.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Bienvenido")

	fmt.Println("Ingresar || Registrarse [I/R]")

	fmt.Scan(&ingreso)

	ingreso = strings.ToLower(ingreso)

	switch ingreso {
	case "i":

		fmt.Println("Ingrese su username")
		fmt.Scan(&usernameScan)
		fmt.Println("Ingrese su password")
		fmt.Scan(&passwordScan)

		user, err := databases.GetUser(db, usernameScan, passwordScan)

		if err != nil {
			if err == sql.ErrNoRows && user == nil {
				log.Fatal(errUsernamePasswordIncorrect)
			}
			log.Fatal(err)
		}

		err = ingresar(db, *user)

		if err != nil {
			log.Fatal(err)
		}

	case "r":
		user, err := registrar(db)

		if err != nil {
			log.Fatal(err)
		}

		err = ingresar(db, *user)

		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("Error: ELECCIÓN INVALIDA")
	}

}
