package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type ISaleGoods interface {
	// 插入一条已售商品数据
	AddSaleGoodOne(saleGood *model.SaleGoods) (err error)
	// 根据 识别码 查询商品
	GetSaleGoodByCode(code string) (saleGood *model.SaleGoods,err error)
	// 根据ID 查询已售商品
	GetSaleGoodById(id int64) (saleGood *model.SaleGoods,err error)
}

type SaleGoodsServer struct {}

func NewSaleGoodsServer() ISaleGoods {
	return &SaleGoodsServer{}
}
var sgd = dal.NewSaleGoodsDal()
// 插入一条已售商品数据
func (sgs *SaleGoodsServer) AddSaleGoodOne(saleGood *model.SaleGoods) (err error) {
	return sgd.InsertSaleGoodOne(saleGood)
}
// 根据 识别码 查询商品
func (sgs *SaleGoodsServer) GetSaleGoodByCode(code string) (saleGood *model.SaleGoods,err error) {
	return sgd.QuerySaleGoodByCode(code)
}
// 根据ID 查询已售商品
func (sgs *SaleGoodsServer) GetSaleGoodById(id int64) (saleGood *model.SaleGoods,err error) {
	return sgd.QuerySaleGoodById(id)
}