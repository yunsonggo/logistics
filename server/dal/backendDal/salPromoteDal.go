package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type SalePromoteInterface interface {
	// 获取所有活动数据
	QuerySalePromoteAll() (salePromotesData []model.Sale, err error)
	// 修改活动数据
	UpdateSalePromote(id int64,promoteData model.Sale) (err error)
	// 新增活动数据
	InsertSalePromoteOne(promoteData model.Sale) (err error)
}

type SalePromoteDal struct {}

func NewSalPromoteDal() SalePromoteInterface {
	return &SalePromoteDal{}
}
// 获取所有活动数据
func (spd *SalePromoteDal) QuerySalePromoteAll() (salePromotesData []model.Sale, err error) {
	err = dal.DB.Find(&salePromotesData)
	return
}
// 修改活动数据
func (spd *SalePromoteDal) UpdateSalePromote(id int64,promoteData model.Sale) (err error) {
	_,err = dal.DB.Where("id = ?",id).Update(&promoteData)
	return
}
// 新增活动数据
func (spd *SalePromoteDal) InsertSalePromoteOne(promoteData model.Sale) (err error) {
	_,err = dal.DB.InsertOne(&promoteData)
	return
}