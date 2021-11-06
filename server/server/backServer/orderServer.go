package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type IOrder interface {
	// 获取所有订单数据
	GetOrderAll() (ordersData []model.Orders, err error)
}
type OrderServer struct {}

var ord = backendDal.NewOrderDal()

func NewOrderServer() IOrder {
	return &OrderServer{}
}
// 获取所有订单数据
func (ors *OrderServer) GetOrderAll() (ordersData []model.Orders, err error) {
	return ord.QueryOrderAll()
}