package api

import (
	"go-echo-api/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetApiOpenHealthcheck
func (api Handler) GetApiOpenHealthcheck(ctx echo.Context) error {
	result := openapi.ResultOK{
		Result: true,
	}
	return ctx.JSON(http.StatusOK, result)
}
