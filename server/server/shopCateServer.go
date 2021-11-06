package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IShopCate interface {
	//根据商家ID 查询商家商品分类
	GetShopCateByShop(shopId int64) (shopCates []model.ShopCates,err error)
}

type ShopCateServer struct {}

func NewShopCateServer() IShopCate {
	return &ShopCateServer{}
}

var scd = dal.NewShopCateDal()
//根据商家ID 查询商家商品分类
func (scs *ShopCateServer) GetShopCateByShop(shopId int64) (shopCates []model.ShopCates,err error) {
	return scd.QueryShopCateByShop(shopId)
}
