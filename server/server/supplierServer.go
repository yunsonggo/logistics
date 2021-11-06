package server

import (
	"orderserve/dal"
	"orderserve/model"
	"orderserve/param"
)

type ISupplier interface {
	AddSupplier(supplierInfo *model.Supplier) (err error)
	DelSupplier(id int64) (err error)
	UpdateSupplier(id int64,supplierInfo *model.Supplier) (err error)
	GetSupplier(id int64) (supplierData *model.Supplier,err error)
	GetSupplierList() (supplierData []model.Supplier,err error)
	GetSupplierAny(key string,data interface{}) (supplierData []model.Supplier,err error)
	GetSupplierByKeys(param *param.SupplierPatam) (supplierData []model.Supplier,err error)
	GetSupplierByIds(key string,data []int64) (supplierData []model.Supplier,err error)
	GetSupplierByName(key string,name []string) (supplierData []model.Supplier,err error)
	GetSupplierByCateId(cateId int64) (supplierData []model.Supplier,err error)
}

type SupplierServer struct {}

var sud = dal.NewSupplierDal()

func NewSupplierServer () ISupplier {
	return &SupplierServer{}
}

func (sus *SupplierServer) AddSupplier(supplierInfo *model.Supplier) (err error) {
	return sud.InsertSupplierOne(supplierInfo)
}

func (sus *SupplierServer) DelSupplier(id int64) (err error) {
	return sud.DelSupplierById(id)
}
func (sus *SupplierServer) UpdateSupplier(id int64,supplierInfo *model.Supplier) (err error) {
	return sud.UpdateSupplierById(id,supplierInfo)
}
func (sus *SupplierServer) GetSupplier(id int64) (supplierData *model.Supplier,err error) {
	return sud.QuerySupplierById(id)
}
func (sus *SupplierServer) GetSupplierList() (supplierData []model.Supplier,err error) {
	return sud.QuerySupplierAll()
}
func (sus *SupplierServer) GetSupplierAny(key string,data interface{}) (supplierData []model.Supplier,err error) {
	return sud.QuerySupplierByAny(key,data)
}
// 根据参数 查询 商家
func (sus *SupplierServer) GetSupplierByKeys(param *param.SupplierPatam) (supplierData []model.Supplier,err error) {
		count := param.Count
		start := param.Page * param.Count - (param.Count-1)
		var key string
		if param.Keys == "is_main" {
			key = "is_main"
			return sud.QuerySupplierByKey(key,count,start)
		}else if param.Condition == "rating" {
			return sud.QuerySupplierByKeyDesc(param.Condition,"desc",count,start)
		}else {
			key = param.Condition
			return sud.QuerySupplierByKey(key,count,start)
		}
}
// 根据ID 数组  查询商家
func (sus *SupplierServer) GetSupplierByIds(key string,data []int64) (supplierData []model.Supplier,err error) {
	return sud.QuerySupplierByIds(key,data)
}
// 根据名字 模糊查询
func (sus *SupplierServer) GetSupplierByName(key string,name []string) (supplierData []model.Supplier,err error) {
	for _,val := range name {
		res,resErr := sud.QuerySupplierByName(key,val)
		if resErr != nil {
			return
		}
		for _,v := range res {
			supplierData = append(supplierData, v)
		}
	}
	return
}
func (sus *SupplierServer) GetSupplierByCateId(cateId int64) (supplierData []model.Supplier,err error) {
	return sud.QuerySupplierByCateId(cateId)
}