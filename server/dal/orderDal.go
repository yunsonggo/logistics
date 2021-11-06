package dal

import "orderserve/model"

type OrderInterface interface {
	// 插入一条订单
	InsertOrderOne(orderData *model.Orders) (err error)
	// 统计订单数量
	QueryOrderCount() (total int64, err error)
	// 根据识别码 查询 订单
	QueryOrderByOrderCode(code string) (orderData *model.Orders,err error)
	// 更新订单
	UpdateOrderById(id int64,orderData *model.Orders) (err error)
	// 根据用户ID 获取所有订单 ID倒序
	QueryOrderAllByUserId(id int64) (ordersData []model.Orders,err error)
}

type OrderDal struct {}

func NewOrderDal() OrderInterface {
	return &OrderDal{}
}
// 插入一条订单
func (od *OrderDal) InsertOrderOne(orderData *model.Orders) (err error) {
	_,err = DB.InsertOne(orderData)
	return
}
// 统计订单数量
func (od *OrderDal) QueryOrderCount() (total int64, err error) {
	order := new(model.Orders)
	total,err = DB.Where("id > ?",1).Count(order)
	return
}
// 根据识别码 查询 订单
func (od *OrderDal) QueryOrderByOrderCode(code string) (orderData *model.Orders,err error) {
	var order model.Orders
	_,err = DB.Where("order_code = ?",code).Get(&order)
	orderData = &order
	return
}
// 更新订单
func (od *OrderDal) UpdateOrderById(id int64,orderData *model.Orders) (err error) {
	_,err = DB.Where("id = ?",id).Update(orderData)
	return
}
// 根据用户ID 获取所有订单 ID倒序
func (od *OrderDal) QueryOrderAllByUserId(id int64) (ordersData []model.Orders,err error) {
	err = DB.Where("user_id = ?",id).Desc("id").Find(&ordersData)
	return
}