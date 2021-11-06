package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IUser interface {
	SignUpUserBySms(phone string) (userId int64, err error)
	GetUserByPhone(phone string) (userData *model.User, err error)
	GetUserById(id interface{}) (userData model.User,err error)
}

var userDal = dal.NewUserDal()

type UserServer struct{}

func NewUserServer() IUser {
	return &UserServer{}
}

// 根据短信 注册新用户
func (us *UserServer) SignUpUserBySms(phone string) (userId int64, err error) {
	return userDal.InsertUserBySms(phone)
}

// 根据手机号 查询用户
func (us *UserServer) GetUserByPhone(phone string) (userData *model.User, err error) {
	return userDal.QueryUserByPhone(phone)
}

// 根据ID  查询用户
func (us *UserServer) GetUserById(id interface{}) (userData model.User,err error) {
	return userDal.QueryUserById(id)
}
