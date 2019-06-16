package pay

import (
	"BackEnd/lib/httpserver"
	paymodel "BackEnd/model/pay"
	"BackEnd/resource"
)

// Status 查询订单
func Status(ctx *httpserver.Context) interface{} {
	// 智能小程序现有的封装没办法得到订单创建成功后 orderID-tpOrderID 的映射
	orderID := ctx.QueryString("tp_order_id")

	data, err := paymodel.Status(ctx, orderID, resource.C.Pay, resource.C.SelfRsaPrivKey)
	if err != nil {
		ctx.Warning(err)
		return err
	}

	return data
}
