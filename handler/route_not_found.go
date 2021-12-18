package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RouteNotFound This handles any other routes that is not
// explicitly specified and returns a 404 http status code
func (h *handler) RouteNotFound(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, "Route not found")
}
