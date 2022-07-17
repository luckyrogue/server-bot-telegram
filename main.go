package main

import (
	"openlog/tgclient"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env-private")
	tgclient.Run()
}
