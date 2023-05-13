package utils

import (
	"fmt"
	// "os"
	"github.com/joho/godotenv"
)

//  Initial checks for SQL DB
// If there is a .env file, switch to dev mode and set path to profiles.json to root of projet
func Prod_check() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("no dotenv")
	}
}