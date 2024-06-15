package utils

import (
    "encoding/csv"
    "os"
)

func LoadCompanySymbols(filePath string) (map[string]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    companySymbols := make(map[string]string)
    for _, record := range records[1:] { // Skip header
        companySymbols[record[0]] = record[1]
    }

    return companySymbols, nil
}
