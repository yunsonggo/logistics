package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IUserAddr interface {
	// 添加地址
	AddUserAddrOne(userAdd *model.UserAddress) (err error)
	// 根据用户ID 查询所有地址
	GetUserAddrById(userId int64) (userAdds []model.UserAddress,err error)
	// 根据用户ID  清除默认收货地址
	ClearAddrPosition(userId int64) (err error)
	// 根据用户ID 地址ID  修改 一条
	EditOneAddrById(userId int64,userAddr *model.UserAddress) (err error)
	// 根据用户ID 地址ID  删除一条
	RemoveOneAddrById(userId int64,userAddr *model.UserAddress) (err error)
}

type UserAddrServer struct {}

func NewUserAddrServer() IUserAddr {
	return &UserAddrServer{}
}

var addrDal = dal.NewUserAddrDal()

func (uas *UserAddrServer) AddUserAddrOne(userAdd *model.UserAddress) (err error) {
	return addrDal.InsertUserAddrOne(userAdd)
}

func (uas *UserAddrServer) GetUserAddrById(userId int64) (userAdds []model.UserAddress,err error) {
	return addrDal.QueryUserAddrById(userId)
}

func (uas *UserAddrServer) ClearAddrPosition(userId int64) (err error) {
	return addrDal.UpdateAddrPosition(userId)
}
// 根据用户ID 地址ID  修改 一条
func (uas *UserAddrServer) EditOneAddrById(userId int64,userAddr *model.UserAddress) (err error) {
	return addrDal.UpdateOneAddrById(userId,userAddr)
}
// 根据用户ID 地址ID  删除一条
func (uas *UserAddrServer) RemoveOneAddrById(userId int64,userAddr *model.UserAddress) (err error) {
	return addrDal.DeleteOneAddrById(userId,userAddr)
}
