package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type ISalePromote interface {
	// 获取所有活动数据
	GetSalePromoteAll() (salePromotesData []model.Sale, err error)
	// 修改活动数据
	EditSalePromote(id int64,promoteData model.Sale) (err error)
	// 新增活动数据
	AddSalePromoteOne(promoteData model.Sale) (err error)
}

type SalePromoteServer struct {}

func NewSalePromoteServer() ISalePromote {
	return &SalePromoteServer{}
}
var spd = backendDal.NewSalPromoteDal()
// 获取所有活动数据
func (sps *SalePromoteServer) GetSalePromoteAll() (salePromotesData []model.Sale, err error) {
	return spd.QuerySalePromoteAll()
}
// 修改活动数据
func (sps *SalePromoteServer) EditSalePromote(id int64,promoteData model.Sale) (err error) {
	return spd.UpdateSalePromote(id,promoteData)
}
// 新增活动数据
func (sps *SalePromoteServer) AddSalePromoteOne(promoteData model.Sale) (err error) {
	return spd.InsertSalePromoteOne(promoteData)
}