package backendDal

import (
	"orderserve/dal"
	"orderserve/model"
)

type CateGoryInterface interface {
	// 更新商家类别
	UpdateCateGoryById(id int64, categoryData model.Cates) (err error)
	// 添加商家类别
	InsertCategoryOne(categoryData model.Cates) (err error)
}

type CategoryDal struct {}

func NewCategoryDal() CateGoryInterface {
	return &CategoryDal{}
}

// 更新商家类别
func (cgd *CategoryDal) UpdateCateGoryById(id int64, categoryData model.Cates) (err error) {
	_,err = dal.DB.Where("id = ?",id).Update(&categoryData)
	return
}
// 添加商家类别
func (cgd *CategoryDal) InsertCategoryOne(categoryData model.Cates) (err error) {
	_,err = dal.DB.InsertOne(&categoryData)
	return
}