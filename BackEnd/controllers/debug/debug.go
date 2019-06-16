package debug

import (
	"BackEnd/data/auth"
	"BackEnd/data/pay"
	"BackEnd/lib/httpserver"
)

func Debug(ctx *httpserver.Context) interface{} {
	user, _ := auth.DefaultAuth.GetAllUser()
	order, _ := pay.DefaulPay.GetAllOrders()

	return map[string]interface{}{
		"user":  user,
		"order": order,
	}
}
