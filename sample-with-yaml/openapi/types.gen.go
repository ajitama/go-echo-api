// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package openapi

import (
	"time"
)

// Defines values for RequestUserUpdateParamsUserDataPrefectures.
const (
	RequestUserUpdateParamsUserDataPrefecturesN01 RequestUserUpdateParamsUserDataPrefectures = "01"

	RequestUserUpdateParamsUserDataPrefecturesN02 RequestUserUpdateParamsUserDataPrefectures = "02"

	RequestUserUpdateParamsUserDataPrefecturesN03 RequestUserUpdateParamsUserDataPrefectures = "03"

	RequestUserUpdateParamsUserDataPrefecturesN04 RequestUserUpdateParamsUserDataPrefectures = "04"

	RequestUserUpdateParamsUserDataPrefecturesN05 RequestUserUpdateParamsUserDataPrefectures = "05"

	RequestUserUpdateParamsUserDataPrefecturesN06 RequestUserUpdateParamsUserDataPrefectures = "06"

	RequestUserUpdateParamsUserDataPrefecturesN07 RequestUserUpdateParamsUserDataPrefectures = "07"

	RequestUserUpdateParamsUserDataPrefecturesN08 RequestUserUpdateParamsUserDataPrefectures = "08"

	RequestUserUpdateParamsUserDataPrefecturesN09 RequestUserUpdateParamsUserDataPrefectures = "09"

	RequestUserUpdateParamsUserDataPrefecturesN10 RequestUserUpdateParamsUserDataPrefectures = "10"

	RequestUserUpdateParamsUserDataPrefecturesN11 RequestUserUpdateParamsUserDataPrefectures = "11"

	RequestUserUpdateParamsUserDataPrefecturesN12 RequestUserUpdateParamsUserDataPrefectures = "12"

	RequestUserUpdateParamsUserDataPrefecturesN13 RequestUserUpdateParamsUserDataPrefectures = "13"

	RequestUserUpdateParamsUserDataPrefecturesN14 RequestUserUpdateParamsUserDataPrefectures = "14"

	RequestUserUpdateParamsUserDataPrefecturesN15 RequestUserUpdateParamsUserDataPrefectures = "15"

	RequestUserUpdateParamsUserDataPrefecturesN16 RequestUserUpdateParamsUserDataPrefectures = "16"

	RequestUserUpdateParamsUserDataPrefecturesN17 RequestUserUpdateParamsUserDataPrefectures = "17"

	RequestUserUpdateParamsUserDataPrefecturesN18 RequestUserUpdateParamsUserDataPrefectures = "18"

	RequestUserUpdateParamsUserDataPrefecturesN19 RequestUserUpdateParamsUserDataPrefectures = "19"

	RequestUserUpdateParamsUserDataPrefecturesN20 RequestUserUpdateParamsUserDataPrefectures = "20"

	RequestUserUpdateParamsUserDataPrefecturesN21 RequestUserUpdateParamsUserDataPrefectures = "21"

	RequestUserUpdateParamsUserDataPrefecturesN22 RequestUserUpdateParamsUserDataPrefectures = "22"

	RequestUserUpdateParamsUserDataPrefecturesN23 RequestUserUpdateParamsUserDataPrefectures = "23"

	RequestUserUpdateParamsUserDataPrefecturesN24 RequestUserUpdateParamsUserDataPrefectures = "24"

	RequestUserUpdateParamsUserDataPrefecturesN25 RequestUserUpdateParamsUserDataPrefectures = "25"

	RequestUserUpdateParamsUserDataPrefecturesN26 RequestUserUpdateParamsUserDataPrefectures = "26"

	RequestUserUpdateParamsUserDataPrefecturesN27 RequestUserUpdateParamsUserDataPrefectures = "27"

	RequestUserUpdateParamsUserDataPrefecturesN28 RequestUserUpdateParamsUserDataPrefectures = "28"

	RequestUserUpdateParamsUserDataPrefecturesN29 RequestUserUpdateParamsUserDataPrefectures = "29"

	RequestUserUpdateParamsUserDataPrefecturesN30 RequestUserUpdateParamsUserDataPrefectures = "30"

	RequestUserUpdateParamsUserDataPrefecturesN31 RequestUserUpdateParamsUserDataPrefectures = "31"

	RequestUserUpdateParamsUserDataPrefecturesN32 RequestUserUpdateParamsUserDataPrefectures = "32"

	RequestUserUpdateParamsUserDataPrefecturesN33 RequestUserUpdateParamsUserDataPrefectures = "33"

	RequestUserUpdateParamsUserDataPrefecturesN34 RequestUserUpdateParamsUserDataPrefectures = "34"

	RequestUserUpdateParamsUserDataPrefecturesN35 RequestUserUpdateParamsUserDataPrefectures = "35"

	RequestUserUpdateParamsUserDataPrefecturesN36 RequestUserUpdateParamsUserDataPrefectures = "36"

	RequestUserUpdateParamsUserDataPrefecturesN37 RequestUserUpdateParamsUserDataPrefectures = "37"

	RequestUserUpdateParamsUserDataPrefecturesN38 RequestUserUpdateParamsUserDataPrefectures = "38"

	RequestUserUpdateParamsUserDataPrefecturesN39 RequestUserUpdateParamsUserDataPrefectures = "39"

	RequestUserUpdateParamsUserDataPrefecturesN40 RequestUserUpdateParamsUserDataPrefectures = "40"

	RequestUserUpdateParamsUserDataPrefecturesN41 RequestUserUpdateParamsUserDataPrefectures = "41"

	RequestUserUpdateParamsUserDataPrefecturesN42 RequestUserUpdateParamsUserDataPrefectures = "42"

	RequestUserUpdateParamsUserDataPrefecturesN43 RequestUserUpdateParamsUserDataPrefectures = "43"

	RequestUserUpdateParamsUserDataPrefecturesN44 RequestUserUpdateParamsUserDataPrefectures = "44"

	RequestUserUpdateParamsUserDataPrefecturesN45 RequestUserUpdateParamsUserDataPrefectures = "45"

	RequestUserUpdateParamsUserDataPrefecturesN46 RequestUserUpdateParamsUserDataPrefectures = "46"

	RequestUserUpdateParamsUserDataPrefecturesN47 RequestUserUpdateParamsUserDataPrefectures = "47"
)

// Defines values for RequestUserUpdateParamsUserDataSex.
const (
	RequestUserUpdateParamsUserDataSexN1 RequestUserUpdateParamsUserDataSex = 1

	RequestUserUpdateParamsUserDataSexN2 RequestUserUpdateParamsUserDataSex = 2
)

// 共通レスポンスヘッダ
type CommonHeader struct {
	// エラー内容
	ErrorMessage string `json:"error_message"`

	// ステータスコード
	Status string `json:"status"`
}

// 共通レスポンス
type CommonResult struct {
	// 共通レスポンスヘッダ
	Header CommonHeader `json:"header"`

	// レスポンス部（Null）
	Response []interface{} `json:"response"`
}

// 共通エラーレスポンス
type CommonResultError struct {
	// 共通レスポンスヘッダ
	Header CommonHeader `json:"header"`

	// レスポンス部（Null）
	Response []interface{} `json:"response"`
}

// エラー情報
type Error struct {
	// エラー内容
	Detail string `json:"detail"`

	// ステータスコード
	StatusCode string `json:"statusCode"`
}

// マイ店舗操作インターフェース
type RequestMyShopParams struct {
	// 会員ID
	MemberId string `json:"member_id"`

	// 登録店舗コード
	ShopList []string `json:"shop_list"`
}

// 会員情報更新
type RequestUserUpdateParams struct {
	// 会員ID
	MemberId string `json:"member_id"`

	// 更新項目一覧
	UserData struct {
		// 生年月日 2000-01-01T00:00:00Z
		Birthday time.Time `json:"birthday"`

		// ニックネーム
		Nickname string `json:"nickname"`

		// 居住地
		Prefectures RequestUserUpdateParamsUserDataPrefectures `json:"prefectures"`

		// 性別 1は男性、2が女性
		Sex RequestUserUpdateParamsUserDataSex `json:"sex"`

		// 更新日
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	} `json:"user_data"`
}

// 居住地
type RequestUserUpdateParamsUserDataPrefectures string

// 性別 1は男性、2が女性
type RequestUserUpdateParamsUserDataSex int

// 失敗
type ResultError struct {
	// エラー情報
	Error Error `json:"error"`

	// 結果
	Result bool `json:"result"`
}

// マイ店舗一覧結果
type ResultMyShop struct {
	// 共通レスポンスヘッダ
	Header CommonHeader `json:"header"`

	// レスポンス部
	Response []ShopList `json:"response"`
}

// 成功
type ResultOK struct {
	// 結果
	Result bool `json:"result"`
}

// メルマガフラグ操作インターフェース
type ResultUserDetail struct {
	// 共通レスポンスヘッダ
	Header CommonHeader `json:"header"`

	// 会員情報詳細
	Response UserData `json:"response"`
}

// 店舗ID配列
type ShopList struct {
	// 店舗名
	Name string `json:"name"`

	// 店舗コード
	ShopCd string `json:"shop_cd"`
}

// 会員情報詳細
type UserData struct {
	// 市区町村
	Address1 string `json:"address1"`

	// 町域
	Address2 string `json:"address2"`

	// マンション名等
	Address3 string `json:"address3"`

	// 住所コード
	AddressCode string `json:"address_code"`

	// 生年月日。フォーマットは"2000-01-01T00:00:00Z"
	Birthday time.Time `json:"birthday"`

	// 解約日
	CancelDate string `json:"cancel_date"`

	// キャリア
	Carrier string `json:"carrier"`

	// こども会員フラグ
	ChildMemberFlag int `json:"child_member_flag"`

	// 登録日
	ContractDate string `json:"contract_date"`

	// 会員ステータス
	ContractStatus string `json:"contract_status"`

	// 端末タイプ(初回登録時)
	DeviceTypeEntry string `json:"device_type_entry"`

	// 端末タイプ（更新時）
	DeviceTypeUpdate string `json:"device_type_update"`

	// 最終アクセス日
	LastContinueTime string `json:"last_continue_time"`

	// メールアドレス
	MailAddress string `json:"mail_address"`

	// 会員ID
	MemberId string `json:"member_id"`

	// ホルダー会員フラグ
	MoscardFlag int `json:"moscard_flag"`

	// 今年度マイモスユーザランク
	MymosRankId int `json:"mymos_rank_id"`

	// 今年度マイモスユーザランク名
	MymosRankName string `json:"mymos_rank_name"`

	// 次年度マイモスユーザランク
	NextMymosRankId int `json:"next_mymos_rank_id"`

	// 次年度マイモスユーザランク名
	NextMymosRankName string `json:"next_mymos_rank_name"`

	// ニックネーム
	Nickname string `json:"nickname"`

	// NOS会員フラグ
	NosFlag int `json:"nos_flag"`

	// NOSID
	NosId string `json:"nos_id"`

	// 旧会員フラグ
	OldMemberFlag int `json:"old_member_flag"`

	// 旧会員ID
	OldUserId string `json:"old_user_id"`

	// パスワード省略フラグ
	OmitPasswordFlag int `json:"omit_password_flag"`

	// 郵便番号
	PostalCode string `json:"postal_code"`

	// 都道府県コード
	Prefectures string `json:"prefectures"`

	// 性別
	Sex int `json:"sex"`

	// セッションID
	Sid string `json:"sid"`

	// 簡単ログインフラグ
	SimpleLoginFlag int `json:"simple_login_flag"`

	// UID
	Uid string `json:"uid"`

	// 経由
	Via string `json:"via"`
}

// GetMyMemberId defines model for getMyMemberId.
type GetMyMemberId string

// PostApiV2MyshopJSONBody defines parameters for PostApiV2Myshop.
type PostApiV2MyshopJSONBody RequestMyShopParams

// PostApiV2MyshopDeleteJSONBody defines parameters for PostApiV2MyshopDelete.
type PostApiV2MyshopDeleteJSONBody RequestMyShopParams

// PostApiV2UserUpdateJSONBody defines parameters for PostApiV2UserUpdate.
type PostApiV2UserUpdateJSONBody RequestUserUpdateParams

// PostApiV2MyshopJSONRequestBody defines body for PostApiV2Myshop for application/json ContentType.
type PostApiV2MyshopJSONRequestBody PostApiV2MyshopJSONBody

// PostApiV2MyshopDeleteJSONRequestBody defines body for PostApiV2MyshopDelete for application/json ContentType.
type PostApiV2MyshopDeleteJSONRequestBody PostApiV2MyshopDeleteJSONBody

// PostApiV2UserUpdateJSONRequestBody defines body for PostApiV2UserUpdate for application/json ContentType.
type PostApiV2UserUpdateJSONRequestBody PostApiV2UserUpdateJSONBody

