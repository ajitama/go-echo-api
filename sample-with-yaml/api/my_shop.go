package api

import (
	"fmt"
	"go-echo-api/openapi"
	"go-echo-api/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MyShop struct {
	gorm.Model
	Id       int
	EmpNo    string
	ShopCode string
	ShopName string
}

// GetApiV2MyshopMemberId マイ店舗 一覧取得
func (api Handler) GetApiV2MyshopEmpNo(ctx echo.Context, empNo openapi.GetMyEmpNo) error {
	strEmpNo := string(empNo)
	//_ = runtime.BindStyledParameterWithLocation("simple", false, "EmpNo", runtime.ParamLocationPath, ctx.Param("memberId"), &strEmpNo)
	tableName := "myshop"
	// テーブルへのクエリ
	var count int64
	var memberShop []MyShop
	var myShop []openapi.ShopList // openapi/types.gen.goにて定義している構造体
	//shop_mstを結合し店舗名を取得する。
	// Todo: shop_code zero padding
	util.Println(strEmpNo)
	joinQuery := fmt.Sprintf("INNER JOIN shop_mst ON shop_mst.shop_code = %s.shop_code", tableName)
	api.DB.Debug().Unscoped().Table("myshop").Select("shop_mst.shop_code", "shop_name").Where("emp_no = ?", strEmpNo).Joins(joinQuery).Find(&memberShop).Count(&count)

	util.Println("お気に入り店舗数: ", count)
	for i := 0; i < int(count); i++ {
		// 店舗を配列にセットする
		myShop = append(myShop, openapi.ShopList{ShopCd: memberShop[i].ShopCode, Name: memberShop[i].ShopName})
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
