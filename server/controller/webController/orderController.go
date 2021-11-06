package webController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server"
	"strconv"
	"strings"
	"time"
)

var sgs = server.NewSaleGoodsServer()
var ods = server.NewOrderServer()
var pays = server.NewPayServer()
var supp = server.NewSupplierServer()
// 生成订单 入库 并 调用 支付
func PostOrder(ctx *gin.Context) {
	// 用户数据
	webUserData,_ := ctx.Get("webUser")
	userData := webUserData.(model.User)
	orderListData,exists := ctx.Get("orderListData")
	if !exists {
		common.Failed(ctx,"解析订单数据失败")
		return
	}
	// 订单数据
	var orderData model.Orders
	// 已售商品数据
	var saleGoods []model.SaleGoods
	// 给订单 赋值 用户名 电话
	orderData.UserName = userData.UserName
	orderData.UserPhone = userData.Phone
	// 整理 订单 数据结构
	for key,value := range orderListData.(map[string]interface{}) {
		switch key {
		case "userid" :
			orderData.UserId,_ = strconv.ParseInt(value.(string),10,64)
		case "selectPay" :
			orderData.SelectPay = value.(string)
		case "totalPrice" :
			orderData.TotalPrice = value.(float64)
		case "deliveryPrice" :
			orderData.DeliveryPrice = value.(float64)
		case "addressInfo" :
			var addrString string
			var addrBottom string
			var addrName string
			for addrKey,addrValue := range value.(map[string]interface{}) {
				switch addrKey {
				case "id" :
					orderData.AddressId = int64(addrValue.(float64))
				case "lat" :
					orderData.Lat = addrValue.(float64)
				case "lng" :
					orderData.Lng = addrValue.(float64)
				case "mobile" :
					orderData.LiaisonPhone = addrValue.(string)
				case "user_name" :
					orderData.Liaison = addrValue.(string)
				case "user_gender" :
					orderData.Gender = addrValue.(string)
				case "address" :
					addrString = addrValue.(string)
				case "bottom" :
					addrBottom = addrValue.(string)
				case "name" :
					addrName = addrValue.(string)
				}
			}
			orderData.Address = addrString + addrBottom + "_" + addrName + "_附近"
		case "orderInfo" :
			for orderInfoKey,orderInfoValue := range value.(map[string]interface{}) {
				switch orderInfoKey {
				case "shopInfo" :
					for shopInfoKey,shopInfoValue := range orderInfoValue.(map[string]interface{}) {
						switch shopInfoKey {
						case "id" :
							orderData.ShopId = int64(shopInfoValue.(float64))
						case "name" :
							orderData.ShopName = shopInfoValue.(string)
						case "mobile" :
							orderData.ShopPhone = shopInfoValue.(string)
						case "is_union" :
							orderData.IsUnion = int(shopInfoValue.(float64))
						}
					}
				case "selectedGoods" :
					for _,arrayValue := range orderInfoValue.([]interface{}) {
						var saleGood model.SaleGoods
						for k,v := range arrayValue.(map[string]interface{}) {

								switch k {
								case "id" :
									saleGood.GoodId = int64(v.(float64))
								case "name" :
									saleGood.GoodName = v.(string)
								case "count" :
									saleGood.GoodCount = int(v.(float64))
								case "cate_id" :
									saleGood.CateId = int64(v.(float64))
								case "price" :
									saleGood.GoodPrice = v.(float64)
								case "shop_id" :
									saleGood.ShopId = int64(v.(float64))
								case "icon" :
									saleGood.Icon = v.(string)
								}
						}
						saleGoods = append(saleGoods,saleGood)
					}
				}
			}
		case "remarkInfo" :
			for rk,rv := range value.(map[string]interface{}) {
				switch rk {
				case "invoice" :
					orderData.Invoice = rv.(string)
				case "remark" :
					orderData.Remark = rv.(string)
				case "ware" :
					orderData.Exchange = rv.(string)
				}
			}
		case "distanceMetre" :
			orderData.Metre = value.(float64)
		case "goodsPrice" :
			orderData.GoodsPrice = value.(float64)
		case "orderDelTime" :
			orderData.DeliveryDateTime = value.(string)
		}
	}
	goodsInfo := ""
	for key,value := range saleGoods {
		value.UserId = orderData.UserId
		value.ShopName = orderData.ShopName
		value.Liaison = orderData.Liaison
		value.LiaisonPhone = orderData.LiaisonPhone
		value.Exchange = orderData.Exchange
		value.Invoice = orderData.Invoice
		goodsInfo += "[商品 " + strconv.Itoa(key + 1) + "]:_id:" + strconv.FormatInt(value.GoodId,10) + "_名称:" +
			value.GoodName + "_单价:" + strconv.FormatFloat(value.GoodPrice,'f',-1,64) + "_数量:" +
			strconv.Itoa(value.GoodCount) + "|"
	}
	orderData.GoodsInfo = goodsInfo
	// 生成订单识别码
	// 六位随机数
	c := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	orderCount,countErr := ods.GetOrderCount()
	if countErr != nil {
		common.Failed(ctx,"获取订单数量错误")
		return
	}
	orderCount = orderCount + 1
	code := strconv.FormatInt(time.Now().Unix(), 10) + c + strconv.FormatInt(orderCount,10)
	orderData.OrderCode = code
	// 订单入库
	err := ods.AddOrderOne(&orderData)
	if err != nil {
		common.Failed(ctx,"订单入库失败")
		return
	}
	// 通过订单识别码 获取新入库订单
	resOrder,resOrderErr := ods.GetOrderByOrderCode(code)
	if resOrderErr != nil {
		common.Failed(ctx,"获取订单ID失败")
		return
	}
	// 订单ID 赋值给 已售商品 并且入库
	for key,value := range saleGoods {
		value.SaleCode = strconv.FormatInt(time.Now().Unix(), 10) + c + strconv.Itoa(key)
		value.OrderId = resOrder.Id
		value.ShopName = resOrder.ShopName
		value.Liaison = resOrder.Liaison
		value.LiaisonPhone = resOrder.LiaisonPhone
		addSaleGoodErr := sgs.AddSaleGoodOne(&value)
		if addSaleGoodErr != nil {
			common.Failed(ctx,"已售商品入库失败")
			return
		}
		resGood,resGoodErr := sgs.GetSaleGoodByCode(value.SaleCode)
		if resGoodErr != nil {
			common.Failed(ctx,"获取商品ID失败")
			return
		}
		resOrder.SaleGoodsId += strconv.FormatInt(resGood.Id,10) + ","
	}
	updateOrderErr := ods.UpdateOrderOne(resOrder.Id,resOrder)
	if updateOrderErr != nil {
		common.Failed(ctx,"更新订单商品ID失败")
		return
	}
	// 调用在线支付
	if resOrder.SelectPay == "0" {
		payResErr := pays.PayByWechat()
		if payResErr != nil {
			common.Failed(ctx,"支付失败")
			return
		}
		//改变 订单状态
		resOrder.State = 1
	} else if resOrder.SelectPay == "1" {
		payResErr := pays.PayByAliPay()
		if payResErr != nil {
			common.Failed(ctx,"支付失败")
			return
		}
		//改变 订单状态
		resOrder.State = 1
	}
	// 更新订单状态
	upStateErr := ods.UpdateOrderOne(resOrder.Id,resOrder)
	if upStateErr != nil {
		common.Failed(ctx,"更新订单状态失败")
		return
	}
	common.Success(ctx,"支付成功")
	return
}
// 获取用户所有订单
func PostOrderAll(ctx *gin.Context) {
	// 用户数据
	webUserData,_ := ctx.Get("webUser")
	userData := webUserData.(model.User)
	// 所有订单数据
	res,resErr := ods.GetOrderAllByUserId(userData.Id)
	if resErr != nil {
		common.Failed(ctx,"获取订单数据失败")
		return
	}
	// 添加 商品 商家 数据
	for key,value := range res {
		// 商品数据
		goodIds := strings.Split(value.SaleGoodsId,",")
		for _,v := range goodIds {
			id,_ := strconv.ParseInt(v,10,64)
			if id > 0 {
				resGood,resGoodErr := sgs.GetSaleGoodById(id)
				if resGoodErr != nil {
					common.Failed(ctx,"查询订单商品失败")
					return
				}
				res[key].SaleGoodsList = append(res[key].SaleGoodsList,*resGood)
			}
		}
		shopId := value.ShopId
		resShop,resShopErr := supp.GetSupplier(shopId)
		if resShopErr != nil {
			common.Failed(ctx,"获取商家数据失败")
			return
		}
		res[key].ShopInfo = *resShop
	}
	common.Success(ctx,res)
}