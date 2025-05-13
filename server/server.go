package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Enable CORS for all origins (for local dev)
	e.Use(middleware.CORS())

	// OR to restrict CORS:
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//     AllowOrigins: []string{"http://localhost:3000"},
	//     AllowMethods: []string{http.MethodGet, http.MethodPost},
	// }))

	e.GET("/api/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello from Go Echo!",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
