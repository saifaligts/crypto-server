package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var currentData map[string]interface{}

func init() {
	// Once the server starts, send a single request to get updated data and save it into a map
	loadData()
}

// Call API once and get all data and set in a global variable currentData
func loadData() {
	apiURL := "https://api.hitbtc.com/api/3/public/currency"

	// Fetch data from the API
	err := fetchData(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
}

func fetchData(apiURL string) error {
	// Run the HTTP request and get the response
	response, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Get the body of the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return err
	}
	print(body)

	// Unmarshal the json into a map
	if err := json.Unmarshal(body, &currentData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	return nil
}

func findSymbol(symbol string) (interface{}, error) {
	// Find Data for the selected Symbol
	symbolData, ok := currentData[symbol]
	if !ok {
		return nil, fmt.Errorf("symbol %s not found", symbol)
	}

	return symbolData, nil
}

func main() {
	r := gin.Default()
	r.GET("/currency/:symbol", func(c *gin.Context) {
		symbol := c.Param("symbol")
		if symbol == "" || symbol == "all" {
			c.JSON(http.StatusOK, gin.H{"currencies": currentData})
		} else {
			result, _ := findSymbol(symbol)
			c.JSON(http.StatusOK, gin.H{"currency": result})
		}
	})

	// Start the server
	if err := r.Run(":8082"); err != nil {
		fmt.Println("Error go restarting the server:", err)
	}
}
