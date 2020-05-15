package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Not an efficient way to do all these structs, but just a quick way for learning purposes
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {

	// New echo instance
	e := echo.New()

	e.GET("/", home)
	e.GET("/cats/:data", getCats)

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamster)

	e.Start(":8000")
}

// Echo context holds a request and response
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Home page")
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

// Native way to do this - fastest
func addCat(c echo.Context) error {
	cat := Cat{}

	// Note to self: A defer statement defers the execution of a function
	// until the surrounding function returns.
	// So this will close the request after this func returns
	defer c.Request().Body.Close()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addCats, %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Failed unmarshalling in addCats, %v", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat")
}

// Native way to do this - fast
func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request, %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog")
}

// 3rd party way (the echo way) - slow
func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addDog request, %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "We got your hamster")
}
