package dal

import "orderserve/model"

type GiveInterface interface {
	InsertGiveOne(giveInfo *model.Give) (err error)
	DelGiveById(id int64) (err error)
	UpdateGiveById(id int64, giveInfo *model.Give) (err error)
	QueryGiveById(id int64) (giveData *model.Give,err error)
	QueryGiveAll() (givesData []model.Give,err error)
	QueryGiveByIdArray(giveIds []interface{}) (givesData []model.Give,err error)
}

type GiveDal struct {}

func NewGiveDal() GiveInterface {
	return &GiveDal{}
}

func (gd *GiveDal) InsertGiveOne(giveInfo *model.Give) (err error) {
	_,err = DB.InsertOne(giveInfo)
	return
}
func (gd *GiveDal) DelGiveById(id int64) (err error) {
	give := new(model.Give)
	_,err =DB.ID(id).Delete(give)
	return
}
func (gd *GiveDal) UpdateGiveById(id int64, giveInfo *model.Give) (err error) {
	_,err = DB.ID(id).Update(giveInfo)
	return
}
func (gd *GiveDal) QueryGiveById(id int64) (giveData *model.Give,err error) {
	_,err = DB.ID(id).Get(giveData)
	return
}
func (gd *GiveDal) QueryGiveAll() (givesData []model.Give,err error) {
	err = DB.Find(&givesData)
	return
}

func (gd *GiveDal) QueryGiveByIdArray(giveIds []interface{}) (givesData []model.Give,err error) {
	err = DB.In("id",giveIds).Find(&givesData)
	return
}
