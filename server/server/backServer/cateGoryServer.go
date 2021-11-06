package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model"
)

type ICategory interface {
	// 更新商家类别
	EditCateGoryById(id int64, categoryData model.Cates) (err error)
	// 添加商家类别
	AddCategoryOne(categoryData model.Cates) (err error)
}

type CategoryServer struct{}

func NewCategoryServer() ICategory {
	return &CategoryServer{}
}
var cgd = backendDal.NewCategoryDal()

// 更新商家类别
func (cgs *CategoryServer) EditCateGoryById(id int64, categoryData model.Cates) (err error) {
	return cgd.UpdateCateGoryById(id,categoryData)
}
// 添加商家类别
func (cgs *CategoryServer) AddCategoryOne(categoryData model.Cates) (err error) {
	return cgd.InsertCategoryOne(categoryData)
}