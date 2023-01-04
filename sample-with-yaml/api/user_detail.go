package api

import (
	"go-echo-api/openapi"
	"go-echo-api/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 会員情報 参照
// (GET /api/v2/user/detail/{empNo})
func (api Handler) GetApiV2UserDetailEmpNo(ctx echo.Context, empNo openapi.GetMyEmpNo) error {
	response := openapi.ResultUserDetail{
		Header: openapi.CommonHeader{
			Status:       "200",
			ErrorMessage: "",
		},
		Response: openapi.UserData{},
	}
	strEmpNo := string(empNo)
	println(strEmpNo)
	// openapi/types.gen.go内の構造体（UserData）を変数として定義する
	userDetail := openapi.UserData{}
	var count int64
	if result := api.DB.Debug().Table("employees").Where("emp_no = ?", strEmpNo).Count(&count).First(&userDetail); result.Error != nil {
		// record not foundやDBエラーも含まれる。
		response.Header.Status = "500"
		response.Header.ErrorMessage = "指定された会員が見つかりません。"

		return ctx.JSON(http.StatusOK, response)
	}
	util.Println(count)
	// response.ResponseにDBのデータをセット
	response.Response = userDetail
	return ctx.JSON(http.StatusOK, response)
}
