package utils

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "sort"
    "stock_prices_api/models"
    "time"
    "strconv"
)

func FetchStockPrices(symbols []string) ([]models.Stock, error) {
    apiKey := os.Getenv("ALPHA_VANTAGE_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("API key not set")
    }

    var stocks []models.Stock
    for _, symbol := range symbols {
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

        timeSeries, ok := result["Time Series (Daily)"].(map[string]interface{})
        if !ok {
            return nil, fmt.Errorf("unexpected response format")
        }

        // Find the most recent date
        var dates []time.Time
        for dateStr := range timeSeries {
            date, err := time.Parse("2006-01-02", dateStr)
            if err != nil {
                continue
            }
            dates = append(dates, date)
        }

        sort.Slice(dates, func(i, j int) bool {
            return dates[i].After(dates[j])
        })

        if len(dates) == 0 {
            return nil, fmt.Errorf("no data available")
        }

        mostRecentDate := dates[0].Format("2006-01-02")
        mostRecentData, ok := timeSeries[mostRecentDate].(map[string]interface{})
        if !ok {
            return nil, fmt.Errorf("unexpected response format for most recent data")
        }

        price, err := strconv.ParseFloat(mostRecentData["4. close"].(string), 64)
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
