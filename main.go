
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    //"os"
    "stock_prices_api/handlers"
)

func main() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    router := mux.NewRouter()
    router.HandleFunc("/api/stocks", handlers.GetStockPrices).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

