package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IOrders interface {
	// 插入一条订单
	AddOrderOne(orderData *model.Orders) (err error)
	// 统计订单数量
	GetOrderCount() (total int64, err error)
	// 根据识别码 查询 订单
	GetOrderByOrderCode(code string) (orderData *model.Orders,err error)
	// 更新订单
	UpdateOrderOne(id int64,orderData *model.Orders) (err error)
	// 根据用户ID 获取所有订单 ID倒序
	GetOrderAllByUserId(id int64) (ordersData []model.Orders,err error)
}

type OrdersServer struct {}

func NewOrderServer() IOrders {
	return &OrdersServer{}
}

var orderDal = dal.NewOrderDal()

// 插入一条订单
func (ods *OrdersServer) AddOrderOne(orderData *model.Orders) (err error) {
	return orderDal.InsertOrderOne(orderData)
}
// 统计订单数量
func (ods *OrdersServer) GetOrderCount() (total int64, err error) {
	return orderDal.QueryOrderCount()
}
// 根据识别码 查询 订单
func (ods *OrdersServer) GetOrderByOrderCode(code string) (orderData *model.Orders,err error) {
	return orderDal.QueryOrderByOrderCode(code)
}
// 更新订单
func (ods *OrdersServer) UpdateOrderOne(id int64,orderData *model.Orders) (err error) {
	return orderDal.UpdateOrderById(id,orderData)
}
// 根据用户ID 获取所有订单 ID倒序
func (ods *OrdersServer) GetOrderAllByUserId(id int64) (ordersData []model.Orders,err error) {
	return orderDal.QueryOrderAllByUserId(id)
}