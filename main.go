package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "Password@123#"
	dbname   = "Users"
)

type User struct {
	Username     string
	PasswordHash string
}

func main() {

}

func setupDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil{
		log.Fatal(err)
	}

	err = db.Ping()
	if(err != nil){
	log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	return db
}

func registerHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST"{
			http.Error(w, "Only POST method is allowed!", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			http.Error(w, "Failed to hash password!", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO Users (username, password_hash) VALUES ($1, $2)", username, string(hashedPassword))

		if err != nil {
				http.Error(w, "Failed to save user!", http.StatusInternalServerError)
				return
		}

		fmt.Fprintf(w, "User registered successfully!")
	}
}