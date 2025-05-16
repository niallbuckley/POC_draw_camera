package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func getApiKey(c echo.Context) error {
     val := os.Getenv("GOOGLE_MAPS_API")
     return c.JSON(http.StatusOK, map[string]string{"apiKey": val,})
}

type sheetData struct {
	Azimuth string `json:"azimuth"`
	Fov int `json:"fov"`
	Height int `json:"height"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Tilt int `json:"tilt"`
}

func uploadSheetData(c echo.Context) error {
	var data sheetData

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Marshal struct to JSON
	sheetDataJson, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to marshal JSON"})
	}
	
	resp, err := http.Post(
		// could get the id of the google sheet from user or env var as well
		"https://script.google.com/macros/s/AKfycbzQ3QVWlcUGYl-pSN9Fct4SCLg7yvtNyak372qxCR84ZQnV7vUsSgDx4A6T6944TPuL/exec", 
		"application/json", 
		bytes.NewBuffer(sheetDataJson),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send data to Google Sheets"})
	}
	defer resp.Body.Close()

	return c.JSON(http.StatusOK, map[string]string{"status": "Data sent to Google Sheets"})
}

func main() {
	e := echo.New()
	e.Static("/", "../frontend")

	// Enable CORS for all origins (for local dev)
	e.Use(middleware.CORS())

    e.GET("/*", func(c echo.Context) error {
	 if !strings.HasPrefix(c.Path(), "/api") {
	   return c.File("../frontend/camera-map.html")
      } else {
		return nil
	  }
	})

	// APIs
	e.GET("/api/mapKey", getApiKey)
	e.POST("/api/sheetData", uploadSheetData)

	// Start server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
