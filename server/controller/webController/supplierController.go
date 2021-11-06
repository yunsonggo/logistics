package webController

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param"
	"orderserve/server"
	"orderserve/tools"
	"regexp"
	"strconv"
)

var suppServer = server.NewSupplierServer()
var giveServer = server.NewGiveServer()
var saleServer = server.NewSaleServer()
var goodServer = server.NewGoodServer()

// 首页展示 商家 条件 查找 筛选
func PostProSupplier(ctx *gin.Context) {
	// 获取参数
	var suppParam param.SupplierPatam
	err := ctx.ShouldBind(&suppParam)
	if err != nil {
		//fmt.Println(err)
		common.Failed(ctx,"获取商家参数失败")
		return
	}
	//fmt.Println(suppParam)
	//如果KEY == good 表示 按商品字段 排序
	if suppParam.Keys == "good" {
		var idArray []int64
		// 字段 = 1
		if suppParam.Condition == "is_top" || suppParam.Condition == "is_hot" {
			var data []string
			data = append(data,"1")
			res,resErr := server.SearchGoodBykey(suppParam.Condition,data)
			if resErr != nil {
				common.Failed(ctx,"查询商品失败")
				return
			}
			for _,val := range res {
				idArray = append(idArray, val.ShopId)
			}
		} else if suppParam.Condition == "sell_count" || suppParam.Condition == "score" {
		// 降序排列
			count := suppParam.Count
			start := suppParam.Page * suppParam.Count - (suppParam.Count - 1)
			res,resErr := server.SearchGoodsOrder(suppParam.Condition,"asc",count,start)
			if resErr != nil {
				common.Failed(ctx,"查询商品失败")
				return
			}
			for _,val := range res {
				idArray = append(idArray, val.ShopId)
			}
		} else {
		// 升序排列
			count := suppParam.Count
			start := suppParam.Page * suppParam.Count - (suppParam.Count - 1)
			res,resErr := server.SearchGoodsOrder(suppParam.Condition,"desc",count,start)
			if resErr != nil {
				common.Failed(ctx,"查询商品失败")
				return
			}
			for _,val := range res {
				idArray = append(idArray, val.ShopId)
			}
		}
		// 去重
		shopIds := tools.RemoveIntByMap(idArray)
		// 搜索商家
		shops,shopErr := server.SearchSuppByKey("id",shopIds)
		if shopErr != nil {
			common.Failed(ctx,"查询商家失败")
			return
		}
		common.Success(ctx,shops)
		return
	} else if suppParam.Keys == "screen" {
// 多条件查询商品
		if suppParam.Condition == "" {
			common.Failed(ctx,"未选择筛选条件")
			return
		}
		conditionArr := regexp.MustCompile(`,`).Split(suppParam.Condition,-1)
		var giveIds []interface{}
		var saleIds []interface{}
		var minString string
		var maxString string
		for _,val := range conditionArr {
			if val != "" {
				if val[:1] == "g" {
					valInt,_ := strconv.ParseInt(val[1:],10,64)
					giveIds = append(giveIds,valInt)
				} else if val[:1] == "s" {
					valInt,_ := strconv.ParseInt(val[1:],10,64)
					saleIds = append(saleIds,valInt)
				} else if val[:3] == "min" {
					minString = val[3:]
				} else if val[:3] == "max" {
					maxString = val[3:]
				}
			}
		}
		minf,_ := strconv.ParseFloat(minString,32)
		maxf,_ := strconv.ParseFloat(maxString,32)
		mif := float32(minf)
		maf := float32(maxf)
		resGoodsData,resGoodsErr := server.SearchGoodsByArray(giveIds,saleIds,mif,maf,suppParam.Count,suppParam.Page)
		if resGoodsErr != nil {
			common.Failed(ctx,"查询商品失败")
			return
		}
		var idArray []int64
		for _,val := range resGoodsData {
			idArray = append(idArray, val.ShopId)
		}
		// 去重
		shopIds := tools.RemoveIntByMap(idArray)
		// 搜索商家
		shops,shopErr := server.SearchSuppByKey("id",shopIds)
		if shopErr != nil {
			common.Failed(ctx,"查询商家失败")
			return
		}
		common.Success(ctx,shops)
		return
	}
	// 根据参数 获取商家
	resData,resErr := suppServer.GetSupplierByKeys(&suppParam)
	if resErr != nil {
		common.Failed(ctx,"获取商家数据失败")
		return
	}
	common.Success(ctx,resData)
	return
}
// 商家详情
func GetSupplierInfo(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		common.Failed(ctx,"参数错误")
		return
	}
	idInt,_ :=strconv.ParseInt(id,10,64)
	res,err := suppServer.GetSupplier(idInt)
	if err != nil {
		common.Failed(ctx,"获取商家信息失败")
		return
	}
	common.Success(ctx,res)
}
// 商家参与服务
func PostSupplierGive(ctx *gin.Context) {
	var giveParam param.GiveParam
	err := ctx.ShouldBind(&giveParam)
	if err != nil {
		common.Failed(ctx,"获取服务参数失败")
		return
	}
	giveStr := regexp.MustCompile(`,`).Split(giveParam.GiveIds,-1)
	var giveIds []interface{}
	for _,val := range giveStr {
		giveIds = append(giveIds,val)
	}
	giveRes,giveResErr := giveServer.GetGiveByIdArray(giveIds)
	if giveResErr != nil {
		common.Failed(ctx,"获取商家提供服务失败")
		return
	}
	common.Success(ctx,giveRes)
	return
}
// 商家参与活动
func PostSupplierSale(ctx *gin.Context) {
	var saleParam param.SaleParam
	err := ctx.ShouldBind(&saleParam)
	if err != nil {
		common.Failed(ctx,"获取活动参数失败")
		return
	}
	saleStr := regexp.MustCompile(`,`).Split(saleParam.SaleIds,-1)
	var saleIds []interface{}
	for _,val := range saleStr {
		saleIds = append(saleIds,val)
	}
	saleRes,saleResErr := saleServer.GetSaleByIdArray(saleIds)
	if saleResErr != nil {
		common.Failed(ctx,"获取商家提供服务失败")
		return
	}
	common.Success(ctx,saleRes)
	return
}
// 商家所有商品
func PostStateGoodsAll(ctx *gin.Context) {
	var shopIdParam param.AllGoodsParam
	err := ctx.ShouldBind(&shopIdParam)
	if err != nil {
		common.Failed(ctx,"获取商家ID失败")
		return
	}
	res,resErr := goodServer.GetGoodsAny("shop_id",shopIdParam.ShopId)
	if resErr != nil {
		common.Failed(ctx,"获取商品失败")
		return
	}
	common.Success(ctx,res)
}
// 商家商品分类
func PostSupplierCatesAll(ctx *gin.Context) {
	var shopIdParam param.AllGoodsParam
	err := ctx.ShouldBind(&shopIdParam)
	if err != nil {
		common.Failed(ctx,"获取商家ID失败")
		return
	}
	var scs = server.NewShopCateServer()
	res,resErr := scs.GetShopCateByShop(shopIdParam.ShopId)
	if resErr != nil {
		common.Failed(ctx,"获取商家商品分类失败")
		return
	}
	common.Success(ctx,res)
}
// 根据分类查找商家
func GetSupplierByCate(ctx *gin.Context) {
	cateStringId := ctx.Query("cate_id")
	cateId,_ :=strconv.ParseInt(cateStringId,10,64)
	res,resErr := suppServer.GetSupplierByCateId(cateId)
	if resErr != nil {
		common.Failed(ctx,"查询商家失败")
		return
	}
	common.Success(ctx,res)
	return
}
