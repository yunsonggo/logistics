package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/server/backServer"
)

var ubs = backServer.NewUserBackServer()
// 获取所有用户数据
func PostUserAll(ctx *gin.Context) {
	res,err := ubs.GetUserAll()
	if err != nil {
		common.Failed(ctx,"获取用户数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
