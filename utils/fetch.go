
package utils

import (
    "encoding/json"
    "net/http"
    "os"
    "stock_prices_api/models"
    "fmt"
    "strconv"
)

func FetchStockPrices() ([]models.Stock, error) {
    apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("API key not set")
    }

    url := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&apikey=" + apiKey
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
        return nil, err
    }

    // Parse the JSON response to extract stock data
    var stocks []models.Stock
    if timeSeries, ok := result["Time Series (Daily)"].(map[string]interface{}); ok {
        for _, data := range timeSeries {
            if stockData, ok := data.(map[string]interface{}); ok {
                price, err := strconv.ParseFloat(stockData["4. close"].(string), 64)
                if err != nil {
                    continue
                }
                stock := models.Stock{
                    Symbol: "IBM",
                    Price:  price,
                }
                stocks = append(stocks, stock)
            }
        }
    }
    return stocks, nil
}
