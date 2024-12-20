package controller

import (
	"fmt"
	"golang-br/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	status := "Success"

	res := &model.Response{
		Status:  status,
		Message: "hellow world",
	}

	return c.JSON(http.StatusOK, res)
}

func Upstream(c echo.Context) error {
	name := c.Param("name")
	res := &model.Response{
		Status:  "Success",
		Message: fmt.Sprintf("hellow %v", name),
	}
	return c.JSON(http.StatusOK, res)
}
