package routes

import (
	"fmt"
	"go-echo-api/api"
	"go-echo-api/models"
	"go-echo-api/openapi"
	"go-echo-api/util"

	"log"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
)

// RegisterHandlers APIのルーティング設定
func RegisterHandlers(router *echo.Echo, handler api.Handler) {
	wrapper := openapi.ServerInterfaceWrapper{
		Handler: handler}

	// 未ログインでも実行可能
	router.GET("/api/open/healthcheck", wrapper.GetApiOpenHealthcheck)

	// 以下ミドルウェアによる認証が必が必要なAPIグループ
	clientGroup := router.Group("api/v2", util.AuthCheckMiddleware)
	clientGroup.POST("/myshop/delete", wrapper.PostApiV2MyshopDelete)
	clientGroup.POST("/myshop", wrapper.PostApiV2Myshop)
	clientGroup.GET("/myshop/:memberId", wrapper.GetApiV2MyshopMemberId)
	clientGroup.GET("/user/detail/:memberId", wrapper.GetApiV2UserDetailMemberId)
	clientGroup.POST("/user/update", wrapper.PostApiV2UserUpdate)

}

// SetUpEchoServer HTTPサーバーのセットアップ
func SetUpEchoServer(swaggerPath string) (*echo.Echo, error) {
	// DB疎通確認
	db := models.GetDBInstance()
	if db == nil {
		return nil, fmt.Errorf("DB接続エラー")
	}

	// Swaggerのセットアップ
	//https://pkg.go.dev/github.com/do87/oapi-codegen/pkg/middleware#OapiValidatorFromYamlFile
	//  YAML ファイル パスからバリデータ ミドルウェアを作成する(swaggerにパス定義されてないとエラー)
	f, err := middleware.OapiValidatorFromYamlFile(swaggerPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	e := echo.New()
	// API処理を行う際のミドルウェア定義
	// 共通で出力するログ
	e.Use(util.CommonLogMiddleware)
	// エラー発生時の共通レスポンス
	e.Use(util.ErrorMessageMiddleware)
	e.Use(f)

	handler := api.Handler{
		DB: db,
	}
	RegisterHandlers(e, handler)
	return e, nil
}
