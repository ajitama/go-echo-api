package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-echo-api/openapi"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	ERR_NO_AUTH_TOKEN = fmt.Errorf("NO AUTH")
	ERR_INVALID_TOKEN = fmt.Errorf("INVALID TOKEN")

	ERR_UNKOWN = openapi.CommonHeader{ErrorMessage: "予期せぬエラーが発生", Status: "500"}
	ERR_DBERR  = openapi.CommonHeader{ErrorMessage: "DB接続でエラーが発生", Status: "500"}
	NORMAL_END = openapi.CommonHeader{ErrorMessage: "", Status: "200"}
)

func ErrorResult(ctx echo.Context, status int, message string) error {
	// 共通レスポンスヘッダ
	header := openapi.CommonHeader{
		Status:       strconv.Itoa(status),
		ErrorMessage: message,
	}
	var responseBody []interface{}

	result := openapi.CommonResult{
		Header:   header,
		Response: responseBody,
	}

	return ctx.JSON(status, result)
}

func ErrorMessageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 処理
		err := next(ctx)
		if err != nil {
			if err == ERR_INVALID_TOKEN {
				return ErrorResult(ctx, http.StatusBadRequest, "トークンが正しくない")
			} else if err == ERR_NO_AUTH_TOKEN {
				return ErrorResult(ctx, http.StatusBadRequest, "トークンがない")
			}
			println(err.Error())
			return ErrorResult(ctx, http.StatusBadRequest, err.Error())
		}
		return nil
	}
}

func AuthCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		// Headerからキーを取得
		//_, ok1 := req.Header["Cid"]
		_, ok2 := req.Header["Apptoken"]
		if !ok2 {
			// エラー
			output := fmt.Sprintf("%#v", req.Header)
			fmt.Println([]byte(output + "\n"))
			return ErrorResult(ctx, http.StatusUnauthorized, "トークンがありません")
		}
		empNo := ""
		if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
			// Bodyを変数に取得
			body, err := ioutil.ReadAll(ctx.Request().Body)
			if err != nil {
				return ErrorResult(ctx, http.StatusUnauthorized, "メンバーが存在しません")
			}
			// emp_no取得用構造体
			type jsonMemberId struct {
				EmpNo string `json:"emp_no"`
			}
			var mem jsonMemberId
			json.Unmarshal(body, &mem)
			empNo = mem.EmpNo
			// Bodyを変数に再セット
			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			// Getリクエスト時のemp_no取得
			empNo = ctx.Param("empNo")
			Println("empNo: ", empNo)
		}
		//// mapps_login_token用
		//type MappsLoginToken struct {
		//	MemberId string `json:"member_id"`
		//}
		//DB チェック
		//var count int64
		//var mappsLoginToken MappsLoginToken

		//db := models.GetDBInstance()
		// TODO: expiration_dateの期限判定(現日時時刻より先の有効期限があるか）
		//db.Table("mapps_login_token").Unscoped().Where("app_token = ? AND cid = ? AND expiration_date > now()", apptoken[0], cid[0]).Order("id").Find(&mappsLoginToken).Count(&count)

		// ログインチェック
		//if mappsLoginToken.MemberId == memberId {
		//fmt.Println("トークンが正しい")
		//} else {
		//	Println("有効なトークン情報が見つからない")
		//	return ErrorResult(ctx, http.StatusUnauthorized, fmt.Sprintf("有効なトークン情報が見つからない %s, %s, %s", memberId, mappsLoginToken.MemberId, apptoken[0]))
		//}
		return next(ctx)
	}
}

// IsNumberOnly 文字列中の数字判定
func IsNumberOnly(str string) bool {
	for _, r := range str {
		if '0' <= r && r <= '9' {
			return true
		}
	}
	return false
}
