package server

type PayInterface interface {
	// 微信支付
	PayByWechat() (err error)
	//支付宝支付
	PayByAliPay() (err error)
}

type PayServer struct {}

func NewPayServer() PayInterface {
	return &PayServer{}
}

// 微信支付
func (pays *PayServer) PayByWechat() (err error) {
	return nil
}
//支付宝支付
func (pays *PayServer) PayByAliPay() (err error) {
	return nil
}