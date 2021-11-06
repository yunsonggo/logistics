package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
)
var evaServe = server.NewEvaServe()
// 根据商品ID 获取 评论
func PostEvaByGoodId(ctx *gin.Context) {
	 //server.InitEva()
	var evaParam param.EvaParam
	err := ctx.ShouldBind(&evaParam)
	if err != nil {
		common.Failed(ctx,"获取评论参数失败")
		return
	}
	start := evaParam.Page * evaParam.Count - (evaParam.Count -1)
	res,resErr := evaServe.GetEvaByKey(evaParam.IdName,evaParam.IdVal,evaParam.Count,start)
	if resErr != nil {
		common.Failed(ctx,"获取评论数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
