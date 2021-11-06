package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server"
)

var saleSer = server.NewSaleServer()
var giveSer = server.NewGiveServer()

func GetTab(ctx *gin.Context) {
	menu := new(model.Menu)
	menuList := []map[string]string {
		{"name":"商家排序","value":"menu_sort","icon":"caret-down"},
		{"name":"商品排序","value":"menu_good","icon":"caret-down"},
		{"name":"签约商家","value":"is_union","icon":""},
		{"name":"筛选","value":"menu_good","icon":"filter"},
	}
	sort := []map[string]string{
		{"name":"综合排序","value":"is_main","icon":""},
		{"name":"距离最近","value":"is_around","icon":""},
		{"name":"好评优先","value":"rating","icon":""},
		{"name":"人气最高","value":"is_hot","icon":""},
		{"name":"最新商家","value":"is_new","icon":""},
	}
	menuGoods := []map[string]string{
		{"name":"商品销量","value":"sell_count","icon":"sketch"},
		{"name":"商品评分","value":"score","icon":"grin-beam"},
		{"name":"推荐商品","value":"is_top","icon":"pagelines"},
		{"name":"热卖商品","value":"is_hot","icon":"hotjar"},
		{"name":"商品价格","value":"price","icon":"money-bill-alt"},
		{"name":"商品运费","value":"over_price","icon":"truck"},
		{"name":"商家包运","value":"is_free","icon":"luggage-cart"},
	}
	menuPrice := []map[string]string {
		{"name":"price0","minValue":"","maxValue":"20","select":""},
		{"name":"price1","minValue":"20","maxValue":"40","select":""},
		{"name":"price2","minValue":"40","maxValue":"60","select":""},
		{"name":"price3","minValue":"60","maxValue":"80","select":""},
		{"name":"price4","minValue":"80","maxValue":"100","select":""},
		{"name":"price5","minValue":"100","maxValue":"","select":""},
	}
	menu.MenuList = menuList
	menu.Sort = sort
	menu.MenuGive,_ = giveSer.GetGiveAll()
	menu.MenuSale,_ = saleSer.GetSaleAll()
	menu.MenuGoods = menuGoods
	menu.MenuPrice = menuPrice
	if len(menu.MenuGive) == 0 || len(menu.MenuSale) == 0 {
		common.Failed(ctx,"获取菜单数据失败")
		return
	}
	common.Success(ctx,menu)
	return
}

