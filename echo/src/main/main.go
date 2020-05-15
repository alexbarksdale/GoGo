package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	// New echo instance
	e := echo.New()
	e.GET("/", hello)
	e.GET("/cats/:data", getCats)
	e.Start(":8000")
}

// Echo context holds a request and response

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Test message")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	switch dataType {
	case "string":
		return c.String(http.StatusOK, fmt.Sprintf("Cat name: %v\nType: %v", catName, catType))
	case "json":
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Bad request made",
		})
	}
}
