package handlers

import (
    "encoding/json"
    "net/http"
    //"stock_prices_api/models"
    "stock_prices_api/utils"
)

func GetStockPrices(w http.ResponseWriter, r *http.Request) {
    stocks, err := utils.FetchStockPrices()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(stocks)
}
