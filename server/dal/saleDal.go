package dal

import "orderserve/model"

type SaleInterface interface {
	InsertSaleOne(saleInfo *model.Sale) (err error)
	DelSaleById(id int64) (err error)
	UpdateSaleById(id int64, saleInfo *model.Sale) (err error)
	QuerySaleById(id int64) (saleData *model.Sale,err error)
	QuerySaleAll() (salesData []model.Sale,err error)
	QuerySaleByIdArray(saleIds []interface{}) (saleData []model.Sale,err error)
}

type SaleDal struct {}

func NewSaleDal() SaleInterface {
	return &SaleDal{}
}

func (sd *SaleDal) InsertSaleOne(saleInfo *model.Sale) (err error) {
	_,err = DB.InsertOne(saleInfo)
	return
}
func (sd *SaleDal) DelSaleById(id int64) (err error) {
	sale := new(model.Sale)
	_,err =DB.ID(id).Delete(sale)
	return
}
func (sd *SaleDal) UpdateSaleById(id int64, saleInfo *model.Sale) (err error) {
	_,err = DB.ID(id).Update(saleInfo)
	return
}
func (sd *SaleDal) QuerySaleById(id int64) (saleData *model.Sale,err error) {
	_,err = DB.ID(id).Get(saleData)
	return
}
func (sd *SaleDal) QuerySaleAll() (salesData []model.Sale,err error) {
	err = DB.Find(&salesData)
	return
}

func (sd *SaleDal) QuerySaleByIdArray(saleIds []interface{}) (saleData []model.Sale,err error) {
	err = DB.In("id",saleIds).Find(&saleData)
	return
}
