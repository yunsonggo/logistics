package middleware

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
	"strconv"
)

func WebAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var us = server.NewUserServer()
		var jwtParam param.JwtParam
		err := ctx.ShouldBind(&jwtParam)
		if err != nil {
			common.Failed(ctx,"获取用户信息参数失败")
			ctx.Abort()
			return
		}
		tokenString := jwtParam.JwtUserId
		token,claims,parseErr := common.ParseToken(tokenString)
		if parseErr != nil || !token.Valid {
			common.Failed(ctx,"需要登陆")
			ctx.Abort()
			return
		}
		userId := claims.UserId
		sessId := strconv.FormatInt(userId, 10)
		id := common.GetSess(ctx,"user_" + sessId)
		res,resErr := us.GetUserById(id)
		if resErr != nil || res.Id == 0 || res.IsDel == 0 {
			common.Failed(ctx,"需要注册")
			ctx.Abort()
			return
		}
		res.PassWord = ""
		ctx.Set("webUser",res)
		if jwtParam.AddressInfo.UserId > 0 {
			ctx.Set("addressInfo",jwtParam.AddressInfo)
		}
		if jwtParam.OrderList != nil {
			ctx.Set("orderListData",jwtParam.OrderList)
		}
		ctx.Next()
	}
}
