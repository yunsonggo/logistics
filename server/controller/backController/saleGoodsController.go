package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server/backServer"
)

var sgs = backServer.NewSaleGoodsServer()

// 根据订单ID 查询
func PostGoodsByOrderId(ctx *gin.Context) {
	orderInfo,hasData := ctx.Get("orderInfo")
	if !hasData {
		common.Failed(ctx,"获取订单ID失败")
	}
	orderData := orderInfo.(model.Orders)
	id := orderData.Id
	res,resErr := sgs.GetSaleGoodsByOrderId(id)
	if resErr != nil {
		common.Failed(ctx,"查询已售商品失败")
		return
	}
	common.Success(ctx,res)
	return
}
