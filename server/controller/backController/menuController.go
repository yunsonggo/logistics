package backController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/model/backendModel"
	"orderserve/server/backServer"
)

var mus = backServer.NewMenuServer()
// 获取所有菜单
func PostMenuAll(ctx *gin.Context) {
	res,resErr := mus.GetMenuAll()
	if resErr != nil {
		common.Failed(ctx,"获取菜单数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 添加菜单
func PostMenuAdd(ctx *gin.Context) {
	addMenuData,_ := ctx.Get("addMenuData")
	menuData := addMenuData.(backendModel.BackendMenu)
	if menuData.MenuTitle != "" {
		// 入库
		err := mus.AddMenuOne(&menuData)
		if err != nil {
			common.Failed(ctx,"添加菜单入数据库错误")
			return
		}
		common.Success(ctx,menuData)
		return
	}
	common.Failed(ctx,"获取菜单数据失败")
	return
}
// 更新菜单
func PostMenuEdit(ctx *gin.Context) {
	addMenuData,_ := ctx.Get("addMenuData")
	menuData := addMenuData.(backendModel.BackendMenu)
	if menuData.MenuTitle != "" {
		// 入库
		err := mus.EditMenuOne(&menuData)
		if err != nil {
			common.Failed(ctx,"编辑菜单入数据库错误")
			return
		}
		common.Success(ctx,menuData)
		return
	}
	common.Failed(ctx,"获取菜单数据失败")
	return
}