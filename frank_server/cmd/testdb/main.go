package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	environment = "dev"
	host        = "SQLHOST"
	port        = 5432
	user        = "SQLUSER"
	password    = "SQLPASSWORD"
	dbname      = "SQLDBNAME"
	sslmode     = "SQLSSLMODE"
)

func main() {
	var env map[string]string
	env, err := godotenv.Read(".env." + environment)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		env[host], port, env[user], env[password], env[dbname], env[sslmode])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected to: %s", env[host])
}
