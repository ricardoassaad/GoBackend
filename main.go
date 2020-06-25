package main

import (
	"github.com/joho/godotenv"
	"github.com/ricardoassaad/GoBackend/router"
)

func main() {
	godotenv.Load(".env")

	router.Route()
}
