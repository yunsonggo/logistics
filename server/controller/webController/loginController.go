package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
	"strconv"
)

var us = server.NewUserServer()
var ss = server.NewSmsServer()

func UserLogin(ctx *gin.Context) {
	var loginPost param.LoginParam
	err := ctx.ShouldBind(&loginPost)
	if err != nil {
		common.Failed(ctx, "获取手机号或验证码失败")
		return
	}
	// 验证 手机 手机短信
	codeId, cErr := ss.CheckSms(loginPost.Phone, loginPost.Code)
	if codeId == 0 || cErr != nil {
		common.Failed(ctx, "手机号或者验证码错误")
		return
	}
	// 登录
	//查询用户是否已经注册
	resData, rErr := us.GetUserByPhone(loginPost.Phone)
	if rErr != nil {
		common.Failed(ctx, "验证手机号是否已注册失败")
		return
	}
	if resData.Id == 0 {
		//注册新用户
		userId, sErr := us.SignUpUserBySms(loginPost.Phone)
		if sErr != nil {
			common.Failed(ctx, "注册新用户失败")
			return
		}
		if userId == 0 {
			common.Failed(ctx, "新用户登录失败")
			return
		}
		// 设置session
		stringId := strconv.FormatInt(userId, 10)
		ssErr := common.SetSess(ctx, "user_"+stringId, userId)
		if ssErr != nil {
			common.Failed(ctx, "设置session错误,用户登录失败")
			return
		}
		ctx.SetCookie("cookie_user", stringId, 4300, "/", "localhost", true, true)
		tokenString,tokenErr := common.ReleaseToken(userId)
		if tokenErr != nil {
			common.Failed(ctx,"TOKEN错误,登录失败")
			return
		}
		common.Success(ctx, tokenString)
		_ = ss.RemoveSms(codeId)
		return
	}
	// 已注册用户登录
	if resData.IsDel == 0 {
		common.Failed(ctx, "用户已注销,请联系管理员")
		return
	}
	// 设置session
	stringResId := strconv.FormatInt(resData.Id, 10)
	ssErr := common.SetSess(ctx, "user_"+stringResId, resData.Id)
	if ssErr != nil {
		common.Failed(ctx, "设置session错误,用户登录失败")
		return
	}
	ctx.SetCookie("cookie_user", stringResId, 4300, "/", "localhost", true, true)
	tokenString,tokenErr := common.ReleaseToken(resData.Id)
	if tokenErr != nil {
		common.Failed(ctx,"TOKEN错误,登录失败")
		return
	}
	common.Success(ctx, tokenString)
	_ = ss.RemoveSms(codeId)
	return
}

// 根据	id 获取用户信息
func PostUserById(ctx *gin.Context) {
	webUser ,_ := ctx.Get("webUser")
	if webUser == nil {
		common.Failed(ctx,"需要登陆")
		return
	}
	common.Success(ctx,webUser)
}
