package routes

import (
	"arnov17/echo-test/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, this is echo")
	})

	e.GET("/pegawai", controllers.FetcAllPegawai)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai", controllers.UpdatePegawai)
	e.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
