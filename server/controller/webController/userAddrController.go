package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server"
)

var userAddrServer = server.NewUserAddrServer()

// 获取用户地址
func PostUserAddr(ctx *gin.Context) {
	webUser , _ := ctx.Get("webUser")
	if webUser == nil {
		common.Failed(ctx,"需要登陆")
		return
	}
	user := webUser.(model.User)
	res,resErr := userAddrServer.GetUserAddrById(user.Id)
	if resErr != nil {
		common.Failed(ctx,"获取用户地址失败")
		return
	}
	common.Success(ctx,res)
	return
}

// 添加 用户 收货地址
func PostAddUserAddr(ctx *gin.Context) {
	addressInfo,_ := ctx.Get("addressInfo")
	info := addressInfo.(model.UserAddress)
	if info.UserId > 0 {
		// 设为默认收货地址
		if info.Position == 0 {
			clearErr := userAddrServer.ClearAddrPosition(info.UserId)
			if clearErr != nil {
				common.Failed(ctx,"清除原默认地址失败")
				return
			}
			addErr := userAddrServer.AddUserAddrOne(&info)
			if addErr != nil {
				common.Failed(ctx,"添加新默认地址失败")
				return
			}
			common.Success(ctx,"添加新默认地址成功")
			return
		} else {
			addErr := userAddrServer.AddUserAddrOne(&info)
			if addErr != nil {
				common.Failed(ctx,"添加地址失败")
				return
			}
			common.Success(ctx,"添加地址成功")
			return
		}
	} else {
		common.Failed(ctx,"添加地址失败:请登陆")
		return
	}
}
// 编辑 用户 收货地址
func PostEditUserAddr(ctx *gin.Context) {
	addressInfo,_ := ctx.Get("addressInfo")
	info := addressInfo.(model.UserAddress)
	if info.UserId > 0 {
		// 设为默认收货地址
		if info.Position == 0 {
			clearErr := userAddrServer.ClearAddrPosition(info.UserId)
			if clearErr != nil {
				common.Failed(ctx,"清除原默认地址失败")
				return
			}
			editErr := userAddrServer.EditOneAddrById(info.UserId,&info)
			if editErr != nil {
				common.Failed(ctx,"编辑地址失败")
				return
			}
			common.Success(ctx,"编辑地址成功")
			return
		} else {
			editErr := userAddrServer.EditOneAddrById(info.UserId,&info)
			if editErr != nil {
				common.Failed(ctx,"编辑地址失败")
				return
			}
			common.Success(ctx,"编辑地址成功")
			return
		}
	} else {
		common.Failed(ctx,"修改地址失败:请登陆")
		return
	}
}
// 删除 用户收货地址
func PostDelUserAddr(ctx *gin.Context) {
	addressInfo,_ := ctx.Get("addressInfo")
	info := addressInfo.(model.UserAddress)
	if info.UserId > 0 {
		delErr := userAddrServer.RemoveOneAddrById(info.UserId,&info)
		if delErr != nil {
			common.Failed(ctx,"删除地址失败")
			return
		}
		common.Success(ctx,"删除地址成功")
		return
	} else {
		common.Failed(ctx,"删除地址失败:请登陆")
		return
	}
}
