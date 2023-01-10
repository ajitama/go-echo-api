package api

import (
	"fmt"
	"go-echo-api/openapi"
	"go-echo-api/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// PostApiV2MyshopDelete マイ店舗削除
func (api Handler) PostApiV2MyshopDelete(ctx echo.Context) error {
	inputJson := new(openapi.RequestMyShopParams)
	ctx.Bind(&inputJson)

	util.Println("リクエストパラメータ: ", inputJson)
	// 会員ID取得
	strEmpNo := fmt.Sprintf("%v", inputJson.EmpNo)
	// 対象テーブル名の生成
	tableName := "myshop"

	var deleteStore []openapi.ShopList
	err := api.DB.Transaction(func(tx *gorm.DB) error {
		for _, shopCode := range inputJson.ShopList {
			var count int64
			var shop_name string
			tx.Table(tableName).Unscoped().Where("emp_no = ? and shop_code= ?", strEmpNo, shopCode).Count(&count)
			//response用
			tx.Table("myshop").Select("shop_name").Where("shop_code = ?", shopCode).Find(&shop_name)
			if int(count) != 0 {
				// 0件以外なら削除する
				var deleteShop RegisterMyShop
				deleteShop.EmpNo = strEmpNo
				deleteShop.ShopCode = shopCode
				// クエリ実行
				tx.Table(tableName).Unscoped().Where("emp_no = ? and shop_code = ?", strEmpNo, shopCode).Delete(&deleteShop)
				// 削除した店舗コードをセットする
				deleteStore = append(deleteStore, openapi.ShopList{ShopCd: shopCode, Name: shop_name})
			} else {
				util.Println("[マイ店舗削除]会員番号 :", strEmpNo, "無効な店舗:", shopCode)
				continue
			}
		}
		return nil
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, openapi.ResultMyShop{
			Header: openapi.CommonHeader{
				"店舗の削除に失敗しました",
				"501",
			},
		})
	}

	if len(deleteStore) == 0 {
		// 削除件数0の場合は空箱をつける
		deleteStore = make([]openapi.ShopList, 0)
	}
	result := openapi.ResultMyShop{
		Header:   util.NORMAL_END,
		Response: deleteStore,
	}
	return ctx.JSON(http.StatusOK, result)
}
