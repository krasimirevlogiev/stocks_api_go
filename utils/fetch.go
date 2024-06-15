package utils

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "sort"
    "stock_prices_api/models"
    "strconv"
)

func FetchStockPrices(symbols []string) ([]models.Stock, error) {
    apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("API key not set")
    }

    var stocks []models.Stock
    for _, symbol := range symbols {
        // Log the symbol being used in the HTTP request
        log.Printf("Fetching data for symbol: %s", symbol)

        url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", symbol, apiKey)
        response, err := http.Get(url)
        if err != nil {
            return nil, err
        }
        defer response.Body.Close()

        var result map[string]interface{}
        if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
            return nil, err
        }

        // Debugging: Print the raw response
        log.Printf("Raw response for %s: %v", symbol, result)

        if message, exists := result["Error Message"]; exists {
            return nil, fmt.Errorf("API error: %s", message)
        }

        timeSeries, ok := result["Time Series (Daily)"].(map[string]interface{})
        if !ok {
            return nil, fmt.Errorf("unexpected response format: %v", result)
        }

        // Find the most recent date
        var dates []string
        for dateStr := range timeSeries {
            dates = append(dates, dateStr)
        }

        sort.Strings(dates)
        mostRecentDate := dates[len(dates)-1]

        mostRecentData, ok := timeSeries[mostRecentDate].(map[string]interface{})
        if !ok {
            return nil, fmt.Errorf("unexpected response format for most recent data: %v", result)
        }

        priceStr, ok := mostRecentData["4. close"].(string)
        if !ok {
            return nil, fmt.Errorf("unexpected price format: %v", mostRecentData)
        }

        price, err := strconv.ParseFloat(priceStr, 64)
        if err != nil {
            return nil, err
        }

        stock := models.Stock{
            Symbol: symbol,
            Price:  price,
        }
        stocks = append(stocks, stock)
    }
    return stocks, nil
}
