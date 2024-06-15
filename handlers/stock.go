package handlers

import (
    "encoding/json"
    "net/http"
    //"stock_prices_api/models"
    "stock_prices_api/utils"
    "github.com/gorilla/mux"
)

func GetStockPrices(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    ticker := vars["ticker"]

    if ticker == "" {
        http.Error(w, "ticker parameter is required", http.StatusBadRequest)
        return
    }

    stocks, err := utils.FetchStockPrices([]string{ticker})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if len(stocks) > 0 {
        json.NewEncoder(w).Encode(stocks[0])
    } else {
        http.Error(w, "no data found for the given ticker", http.StatusNotFound)
    }
}
