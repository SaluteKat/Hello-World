package pay

import (
	paydata "BackEnd/data/pay"
)

func SuccCallBack(params *paydata.SuccCallBackParam) error {
	return paydata.DefaulPay.PaySucc(params)
}
