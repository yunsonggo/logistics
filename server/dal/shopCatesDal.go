package dal

import "orderserve/model"

type  ShopCateInterface interface {
	//根据商家ID 查询商家商品分类
	QueryShopCateByShop(shopId int64) (shopCates []model.ShopCates,err error)
}

type ShopCateDal struct {}

func NewShopCateDal() ShopCateInterface {
	return &ShopCateDal{}
}
//根据商家ID 查询商家商品分类
func (scd *ShopCateDal) QueryShopCateByShop(shopId int64) (shopCates []model.ShopCates,err error) {
	err = DB.Where("shop_id = ? and state = 1",shopId).Find(&shopCates)
	return
}
