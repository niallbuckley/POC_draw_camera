package main

import (
	"fmt"
	// "io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func basicGetRequest(c echo.Context) error {
     return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Go Echo!",})
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
	 fmt.Printf("%+v", data)

	 return nil
}

func main() {
	e := echo.New()

	// Enable CORS for all origins (for local dev)
	e.Use(middleware.CORS())


	e.GET("/api/hello", basicGetRequest)
	e.POST("/api/sheetData", uploadSheetData)

	e.Logger.Fatal(e.Start(":8080"))
}
