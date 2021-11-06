package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IGoods interface {
	AddGood(goodInfo *model.Goods) (err error)
	DelGood(id int64) (err error)
	UpdateGood(id int64,goodInfo *model.Goods) (err error)
	GetGood(id int64) (goodData *model.Goods,err error)
	GetGoodsList() (goodsData []model.Goods,err error)
	GetGoodsAny(key string,data interface{}) (goodsData []model.Goods,err error)
	GetGoodsByKeywords(key string,data []string) (goodsData []model.Goods,err error)
	GetGoodsByKeyOrder(key string,order string,count int,start int) (goodsData []model.Goods,err error)
	GetByArrayAndPrice(giveIds []interface{},saleIds []interface{},minPrice float32,maxPrice float32,count int,start int) (goodsData []model.Goods,err error)
}

type GoodServer struct {}

var goodsDal = dal.NewGoodsDal()

func NewGoodServer () IGoods {
	return &GoodServer{}
}

func (gs *GoodServer) AddGood(goodInfo *model.Goods) (err error) {
	return goodsDal.InsertGoodOne(goodInfo)
}

func (gs *GoodServer) DelGood(id int64) (err error) {
	return goodsDal.DelGoodById(id)
}
func (gs *GoodServer) UpdateGood(id int64,goodInfo *model.Goods) (err error) {
	return goodsDal.UpdateGoodById(id,goodInfo)
}
func (gs *GoodServer) GetGood(id int64) (goodData *model.Goods,err error) {
	return goodsDal.QueryGoodById(id)
}
func (gs *GoodServer) GetGoodsList() (goodsData []model.Goods,err error) {
	return goodsDal.QueryGoodsAll()
}
// 根据字段查询
func (gs *GoodServer) GetGoodsAny(key string,data interface{}) (goodsData []model.Goods,err error) {
	return goodsDal.QueryGoodsByAny(key,data)
}
// 根据关键词模糊查询
func (gs *GoodServer) GetGoodsByKeywords(key string,data []string) (goodsData []model.Goods,err error) {
	for _,val := range data {
		res,resErr := goodsDal.QueryGoodsByKeywords(key,val)
		if resErr != nil {
			return
		}
		for _,v := range res {
			goodsData = append(goodsData, v)
		}
	}
	return
}
// 某字段  升 / 降序
func (gs *GoodServer) GetGoodsByKeyOrder(key string,order string,count int,start int) (goodsData []model.Goods,err error) {
	return goodsDal.QueryGoodsByKeyOrder(key,order,count,start)
}
func (gs *GoodServer) GetByArrayAndPrice(giveIds []interface{},saleIds []interface{},minPrice float32,maxPrice float32,count int,start int) (goodsData []model.Goods,err error) {
	return goodsDal.QueryByArrayAndPrice(giveIds,saleIds,minPrice,maxPrice,count,start)
}
