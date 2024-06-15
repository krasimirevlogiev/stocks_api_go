
package handlers

import (
    "encoding/json"
    "net/http"
    //"stock_prices_api/models"
    "stock_prices_api/utils"
    "strings"
    "github.com/gorilla/mux"
)

func GetStockPrices(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    company := vars["company"]

    if company == "" {
        http.Error(w, "company parameter is required", http.StatusBadRequest)
        return
    }

    companySymbols := map[string]string{
        "apple":  "AAPL",
        "google": "GOOGL",
        "ibm":    "IBM",
    }

    symbol, exists := companySymbols[strings.ToLower(company)]
    if !exists {
        http.Error(w, "unknown company", http.StatusBadRequest)
        return
    }

    stocks, err := utils.FetchStockPrices([]string{symbol})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(stocks)
}

