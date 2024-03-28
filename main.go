package main

import (
	"database/sql"
	"fmt"
	"log"
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