package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/server"
)

var gs = server.NewGoodServer()

// 首页展示 轮播商品
func GetProShopping(ctx *gin.Context) {
	mainGoodsList, err := gs.GetGoodsAny("is_main",1)
	if err != nil {
		common.Failed(ctx,"获取商品数据失败")
		return
	}
	common.Success(ctx,mainGoodsList)
}
