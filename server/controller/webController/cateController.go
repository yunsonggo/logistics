package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/server"
)

var cs = server.NewCatesServer()

func GetProCate(ctx *gin.Context) {
	catesData,err := cs.GetCatesAny("state",1)
	if err != nil {
		common.Failed(ctx,"获取商品类别失败")
		return
	}
	common.Success(ctx,catesData)
	return
}
