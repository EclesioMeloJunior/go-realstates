package main

import (
	"go-realstates/server"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")

	app := server.Run(port)

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
