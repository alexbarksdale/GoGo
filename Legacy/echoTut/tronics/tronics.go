package tronics

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

// init will run before any other method in this package.
func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
	fmt.Printf("Config: %+v", cfg)
}

// Start starts the application.
func Start() {
	// GET
	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)

	// POST
	e.POST("/products", createProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
