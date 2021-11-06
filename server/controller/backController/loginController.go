package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param/backParam"
	"orderserve/server/backServer"
	"orderserve/tools"
	"strconv"
)

type ResManagerData struct {
	ResToken string `json:"res_token"`
	ResUserId int64 `json:"res_user_id"`
}

// 实例化管理员 服务
var ms = backServer.NewManagerServer()
// 注册用户
func PostSignUp(ctx *gin.Context) {
	var managerParam backParam.ManagerLoginParam
	err := ctx.ShouldBind(&managerParam)
	if err != nil {
		common.Failed(ctx,"获取注册参数失败")
		return
	}
	userName := managerParam.UserName
	passWord := tools.EncoderSha256(managerParam.UserPassword)
	// 判断用户是否已经存在
	res,resErr := ms.GetManagerByName(userName)
	if resErr != nil {
		common.Failed(ctx,"查询用户是否存在失败")
		return
	}
	if res.Id > 0 {
		common.Failed(ctx,"用户已经存在")
		return
	}
	// 添加新用户
	addErr := ms.AddManager(userName,passWord)
	if addErr != nil {
		common.Failed(ctx,"注册新用户失败")
		return
	}
	// 添加新用户成功
	common.Success(ctx,"ok")
	return
}
// 验证用户登录
func PostLogin(ctx *gin.Context) {
	var managerParam backParam.ManagerLoginParam
	err := ctx.ShouldBind(&managerParam)
	if err != nil {
		common.Failed(ctx,"获取登录参数失败")
		return
	}
	userName := managerParam.UserName
	passWord := tools.EncoderSha256(managerParam.UserPassword)
	//验证 用户
	res,resErr := ms.GetManagerByPass(userName,passWord)
	if resErr != nil || res.Id == 0 {
		common.Failed(ctx,"验证用户失败")
		return
	}
	// 生成JWT TOKEN
	managerToken,jwtErr := common.ReleaseToken(res.Id)
	if jwtErr != nil {
		common.Failed(ctx,"生成用户TOKEN失败,请重新登录")
		return
	}
	// 设置session
	stringId := strconv.FormatInt(res.Id,10)
	ssErr := common.SetSess(ctx,"manager_" + stringId,res.Id)
	if ssErr != nil {
		common.Failed(ctx,"设置session错误,请重新登录")
		return
	}
	var data ResManagerData
	data.ResToken = managerToken
	data.ResUserId = res.Id
	// 返回 	JWT TOKEN USER_ID
	common.Success(ctx,data)
}
