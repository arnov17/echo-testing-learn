package controllers

import (
	"arnov17/echo-test/helper"
	"arnov17/echo-test/models"
	"net/http"

	"github.com/labstack/echo"
)

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HassPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "berhasil")
}
