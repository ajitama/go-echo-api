package api

import (
	"fmt"
	"go-echo-api/openapi"
	"go-echo-api/util"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MyShop struct {
	gorm.Model
	Id       int
	MemberId string
	ShopCd   string
	Name     string
}

type RegisterMyShop struct {
	MemberId  string
	ShopCd    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetApiV2MyshopMemberId マイ店舗 一覧取得
func (api Handler) GetApiV2MyshopMemberId(ctx echo.Context, param openapi.GetMyMemberId) error {
	var memberId string
	_ = runtime.BindStyledParameterWithLocation("simple", false, "memberId", runtime.ParamLocationPath, ctx.Param("memberId"), &memberId)
	tableName := "myshop"
	// テーブルへのクエリ
	var count int64
	var memberShop []MyShop
	var myShop []openapi.ShopList
	joinQuery := fmt.Sprintf("INNER JOIN shop_official ON shop_official.shop_cd = %s.shop_cd", tableName)
	api.DB.Unscoped().Table(tableName).Select("shop_official.shop_cd", "name").Where("member_id = ?", memberId).Joins(joinQuery).Find(&memberShop).Count(&count)

	for i := 0; i < int(count); i++ {
		myShop = append(myShop, openapi.ShopList{ShopCd: memberShop[i].ShopCd, Name: memberShop[i].Name})
	}
	if int(count) == 0 {
		// 0件のときは空箱だけ作成
		myShop = make([]openapi.ShopList, 0)
	}
	result := openapi.ResultMyShop{
		Header:   util.NORMAL_END,
		Response: myShop,
	}

	return ctx.JSON(http.StatusOK, result)
}

// PostApiV2Myshop マイ店舗登録
func (api Handler) PostApiV2Myshop(ctx echo.Context) error {
	inputJson := new(openapi.RequestMyShopParams)
	ctx.Bind(&inputJson)

	//debug log 出力
	util.Println("リクエストパラメータ", inputJson)
	if len(inputJson.ShopList) > 100 {
		return ctx.JSON(http.StatusOK, openapi.ResultMyShop{
			Header: util.ERR_UNKOWN,
		})
	}

	memberId := fmt.Sprintf("%v", inputJson.MemberId)
	var registerStore []openapi.ShopList

	err := api.DB.Transaction(func(tx *gorm.DB) error {
		for _, shopCode := range inputJson.ShopList {
			var count int64
			var shop_name string
			tx.Table("shop_official").Select("name").Where("shop_cd = ? AND now() BETWEEN block1_from AND block1_to", shopCode).Find(&shop_name).Count(&count)

			if count == 0 {
				util.Println("[マイ店舗登録] 会員番号 :", memberId, "無効な店舗:", shopCode)
				continue
			}

			//登録済みか確認
			tableName := "myshop"
			tx.Table(tableName).Unscoped().Where("member_id = ? AND shop_cd = ?", memberId, shopCode).Count(&count)
			if int(count) == 0 {
				// 0件なら登録する
				tx.Table(tableName).Unscoped().Create(&RegisterMyShop{
					MemberId: memberId,
					ShopCd:   shopCode,
				})
				// 登録した店舗コードをセットする
				registerStore = append(registerStore, openapi.ShopList{ShopCd: shopCode, Name: shop_name})
			}
		}
		return nil
	})
	if err != nil {
		//Todo
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

// PostApiV2MyshopDelete マイ店舗削除
func (api Handler) PostApiV2MyshopDelete(ctx echo.Context) error {
	inputJson := new(openapi.RequestMyShopParams)
	ctx.Bind(&inputJson)

	util.Println("リクエストパラメータ", inputJson)
	// 会員ID末尾1文字の取得
	memberId := fmt.Sprintf("%v", inputJson.MemberId)
	// 対象テーブル名の生成
	tableName := "myshop"

	var deleteStore []openapi.ShopList
	err := api.DB.Transaction(func(tx *gorm.DB) error {
		for _, shopCode := range inputJson.ShopList {
			var count int64
			var shop_name string
			tx.Table(tableName).Unscoped().Where("member_id = ? and shop_cd = ?", memberId, shopCode).Count(&count)
			//response用
			tx.Table("shop_official").Select("name").Where("shop_cd = ? AND now() BETWEEN block1_from AND block1_to", shopCode).Find(&shop_name)
			if int(count) != 0 {
				// 0件以外なら削除する
				var deleteShop RegisterMyShop
				deleteShop.MemberId = memberId
				deleteShop.ShopCd = shopCode
				// クエリ実行
				tx.Table(tableName).Unscoped().Where("member_id = ? and shop_cd = ?", memberId, shopCode).Delete(&deleteShop)
				// 削除した店舗コードをセットする
				deleteStore = append(deleteStore, openapi.ShopList{ShopCd: shopCode, Name: shop_name})
			} else {
				util.Println("[マイ店舗削除]会員番号 :", memberId, "無効な店舗:", shopCode)
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
