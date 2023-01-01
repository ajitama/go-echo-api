package api

import (
	"net/http"

	"go-echo-api/openapi"
	"go-echo-api/util"

	"github.com/labstack/echo/v4"
)

// PostApiV2UserUpdate ユーザー情報更新
func (api Handler) PostApiV2UserUpdate(ctx echo.Context) error {
	response := openapi.CommonResult{
		Header: openapi.CommonHeader{
			Status:       "200",
			ErrorMessage: "",
		},
		Response: nil,
	}
	var params openapi.RequestUserUpdateParams
	ctx.Bind(&params)
	member := openapi.UserData{}
	util.Println(member)
	return ctx.JSON(http.StatusOK, response)
}

// GetApiV2UserTotalAmountMemberId マイモス集計参照
func (api Handler) GetApiV2UserTotalAmountMemberId(ctx echo.Context, param openapi.GetMyMemberId) error {

	return ctx.JSON(http.StatusOK, "ok")
}
