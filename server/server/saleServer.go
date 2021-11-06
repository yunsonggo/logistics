package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type ISale interface {
	AddSale(saleInfo *model.Sale) (err error)
	DelSale(id int64) (err error)
	UpdateSale(id int64, saleInfo *model.Sale) (err error)
	GetSale(id int64) (saleData *model.Sale,err error)
	GetSaleAll() (salesData []model.Sale,err error)
	GetSaleByIdArray(saleIds []interface{}) (saleData []model.Sale,err error)
}

var sd = dal.NewSaleDal()

type SaleServer struct {}

func NewSaleServer() ISale {
	return &SaleServer{}
}

func (ss *SaleServer) AddSale(saleInfo *model.Sale) (err error) {
	return sd.InsertSaleOne(saleInfo)
}
func (ss *SaleServer) DelSale(id int64) (err error) {
	return sd.DelSaleById(id)
}
func (ss *SaleServer) UpdateSale(id int64, saleInfo *model.Sale) (err error) {
	return sd.UpdateSaleById(id,saleInfo)
}
func (ss *SaleServer) GetSale(id int64) (saleData *model.Sale,err error) {
	return sd.QuerySaleById(id)
}
func (ss *SaleServer) GetSaleAll() (salesData []model.Sale,err error) {
	return sd.QuerySaleAll()
}

func (ss *SaleServer) GetSaleByIdArray(saleIds []interface{}) (saleData []model.Sale,err error) {
	return sd.QuerySaleByIdArray(saleIds)
}
