package swaggerui

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/getkin/kin-openapi/openapi3"
)

func Register(e *echo.Echo, spec *openapi3.T) {
	e.GET("/spec.json", specHandlerFactory(spec))
	e.GET("/docs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, swaggerHTML)
	})
}

func specHandlerFactory(spec *openapi3.T) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, spec)
	}
}

//go:embed index.html
var swaggerHTML string
