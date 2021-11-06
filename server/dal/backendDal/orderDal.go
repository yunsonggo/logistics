package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type OrderInterface interface {
	// 获取所有订单数据
	QueryOrderAll() (ordersData []model.Orders, err error)
}
type OrderDal struct {}

func NewOrderDal() OrderInterface {
	return &OrderDal{}
}
// 获取所有订单数据
func (ord *OrderDal) QueryOrderAll() (ordersData []model.Orders, err error) {
	err = dal.DB.Find(&ordersData)
	return
}