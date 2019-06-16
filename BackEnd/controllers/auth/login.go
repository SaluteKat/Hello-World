package auth

import (
	"BackEnd/lib/httpserver"
	"BackEnd/lib/rsp"
	"BackEnd/model/auth"
	"BackEnd/resource"
)

// Login 调用swan.Login 后第三方开发者前端回调第三方服务端
func Login(ctx *httpserver.Context) interface{} {
	code := ctx.QueryString("code")
	if code == "" {
		return rsp.ParamIllegal
	}

	openID, err := auth.Login(ctx, code, resource.C.SmartApp)
	if err != nil {
		return err
	}

	return map[string]interface{}{
		"open_id": openID,
	}
}
