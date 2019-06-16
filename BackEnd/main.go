package main

import (
	"path/filepath"
	"runtime"

	"BackEnd/controllers/auth"
	"BackEnd/controllers/debug"
	"BackEnd/controllers/pay"
	"BackEnd/lib/conf"
	"BackEnd/lib/httpserver"
	"BackEnd/lib/log"
	"BackEnd/resource"
)

func main() {
	conf.Init(root())

	if err := resource.Init(map[string]interface{}{
		"server.json":    resource.C.Server,
		"log.json":       resource.C.Log,
		"pay.json":       resource.C.Pay,
		"smart_app.json": resource.C.SmartApp,
	}); err != nil {
		panic(err)
	}

	if err := log.Init(resource.C.Log); err != nil {
		panic(err)
	}

	var err error
	if resource.C.SelfRsaPrivKey, err = conf.LoadFile("self_rsa_private_key.pem"); err != nil {
		panic(err)
	}

	if resource.C.PlatformRsaPubKey, err = conf.LoadFile("platform_rsa_public_key.pem"); err != nil {
		panic(err)
	}

	// 启动web服务
	server := &httpserver.Server{
		Routers: map[string]httpserver.HandleFunc{
			"/pay/gen":           pay.Gen,
			"/pay/callback/succ": pay.SuccCallBack,
			"/pay/status":        pay.Status,
			"/pay/refund":        pay.Refund,

			"/auth/login":    auth.Login,
			"/auth/userinfo": auth.GetUserInfo,
			"/auth/phone":    auth.GetPhone,

			"/debug": debug.Debug,
		},
	}
	err = server.RunServer(resource.C.Server)
	panic(err)
}

func root() string {
	_, file, _, _ := runtime.Caller(1)

	return filepath.Dir(file)
}
