package util

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/labstack/echo/v4"
)

func CommonLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// リクエストされたURIとPOSTされたパラメータを出力する。
		// Todo: AWS FireLens 対応
		req := ctx.Request()
		// "/healthcheck"を含むURIへのリクエストはログを出力しない
		if strings.Index(ctx.Request().RequestURI, "/healthcheck") != -1 {
			return next(ctx)
		}
		// URI出力
		Println("Request:", req.Method, ctx.Request().RequestURI)
		// GET以外はボディを出力
		if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
			body, err := ioutil.ReadAll(ctx.Request().Body)
			if err == nil {
				Println("Request Body:", string(body))
			}

			// Bodyを変数に再セット
			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		return next(ctx)
	}
}
