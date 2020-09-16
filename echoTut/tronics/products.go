package tronics

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "phone"}, {2: "tv"}, {3: "laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	var product map[int]string

	for _, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}

			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
		//Vendor          string `json:"vendor" validate:"min=5,max=10"`
		//Email           string `json:"email" validate:"required_with=Vendor,email"`
		//Website         string `json:"website" validate:"url"`
		//Country         string `json:"country" validate:"len=2"`
		//DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}

	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	if err := c.Validate(reqBody); err != nil {
		return err
	}

	//if err := v.Struct(reqBody); err != nil {
	//	return err
	//}

	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}

	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}
