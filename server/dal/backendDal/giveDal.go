package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type GivesInterface interface {
	// 获取所有服务数据
	QueryGivesAll() (givesData []model.Give, err error)
	// 更新服务数据
	UpdateGiveOne(id int64,giveData model.Give) (err error)
	// 添加服务数据
	InsertGiveOne(giveData model.Give) (err error)
}

type GiveDal struct {}

func NewGiveDal() GivesInterface {
	return &GiveDal{}
}

// 获取所有服务数据
func (gid *GiveDal) QueryGivesAll() (givesData []model.Give, err error) {
	err = dal.DB.Find(&givesData)
	return
}
// 更新服务数据
func (gid *GiveDal) UpdateGiveOne(id int64,giveData model.Give) (err error) {
	_,err = dal.DB.Where("id = ?",id).Update(&giveData)
	return
}
// 添加服务数据
func (gid *GiveDal) InsertGiveOne(giveData model.Give) (err error) {
	_,err = dal.DB.InsertOne(&giveData)
	return
}