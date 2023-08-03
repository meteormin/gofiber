package dotenv

import (
	"github.com/joho/godotenv"
	"log"
)

var IsLoaded = false

// init
// load dotenv
func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		log.Println("failed dotenv load")
	}

	IsLoaded = true
}
