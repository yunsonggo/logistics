package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type IUserBack interface {
	// 获取所有用户数据
	GetUserAll() (usersData []model.User,err error)
}
type UserBackServer struct {}

func NewUserBackServer() IUserBack {
	return &UserBackServer{}
}

var ubd = backendDal.NewUserBackDal()
// 获取所有用户数据
func (ubs *UserBackServer) GetUserAll() (usersData []model.User,err error) {
	return ubd.QueryUserAll()
}
