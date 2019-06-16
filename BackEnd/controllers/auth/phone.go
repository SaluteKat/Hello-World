package auth

import (
	"BackEnd/lib/httpserver"
	"BackEnd/lib/rsp"
	"BackEnd/model/auth"
	"BackEnd/resource"
)

// GetUserInfo swan.getUserInfo 回调 账户管理示例
func GetPhone(ctx *httpserver.Context) interface{} {
	data := &struct {
		Data   string `json:"data"`
		IV     string `json:"iv"`
		OpenID string `json:"open_id"`
	}{}
	if err := ctx.ReqJson(data); err != nil {
		return err
	}

	if data.Data == "" || data.IV == "" || data.OpenID == "" {
		return rsp.ParamIllegal
	}

	phone, err := auth.DecryptUPhoneData(ctx, data.Data, data.IV, data.OpenID, resource.C.SmartApp)
	if err != nil {
		return err
	}

	return phone
}
