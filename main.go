package main

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)


func main() {
    // Load API Key
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    riotKey := os.Getenv("RIOT_KEY")

    account := r_account_riot_id(riotKey, "americas", "going 8th", "7330")
    fmt.Println(account)
}
