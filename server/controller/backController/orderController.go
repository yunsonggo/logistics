package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/server/backServer"
)

var ors = backServer.NewOrderServer()

// 获取所有订单数据
func PostOrderAll(ctx *gin.Context) {
	res,err := ors.GetOrderAll()
	if err != nil {
		common.Failed(ctx,"获取订单数据错误")
		return
	}
	common.Success(ctx,res)
	return
}
