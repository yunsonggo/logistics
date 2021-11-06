package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type IGive interface {
	// 获取所有服务数据
	GetGivesAll() (givesData []model.Give, err error)
	// 更新服务数据
	EditGiveOne(id int64,giveData model.Give) (err error)
	// 添加服务数据
	AddGiveOne(giveData model.Give) (err error)
}

type GiveServer struct {}

func NewGiveServer() IGive {
	return &GiveServer{}
}

var gid = backendDal.NewGiveDal()

// 获取所有服务数据
func (gis *GiveServer) GetGivesAll() (givesData []model.Give, err error) {
	return gid.QueryGivesAll()
}
// 更新服务数据
func (gis *GiveServer) EditGiveOne(id int64,giveData model.Give) (err error) {
	return gid.UpdateGiveOne(id,giveData)
}
// 添加服务数据
func (gis *GiveServer) AddGiveOne(giveData model.Give) (err error) {
	return gid.InsertGiveOne(giveData)
}