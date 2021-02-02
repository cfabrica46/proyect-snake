package db

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

//User representa los datos de la tabla users de databases.db
type User struct {
	ID       int
	Username string
	Password string
	Scores   []Score
}

//Score representa los datos de la tabla scores de databases.db
type Score struct {
	ID    int
	Score int
	Date  string
}

//Estos errores seran utilizados a lo largo del paquete
var (
	ErrUserExist = errors.New("El username que usted escogió ya esta en uso")
	ErrNotScores = errors.New("Aun no tiene un  marcador registrado")
)

//Migracion es una funcion complementaria de Open
//Al ejecutarla se migraran los datos del archivo .sql a un archivo .db
func Migracion() (databases *sql.DB, err error) {
	archivoDB, err := os.Create("databases.db")

	if err != nil {
		return
	}
	archivoDB.Close()

	databases, err = sql.Open("sqlite3", "./databases.db?_foreign_keys=on")

	if err != nil {
		return
	}

	archivoSQL, err := os.Open("databases.sql")

	if err != nil {
		return
	}

	defer archivoSQL.Close()

	contendio, err := ioutil.ReadAll(archivoSQL)

	if err != nil {
		return
	}

	_, err = databases.Exec(string(contendio))
	if err != nil {
		return
	}

	return
}

//Open Abrira el archivo .db o en su defecto ejecutara Migracion
func Open() (databases *sql.DB, err error) {

	archivo, err := os.Open("databases.db")

	if err != nil {
		if os.IsNotExist(err) {

			databases, err := Migracion()

			if err != nil {
				return databases, err
			}

			return databases, err
		}
		return
	}
	defer archivo.Close()

	databases, err = sql.Open("sqlite3", "./databases.db?_foreign_keys=on")

	if err != nil {
		return
	}

	return
}

//GetUser vericara si existe un user registrado con los parametros predefinidos
func GetUser(databases *sql.DB, usernameScan, passwordScan string) (user *User, err error) {

	var userAux User

	row := databases.QueryRow("SELECT id,username,password FROM users WHERE username = ? AND password = ?", usernameScan, passwordScan)

	err = row.Scan(&userAux.ID, &userAux.Username, &userAux.Password)

	if err != nil {
		return
	}

	user = &userAux

	return

}

//CheckIfUserAlreadyExist verifica si ya existe un usuario registrado con el mismo username
func CheckIfUserAlreadyExist(databases *sql.DB, usernameScan string) (check bool, err error) {

	var id int

	row := databases.QueryRow("SELECT id FROM users WHERE username = ? ", usernameScan)

	err = row.Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		check = true
		return
	}
	check = true
	return
}

//InsertUser inserta a la base de datos el nuevo usuario
func InsertUser(databases *sql.DB, usernameScan, passwordScan string) (err error) {

	stmt, err := databases.Prepare("INSERT INTO users(username, password) VALUES(?,?)")

	if err != nil {
		return
	}
	_, err = stmt.Exec(usernameScan, passwordScan)

	if err != nil {

		return
	}

	return
}

//GetScoresWithUserID da como retorno un slice de pokemon
func GetScoresWithUserID(databases *sql.DB, id int) (scores []Score, err error) {

	rows, err := databases.Query("SELECT scores.score,scores.date FROM users_scores INNER JOIN scores ON users_scores.score_id = scores.id WHERE users_scores.user_id = ? ORDER BY scores.score DESC;", id)

	if err != nil {
		return scores, err
	}

	defer rows.Close()

	for rows.Next() {
		var newScore Score

		err = rows.Scan(&newScore.Score, &newScore.Date)

		if err != nil {
			return
		}

		scores = append(scores, newScore)

	}

	return
}

//InsertNewScore inserta a la base de datos la el ultimo marcador con los puntos de tu ultima partida
func InsertNewScore(databases *sql.DB, userID, newScore int) (err error) {

	stmt, err := databases.Prepare("INSERT INTO scores(score,date) VALUES(?,datetime('now'))")

	if err != nil {
		return
	}

	res, err := stmt.Exec(newScore)

	if err != nil {
		return
	}

	scoreID, err := res.LastInsertId()

	err = InsertNewRelationUserScore(databases, userID, int(scoreID))

	if err != nil {
		log.Fatal(err)
	}
	return
}

//InsertNewRelationUserScore inserta a la base de datos la relacion entre el usuario y su nuevo puntaje
func InsertNewRelationUserScore(databases *sql.DB, userID, scoreID int) (err error) {
	stmt, err := databases.Prepare("INSERT INTO users_scores(user_id, score_id) VALUES(?,?)")

	if err != nil {
		return
	}

	_, err = stmt.Exec(userID, scoreID)

	return
}
