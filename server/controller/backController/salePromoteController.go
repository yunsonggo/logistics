package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server/backServer"
)

var sps = backServer.NewSalePromoteServer()

// 获取所有活动数据
func PostSalePromoteList(ctx *gin.Context) {
	res,resErr := sps.GetSalePromoteAll()
	if resErr != nil {
		common.Failed(ctx,"获取活动数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 修改活动数据
func PostSalePromoteEdit(ctx *gin.Context) {
	promoteData,hasData := ctx.Get("promoteData")
	if !hasData {
		common.Failed(ctx,"获取活动数据失败")
		return
	}
	promoteInfo := promoteData.(model.Sale)
	id := promoteInfo.Id
	err := sps.EditSalePromote(id,promoteInfo)
	if err != nil {
		common.Failed(ctx,"更新活动数据失败")
	}
	common.Success(ctx,promoteInfo)
	return
}
// 新增活动数据
func PostSalePromoteAdd(ctx *gin.Context) {
	promoteData,hasData := ctx.Get("promoteData")
	if !hasData {
		common.Failed(ctx,"获取活动数据失败")
		return
	}
	promoteInfo := promoteData.(model.Sale)
	err := sps.AddSalePromoteOne(promoteInfo)
	if err != nil {
		common.Failed(ctx,"新增活动数据失败")
	}
	common.Success(ctx,promoteInfo)
	return
}