package main

import (
	"custom-quizer-backend/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	err = godotenv.Load(".env")
	if err != nil {
		log.Println("Error reading env file:", err)
	}
	log.Println("HOST:", os.Getenv("host"))
	fmt.Println("hello world")
	db.Init()
}
