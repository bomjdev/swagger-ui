package echo5

import (
	"net/http"

	swaggerui "github.com/bomjdev/swagger-ui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo, spec *openapi3.T) {
	e.GET("/spec.json", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, spec)
	})
	e.GET("/docs", func(c *echo.Context) error {
		return c.HTML(http.StatusOK, swaggerui.SwaggerHTML)
	})
}
