package param

import "orderserve/model"

type AddUserAddrParam struct {
	UserToken string `json:"user_id"`
	AddressInfo model.UserAddress `json:"addressInfo"`
}
