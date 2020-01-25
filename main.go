package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sanket143/lisa/src/system/app"
	DB "github.com/sanket143/lisa/src/system/db"
)

var port string

func init() {
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	s := app.NewServer()

	db := DB.Connect()

	s.Init(port, db)
	s.Start()
}
