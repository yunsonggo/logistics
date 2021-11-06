package dal

import "orderserve/model"

type SaleGoodsInterface interface {
	// 插入一条已售商品数据
	InsertSaleGoodOne(saleGood *model.SaleGoods) (err error)
	// 根据 识别码 查询商品
	QuerySaleGoodByCode(code string) (saleGood *model.SaleGoods,err error)
	// 根据ID 查询已售商品
	QuerySaleGoodById(id int64) (saleGood *model.SaleGoods,err error)
}

type SaleGoodsDal struct {}

func NewSaleGoodsDal() SaleGoodsInterface {
	return &SaleGoodsDal{}
}

func (sgd *SaleGoodsDal) InsertSaleGoodOne(saleGood *model.SaleGoods) (err error) {
	_,err = DB.InsertOne(saleGood)
	return
}

// 根据 识别码 查询商品
func (sgd *SaleGoodsDal) QuerySaleGoodByCode(code string) (saleGood *model.SaleGoods,err error) {
	var good model.SaleGoods
	_,err = DB.Where("sale_code = ?",code).Get(&good)
	saleGood = &good
	return
}
// 根据ID 查询已售商品
func (sgd *SaleGoodsDal) QuerySaleGoodById(id int64) (saleGood *model.SaleGoods,err error) {
	var good model.SaleGoods
	_,err = DB.Where("id = ?",id).Get(&good)
	saleGood = &good
	return
}