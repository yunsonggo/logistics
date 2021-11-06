package server

import (
	"fmt"
	"orderserve/dal"
	"orderserve/model"
)

/**
 * 向goods表中插入初始测试数据
 */
func InitGoodsData() {
	goods := []model.Goods{
		{
			Name: "测试商品六",
			Description: "测试商品六描述",
			Content: "测试商品六详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 1,
			CateId: 1,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品7",
			Description: "测试商品7描述",
			Content: "测试商品7详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 2,
			CateId: 2,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品8",
			Description: "测试商品8描述",
			Content: "测试商品8详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 3,
			CateId: 3,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4",
			Score: 4,
		},
		{
			Name: "测试商品9",
			Description: "测试商品9描述",
			Content: "测试商品9详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 4,
			CateId: 4,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品11",
			Description: "测试商品11描述",
			Content: "测试商品11详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 1,
			CateId: 1,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品12",
			Description: "测试商品12描述",
			Content: "测试商品12详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 5,
			CateId: 5,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品123",
			Description: "测试商品123描述",
			Content: "测试商品123详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 1,
			CateId: 1,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品111",
			Description: "测试商品111描述",
			Content: "测试商品111详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 2,
			CateId: 2,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品22",
			Description: "测试商品22描述",
			Content: "测试商品22详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 2,
			CateId: 2,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品33",
			Description: "测试商品33描述",
			Content: "测试商品33详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 3,
			CateId: 3,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品44",
			Description: "测试商品44描述",
			Content: "测试商品44详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 4,
			CateId: 4,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品55",
			Description: "测试商品55描述",
			Content: "测试商品55详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 5,
			CateId: 5,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3",
			Score: 4,
		},
		{
			Name: "测试商品66",
			Description: "测试商品66描述",
			Content: "测试商品66详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 2,
			CateId: 2,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品77",
			Description: "测试商品77描述",
			Content: "测试商品77详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 3,
			CateId: 3,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品88",
			Description: "测试商品88描述",
			Content: "测试商品88详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 4,
			CateId: 4,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4,5",
			Score: 4,
		},
		{
			Name: "测试商品99",
			Description: "测试商品99描述",
			Content: "测试商品99详情",
			Icon: "upload/goods/shop_id/2021001.jpg",
			Price: 20,
			OldPrice: 25,
			GoodsUnit: "箱",
			TotalNum: 200,
			SellCount: 30,
			MinNum: 5,
			DeliveryPrice: 22,
			OverPrice: 8,
			ShopId: 5,
			CateId: 5,
			IsHot: 1,
			IsMain: 1,
			IsTop: 1,
			State: 1,
			SaleId: "1,2,3,4",
			Score: 4,
		},
	}
	//事务
	session := dal.DB.NewSession()
	defer session.Close()
	//事务操作：事务开始, 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, good := range goods {
		_, err = session.Insert(&good)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit() //提交事务
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 向supplier表中插入初始测试数据
 */
func InitSuppData() {
	suppliers := []model.Supplier{
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
		{Name: "测试商家111", Description   : "测试商家111描述", Content       : "测试商家111详情", Phone         : "0390-0123456", Mobile        : "15516132888", Address       : "濮阳市华龙区胜利东路", Lng           : 31.179695, Lat           : 121.406051, Icon          : "upload/supp/2021001.jpg", TotalCount    : 50, MinNum        : 10, DeliveryPrice : 11, OverPrice     : 6, Rating        : 4.8, RatingCount   : 5, OpeningHours  : "周一至周六全天", GiveId        : "1,2,3,4,5", SaleId        : "1,2,3,4,5", IsHot         : 1, IsTop         : 1, IsNew         : 1, State         : 1, IsUnion		  : 1, IsMain        : 1},
	}
	//事务
	session := dal.DB.NewSession()
	defer session.Close()
	//事务操作：事务开始, 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, supplier := range suppliers {
		_, err = session.Insert(&supplier)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit() //提交事务
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 向cates表中插入初始测试数据
 */
func InitShopCates() {
	cates := []model.ShopCates{
		{ScName:"热销",ScIcon:"upload/shopcate/shop_1/hot.png",ShopId: 1,GoodIds:"1,8,9,10",State:1,Position: 1,Desc:"百年经典"},
		{ScName:"优惠",ScIcon:"upload/shopcate/shop_1/youhui.png",ShopId: 1,GoodIds:"12,17,19",State:1,Position: 2,Desc:"优品钜惠"},
		{ScName:"百岁山矿泉水",ScIcon:"",ShopId: 1,GoodIds:"21,23,25,32,33",State:1,Position: 3,Desc:"百年经典"},
		{ScName:"百岁山纯净水",ScIcon:"",ShopId: 1,GoodIds:"21,23,25,32,33",State:1,Position: 4,Desc:"优品钜惠"},
		{ScName:"百岁山桶装水",ScIcon:"",ShopId: 1,GoodIds:"21,23,25,32,33",State:1,Position: 5,Desc:"百年经典"},
		{ScName:"百岁山商用水",ScIcon:"",ShopId: 1,GoodIds:"21,23,25,32,33",State:1,Position: 6,Desc:"优品钜惠"},
		{ScName:"百岁山家用水",ScIcon:"",ShopId: 1,GoodIds:"21,23,25,32,33",State:1,Position: 7,Desc:"优品钜惠"},
		{ScName:"热销",ScIcon:"upload/shopcate/shop_1/hot.png",ShopId: 2,GoodIds:"2,13,14,23",State:1,Position: 1,Desc:"百年经典"},
		{ScName:"优惠",ScIcon:"upload/shopcate/shop_1/youhui.png",ShopId: 2,GoodIds:"18,20,22,31",State:1,Position: 2,Desc:"优品钜惠"},
		{ScName:"农夫山泉矿泉水",ScIcon:"",ShopId: 2,GoodIds:"26,27,28,29,30",State:1,Position: 3,Desc:"优品钜惠"},
		{ScName:"农夫山泉纯净水",ScIcon:"",ShopId: 2,GoodIds:"26,27,28,29,30",State:1,Position: 4,Desc:"优品钜惠"},
		{ScName:"农夫山泉桶装水",ScIcon:"",ShopId: 2,GoodIds:"26,27,28,29,30",State:1,Position: 5,Desc:"优品钜惠"},
		{ScName:"农夫山泉商用水",ScIcon:"",ShopId: 2,GoodIds:"26,27,28,29,30",State:1,Position: 6,Desc:"优品钜惠"},
		{ScName:"农夫山泉家用水",ScIcon:"",ShopId: 2,GoodIds:"26,27,28,29,30",State:1,Position: 7,Desc:"百年经典"},
	}
	//事务
	session := dal.DB.NewSession()
	defer session.Close()
	//事务操作：事务开始, 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, cate := range cates {
		_, err = session.Insert(&cate)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit() //提交事务
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 向eva表中插入初始测试数据
 */
func InitEva() {
	evas := []model.Evaluation {
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "好评好评1",IsFine: 1,IsGeneral: 0,IsBad: 0,Score: 4.8,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "好评好评2",IsFine: 1,IsGeneral: 0,IsBad: 0,Score: 4.8,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "好评好评3",IsFine: 1,IsGeneral: 0,IsBad: 0,Score: 4.8,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "中评中评4",IsFine: 0,IsGeneral: 1,IsBad: 0,Score: 3.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "中评中评5",IsFine: 0,IsGeneral: 1,IsBad: 0,Score: 3.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "中评中评6",IsFine: 0,IsGeneral: 1,IsBad: 0,Score: 3.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "中评中评7",IsFine: 0,IsGeneral: 1,IsBad: 0,Score: 3.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "差评差评8",IsFine: 0,IsGeneral: 0,IsBad: 1,Score: 1.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "差评差评9",IsFine: 0,IsGeneral: 0,IsBad: 1,Score: 1.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
		{ParentId: 0,UserId: 1,ShopId: 2,GoodId: 2,UserName:"yonghu1111",ContentText: "差评差评10",IsFine: 0,IsGeneral: 0,IsBad: 1,Score: 1.5,PackScore: 4.8,ServeScore: 4.8,DeliveryScore: 4.8,State: 1},
	}

	//事务
	session := dal.DB.NewSession()
	defer session.Close()
	//事务操作：事务开始, 执行操作（回滚），提交事务
	err := session.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, eva := range evas {
		_, err = session.Insert(&eva)
		if err != nil {
			session.Rollback()
			return
		}
	}
	err = session.Commit() //提交事务
	if err != nil {
		fmt.Println(err.Error())
	}
}
