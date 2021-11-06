package param

import "orderserve/model"

type JwtParam struct {
	JwtUserId string `json:"user_id"`
	UserId string `json:"userid"`
	AddressInfo model.UserAddress `json:"addressInfo"`
	OrderList interface{} `json:"order_list"`
}
