package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type SuppInterface interface {
	// 获取所有商家数据
	QuerySuppDataAll() (suppliersData []model.Supplier,err error)
	// 根据ID 更新商家banner
	UpdateSuppBanner(id int64,fileName string) (err error)
	// 根据ID 更新商家logo
	UpdateSuppLogo(id int64,fileName string) (err error)
	// 根据ID  更新商家数据
	UpdateSuppById(id int64,suppData model.Supplier) (err error)
	// 添加商家
	InsertSuppOne(suppData model.Supplier) (err error)
}

type SuppDal struct {}

func NewSuppDal() SuppInterface {
	return &SuppDal{}
}

// 获取所有商家数据
func (supD *SuppDal) QuerySuppDataAll() (suppliersData []model.Supplier,err error) {
	err = dal.DB.Asc("id").Find(&suppliersData)
	return
}
// 根据ID 更新商家banner
func (supD *SuppDal) UpdateSuppBanner(id int64,fileName string) (err error) {
	newSupp := new(model.Supplier)
	newSupp.Banner = fileName
	_,err = dal.DB.Where("id = ?",id).Update(newSupp)
	return
}
// 根据ID 更新商家logo
func (supD *SuppDal) UpdateSuppLogo(id int64,fileName string) (err error) {
	newSupp := new(model.Supplier)
	newSupp.Icon = fileName
	_,err = dal.DB.Where("id = ?",id).Update(newSupp)
	return
}
// 根据ID  更新商家数据
func (supD *SuppDal) UpdateSuppById(id int64,suppData model.Supplier) (err error) {
	_,err = dal.DB.Where("id = ?",id).Update(&suppData)
	return
}
// 添加商家
func (supD *SuppDal) InsertSuppOne(suppData model.Supplier) (err error) {
	_,err = dal.DB.InsertOne(&suppData)
	return
}