package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
)

var smsServer = server.NewSmsServer()

func SendCode(ctx *gin.Context) {
	var smsPost param.SmsParam
	err := ctx.ShouldBind(&smsPost)
	if err != nil {
		common.Failed(ctx, "获取验证码参数失败")
		return
	}
	codeStr,resErr := smsServer.SendSms(smsPost.Phone, smsPost.TplId, smsPost.Key)
	if resErr != nil {
		common.Failed(ctx, "发送验证码失败")
		return
	}
	common.Success(ctx, codeStr)
	return
}
