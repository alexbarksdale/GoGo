package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func main() {
	// New echo instance
	e := echo.New()

	e.Use(ServerHeader)

	e.GET("/", home)
	e.GET("/cats/:data", getCats)

	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/hamsters", addHamster)

	e.GET("/login", login)

	// ADMIN GROUPING
	adminGrp := e.Group("/admin")

	// logs the server interaction
	adminGrp.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	adminGrp.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// *Check in the DB if the password valid*
		if username == "bob" && password == "123" {
			return true, nil
		}
		return false, nil
	}))

	adminGrp.GET("/main", mainAdmin)

	// COOKIE GROUPING
	cookieGrp := e.Group("/cookie")
	cookieGrp.Use(checkCookie)

	cookieGrp.GET("/main", mainCookie)

	// JWT GROUPING
	jwtGrp := e.Group("/jwt")
	jwtGrp.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("dontUseThis"),
	}))

	jwtGrp.GET("/main", mainJwt)

	// Starts echo instance
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

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// *check user and pass in db*
	if username == "bob" && password == "123" {
		cookie := &http.Cookie{}

		// Same thing
		// cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "example_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in",
			"token":   token,
		})
	}
	return c.String(http.StatusUnauthorized, "You were not logged in")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"bob",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("dontUseThis"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Secret admin page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You are viewing this because you have a cookie")
}

func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("Username:", claims["name"], "UserID:", claims["jti"])

	return c.String(http.StatusOK, "you are on the super secret page!")
}

// ======= MIDDLEWARE ======

// Custom header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "TestServer/1.0")
		c.Response().Header().Set("Anything", "CrazyHeader")
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you don't have any cookie")
			}
			return err
		}

		if cookie.Value == "example_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "You don't have the right cookie")
	}
}
