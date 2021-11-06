package dal

import "orderserve/model"

type SupplierInterface interface {
	InsertSupplierOne(supplierInfo *model.Supplier) (err error)
	DelSupplierById(id int64) (err error)
	UpdateSupplierById(id int64,supplierInfo *model.Supplier) (err error)
	QuerySupplierById(id int64) (supplierData *model.Supplier,err error)
	QuerySupplierAll() (supplierData []model.Supplier,err error)
	QuerySupplierByAny(key string,data interface{}) (supplierData []model.Supplier,err error)
	QuerySupplierByKey(key string, count int, start int) (supplierData []model.Supplier,err error)
	QuerySupplierByIds(key string,data []int64) (supplierData []model.Supplier,err error)
	QuerySupplierByName(key string,name string) (supplierData []model.Supplier,err error)
	QuerySupplierByKeyDesc(key string,order string,count int, start int) (supplierData []model.Supplier,err error)
	QuerySupplierByCateId(cateId int64) (supplierData []model.Supplier,err error)
}

type SupplierDal struct {}

func NewSupplierDal() SupplierInterface {
	return &SupplierDal{}
}

// 添加 商家
func (sd *SupplierDal) InsertSupplierOne(supplierInfo *model.Supplier) (err error) {
	_,err = DB.InsertOne(supplierInfo)
	return
}

// 根据ID 删除商家
func (sd *SupplierDal) DelSupplierById(id int64) (err error) {
	supplier := new(model.Supplier)
	_,err = DB.ID(id).Delete(supplier)
	return
}
// 更新商家
func (sd *SupplierDal) UpdateSupplierById(id int64,supplierInfo *model.Supplier) (err error) {
	_,err = DB.ID(id).Update(supplierInfo)
	return
}
// 根据ID 查商家
func (sd *SupplierDal) QuerySupplierById(id int64) (supplierData *model.Supplier,err error) {
	var data model.Supplier
	_,err = DB.Where("id = ? and state = 1",id).Get(&data)
	supplierData = &data
	return
}
// 查询首页所有
func (sd *SupplierDal) QuerySupplierAll() (supplierData []model.Supplier,err error) {
	err = DB.Where("state = ? and is_main = ?",1,1).Find(&supplierData)
	return
}
// 查询 关键字 商家
func (sd *SupplierDal) QuerySupplierByAny(key string,data interface{}) (supplierData []model.Supplier,err error) {
	err = DB.In(key,data).Where("state = ?",1).Find(&supplierData)
	return
}
// 根据key 查询 limit 商家
func (sd *SupplierDal) QuerySupplierByKey(key string, count int, start int) (supplierData []model.Supplier,err error) {
	err = DB.Where(key + " = ? and state = ?",1,1).Limit(count,start).Find(&supplierData)
	return
}
// 根据商家类别 查询商家
func (sd *SupplierDal) QuerySupplierByCateId(cateId int64) (supplierData []model.Supplier,err error) {
	err = DB.Where("cate_id = ? and state = ?",cateId,1).Find(&supplierData)
	return
}
// 根据 ID 数组 查询商家
func (sd *SupplierDal) QuerySupplierByIds(key string,data []int64) (supplierData []model.Supplier,err error) {
	err = DB.In(key,data).Find(&supplierData)
	return
}
// 根据名字 模糊查询商家
func (sd *SupplierDal) QuerySupplierByName(key string,name string) (supplierData []model.Supplier,err error) {
	err = DB.Where(key + " like ? and state = 1","%"+name+"%").Find(&supplierData)
	return
}
//根据某字段 排序
func (sd *SupplierDal) QuerySupplierByKeyDesc(key string,order string,count int, start int) (supplierData []model.Supplier,err error) {
	err = DB.Where("state = 1").Limit(count,start).OrderBy(key + " " + order).Find(&supplierData)
	return
}
