package api

import (
	"go-echo-api/openapi"
	"go-echo-api/util"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RegisterMyShop struct {
	EmpNo     string
	ShopCode  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PostApiV2Myshop マイ店舗登録
func (api Handler) PostApiV2Myshop(ctx echo.Context) error {
	inputJson := new(openapi.RequestMyShopParams)
	ctx.Bind(&inputJson)

	//debug log 出力
	util.Println("リクエストパラメータ:", inputJson)
	// 一度に100件以上はエラーにする
	if len(inputJson.ShopList) > 100 {
		return ctx.JSON(http.StatusOK, openapi.ResultMyShop{
			Header: util.ERR_UNKOWN,
		})
	}

	// 会員IDの抽出
	strEmpNo := inputJson.EmpNo
	var registerStore []openapi.ShopList

	err := api.DB.Transaction(func(tx *gorm.DB) error {
		for _, shopCode := range inputJson.ShopList {
			var count int64
			var shop_name string
			// 店舗マスタの確認（存在確認）
			tx.Table("shop_mst").Debug().Select("shop_name").Where("shop_code = ?", shopCode).Find(&shop_name).Count(&count)

			if count == 0 {
				// 店舗マスタに登録がなかった場合でも処理を続ける。
				util.Println("[マイ店舗登録] 会員番号 :", strEmpNo, "無効な店舗:", shopCode)
				continue
			}

			//登録済みか確認
			tableName := "myshop"
			tx.Table(tableName).Unscoped().Where("emp_no = ? AND shop_code = ?", strEmpNo, shopCode).Count(&count)
			if int(count) == 0 {
				// 0件なら登録する(構造体に値をセットし、DBにInsert)
				tx.Table(tableName).Unscoped().Create(&RegisterMyShop{
					EmpNo:    strEmpNo,
					ShopCode: shopCode,
				})
				// 登録した店舗コードをセットする
				registerStore = append(registerStore, openapi.ShopList{ShopCd: shopCode, Name: shop_name})
			}
		}
		return nil
	})
	if err != nil {
		//Todo : まあ。なんとかしなきゃ。
		return ctx.JSON(http.StatusInternalServerError, openapi.ResultMyShop{
			Header: openapi.CommonHeader{
				"店舗の登録に失敗しました",
				"500",
			},
		})
	}
	if len(registerStore) == 0 {
		// 登録件数0の場合は空箱をつける
		registerStore = make([]openapi.ShopList, 0)
	}
	return ctx.JSON(http.StatusOK, openapi.ResultMyShop{
		Header:   util.NORMAL_END,
		Response: registerStore,
	})
}
