package api

import (
	"go-echo-api/openapi"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetApiOpenHealthcheck
func (api Handler) GetApiV2UserDetailMemberId(ctx echo.Context, memberId openapi.GetMyMemberId) error {
	result := openapi.ResultOK{
		Result: true,
	}
	return ctx.JSON(http.StatusOK, result)
}
