package dal

import "orderserve/model"

type CatesInterface interface {
	InsertCateOne(cateInfo *model.Cates) (err error)
	DelCateById(id int64) (err error)
	UpdateCateById(id int64,cateInfo *model.Cates) (err error)
	QueryCateById(id int64) (cateData *model.Cates,err error)
	QueryCatesAll() (catesData []model.Cates,err error)
	QueryCatesByAny(key string,data interface{}) (catesData []model.Cates,err error)
}

type CatesDal struct {}

func NewCatesDal() CatesInterface {
	return &CatesDal{}
}

// 增加 类别
func (cd *CatesDal) InsertCateOne(cateInfo *model.Cates) (err error) {
	_,err = DB.InsertOne(cateInfo)
	return
}
// 删除 类别
func (cd *CatesDal) DelCateById(id int64) (err error) {
	cate := new(model.Cates)
	_,err = DB.ID(id).Delete(cate)
	return
}
// 更新 类别
func (cd *CatesDal) UpdateCateById(id int64,cateInfo *model.Cates) (err error) {
	_,err = DB.ID(id).Update(cateInfo)
	return
}
// 查询 类别
func (cd *CatesDal) QueryCateById(id int64) (cateData *model.Cates,err error) {
	_,err = DB.ID(id).Get(cateData)
	return
}
// 查询 所有类别
func (cd *CatesDal) QueryCatesAll() (catesData []model.Cates,err error) {
	err =DB.Find(&catesData)
	return
}
// 查询 条件类别
func (cd *CatesDal) QueryCatesByAny(key string,data interface{}) (catesData []model.Cates,err error) {
	err = DB.Where(key + " = ?",data).Find(&catesData)
	return
}
