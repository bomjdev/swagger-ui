package swaggerui

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, spec *openapi3.T) {
	e.GET("/spec.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, spec)
	})
	e.GET("/docs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, SwaggerHTML)
	})
}
