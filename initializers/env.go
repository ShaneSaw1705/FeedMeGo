package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Env() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
