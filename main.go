package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    riotKey := os.Getenv("RIOT_KEY")
    fmt.Println(riotKey)
}
