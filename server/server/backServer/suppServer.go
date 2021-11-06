package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type ISupp interface {
	// 获取所有商家数据
	GetSuppDataAll() (suppliersData []model.Supplier,err error)
	// 根据ID 更新商家banner
	EditSuppBanner(id int64,fileName string) (err error)
	// 根据ID 更新商家logo
	EditSuppLogo(id int64,fileName string) (err error)
	// 根据ID  更新商家数据
	EditSuppById(id int64,suppData model.Supplier) (err error)
	// 添加商家
	AddSuppOne(suppData model.Supplier) (err error)
}

type SuppServer struct {}

func NewSuppServer() ISupp {
	return &SuppServer{}
}

var suppDal = backendDal.NewSuppDal()

// 获取所有商家数据
func (supS *SuppServer) GetSuppDataAll() (suppliersData []model.Supplier,err error) {
	return suppDal.QuerySuppDataAll()
}
// 根据ID 更新商家banner
func (supS *SuppServer) EditSuppBanner(id int64,fileName string) (err error) {
	return suppDal.UpdateSuppBanner(id,fileName)
}
// 根据ID 更新商家logo
func (supS *SuppServer) EditSuppLogo(id int64,fileName string) (err error) {
	return suppDal.UpdateSuppLogo(id,fileName)
}
// 根据ID  更新商家数据
func (supS *SuppServer) EditSuppById(id int64,suppData model.Supplier) (err error) {
	return suppDal.UpdateSuppById(id,suppData)
}
// 添加商家
func (supS *SuppServer) AddSuppOne(suppData model.Supplier) (err error) {
	return suppDal.InsertSuppOne(suppData)
}