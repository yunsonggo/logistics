package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type ISaleGoods interface {
	// 根据订单ID 查询
	GetSaleGoodsByOrderId(id int64) (saleGoods []model.SaleGoods, err error)
}

type SaleGoodsServer struct {}

func NewSaleGoodsServer() ISaleGoods {
	return &SaleGoodsServer{}
}

var sgd = backendDal.NewSaleGoodsDal()

// 根据订单ID 查询
func (sgs *SaleGoodsServer) GetSaleGoodsByOrderId(id int64) (saleGoods []model.SaleGoods, err error) {
	return sgd.QuerySaleGoodsByOrderId(id)
}
