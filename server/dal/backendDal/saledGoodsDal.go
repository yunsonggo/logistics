package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type SaleGoodsInterface interface {
	// 根据订单ID 查询
	QuerySaleGoodsByOrderId(id int64) (saleGoods []model.SaleGoods, err error)
}

type SaleGoodsDal struct {}

func NewSaleGoodsDal() SaleGoodsInterface {
	return &SaleGoodsDal{}
}

// 根据订单ID 查询
func (sgd *SaleGoodsDal) QuerySaleGoodsByOrderId(id int64) (saleGoods []model.SaleGoods, err error) {
	err = dal.DB.Where("order_id = ?",id).Desc("create_time").Find(&saleGoods)
	return
}