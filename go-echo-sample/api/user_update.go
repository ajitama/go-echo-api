package api

import (
	"fmt"
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

	fmt.Printf("関数内データ: %+v\n", params.UserData)
	var count int64
	if result := api.DB.Debug().Table("employees").Where("emp_no = ?", params.EmpNo).Count(&count); result.Error != nil {
		response.Header.Status = "500"
		response.Header.ErrorMessage = "DBエラー"
	}
	if count == 0 {
		response.Header.Status = "500"
		response.Header.ErrorMessage = "会員情報が見つかりませんでした"

	} else {
		// DebugでSQLを表示
		res := api.DB.Debug().Table("employees").Where("emp_no = ?", params.EmpNo).Updates(&params.UserData)
		if res.Error != nil {
			util.Println(res.Error)
			response.Header.Status = "400"
			response.Header.ErrorMessage = "会員情報の更新に失敗しました"
		}
	}
	return ctx.JSON(http.StatusOK, response)
}
