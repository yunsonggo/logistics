package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
	"orderserve/tools"
	"regexp"
	"strings"
)
// 根据商品关键词 搜索 商家
func PostSearch(ctx *gin.Context) {
	var searchParam param.SearchParam
	err := ctx.ShouldBind(&searchParam)
	if err != nil {
		common.Failed(ctx,"获取搜索商品参数失败")
		return
	}
	// 去除空格
	trimParam := strings.TrimSpace(searchParam.KeyWord)
	// 关键字切片
	paramSlice := regexp.MustCompile(" +").Split(trimParam,-1)
	// 根据关键字搜索商品
	if paramSlice[0] == "" {
		common.Failed(ctx,"没有搜索关键词")
		return
	}
	res,resErr := server.SearchGoodBykey("name",paramSlice)
	suppRes,suppResErr := server.SearchSuppByName("name",paramSlice)
	if suppResErr != nil || resErr != nil {
		common.Failed(ctx,"搜索商家或商品失败")
		return
	}
	// 根据商品 提取 商家ID
	var idArray []int64
	for _,val := range res {
		idArray = append(idArray, val.ShopId)
	}
	for _,val := range suppRes {
		idArray = append(idArray,val.Id)
	}
	ids := tools.SortSliceInt(idArray)
	shopIds := tools.RemoveDup(ids)
	// 搜索商家
	shops,shopErr := server.SearchSuppByKey("id",shopIds)
	if shopErr != nil {
		common.Failed(ctx,"查询商家失败")
		return
	}
	common.Success(ctx,shops)
	return
}

