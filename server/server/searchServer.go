package server

import "orderserve/model"

var suppSer = NewSupplierServer()
var goodSer = NewGoodServer()

//关键词搜索商家
func SearchSuppByName(key string,name []string) (supplierData []model.Supplier,err error) {
	return suppSer.GetSupplierByName(key,name)
}
//id列表 搜索 商家
func SearchSuppByKey(key string,data []int64) (supplierData []model.Supplier,err error) {
	return suppSer.GetSupplierByIds(key,data)
}
//关键词 搜索 商品
func SearchGoodBykey(key string,data []string) (goodsData []model.Goods,err error) {
	return goodSer.GetGoodsByKeywords(key,data)
}
// 根据某字段 升降序 排列
func SearchGoodsOrder(key string,order string ,count int, start int) (goodsData []model.Goods,err error) {
	return goodSer.GetGoodsByKeyOrder(key,order,count,start)
}

func SearchGoodsByArray (giveIds []interface{},saleIds []interface{},minPrice float32,maxPrice float32,count int,page int) (goodsData []model.Goods,err error) {
	start := page * count - (count-1)
	return goodSer.GetByArrayAndPrice(giveIds,saleIds,minPrice,maxPrice,count,start)
}
