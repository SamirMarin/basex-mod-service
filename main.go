package main

import (
	"github.com/SamirMarin/basex-service/internal/basex"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/convert", convert)
	e.Logger.Fatal(e.Start(":1323"))
}

func convert(c echo.Context) error {
	base := basex.Basex{}
	if err := c.Bind(&base); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	convNum := base.Convert()

	return c.String(http.StatusOK, convNum)
}
