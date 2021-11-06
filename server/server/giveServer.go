package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IGive interface {
	AddGive(giveInfo *model.Give) (err error)
	DelGive(id int64) (err error)
	UpdateGive(id int64, giveInfo *model.Give) (err error)
	GetGive(id int64) (giveData *model.Give,err error)
	GetGiveAll() (givesData []model.Give,err error)
	GetGiveByIdArray(giveIds []interface{}) (givesData []model.Give,err error)
}

var gd = dal.NewGiveDal()

type GiveServer struct {}

func NewGiveServer() IGive {
	return &GiveServer{}
}

func (gs *GiveServer) AddGive(giveInfo *model.Give) (err error) {
	return gd.InsertGiveOne(giveInfo)
}
func (gs *GiveServer) DelGive(id int64) (err error) {
	return gd.DelGiveById(id)
}
func (gs *GiveServer) UpdateGive(id int64, giveInfo *model.Give) (err error) {
	return gd.UpdateGiveById(id,giveInfo)
}
func (gs *GiveServer) GetGive(id int64) (giveData *model.Give,err error) {
	return gd.QueryGiveById(id)
}

func (gs *GiveServer) GetGiveAll() (givesData []model.Give,err error) {
	return gd.QueryGiveAll()
}

func (gs *GiveServer) GetGiveByIdArray(giveIds []interface{}) (givesData []model.Give,err error) {
	return gd.QueryGiveByIdArray(giveIds)
}
