package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type ICates interface {
	AddCate(cateInfo *model.Cates) (err error)
	DelCate(id int64) (err error)
	UpdateCate(id int64,cateInfo *model.Cates) (err error)
	GetCate(id int64) (cateData *model.Cates,err error)
	GetCatesAll() (catesData []model.Cates,err error)
	GetCatesAny(key string,data interface{}) (catesData []model.Cates,err error)
}

var cd = dal.NewCatesDal()

type CatesServer struct {}

func NewCatesServer() ICates {
	return &CatesServer{}
}

func (cs *CatesServer) AddCate(cateInfo *model.Cates) (err error) {
	return cd.InsertCateOne(cateInfo)
}

func (cs *CatesServer) DelCate(id int64) (err error) {
	return cd.DelCateById(id)
}
func (cs *CatesServer) UpdateCate(id int64,cateInfo *model.Cates) (err error) {
	return cd.UpdateCateById(id,cateInfo)
}
func (cs *CatesServer) GetCate(id int64) (cateData *model.Cates,err error) {
	return cd.QueryCateById(id)
}
func (cs *CatesServer) GetCatesAll() (catesData []model.Cates,err error) {
	return cd.QueryCatesAll()
}
func (cs *CatesServer) GetCatesAny(key string,data interface{}) (catesData []model.Cates,err error) {
	return cd.QueryCatesByAny(key,data)
}
