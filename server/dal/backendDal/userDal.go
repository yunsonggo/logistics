package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type UserBackInterface interface {
	// 获取所有用户数据
	QueryUserAll() (usersData []model.User,err error)
}

type UserBackDal struct {}

func NewUserBackDal() UserBackInterface {
	return &UserBackDal{}
}
// 获取所有用户数据
func (ubd *UserBackDal) QueryUserAll() (usersData []model.User,err error) {
	err = dal.DB.Find(&usersData)
	return
}