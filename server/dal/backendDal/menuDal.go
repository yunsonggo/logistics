package backendDal

import (
	"orderserve/dal"
	"orderserve/model/backendModel"
)

type MenuDalInterface interface {
	// 查询所有菜单
	QueryMenuAll() (menusData []backendModel.BackendMenu,err error)
	// 添加菜单
	InsertMenuOne(menuData *backendModel.BackendMenu) (err error)
	// 更新菜单
	UpdateMenuOne(menuData *backendModel.BackendMenu) (err error)
}

type MenuDal struct {}

func NewMenuDal() MenuDalInterface {
	return &MenuDal{}
}
// 查询所有菜单
func (nd *MenuDal) QueryMenuAll() (menusData []backendModel.BackendMenu,err error) {
	err = dal.DB.Find(&menusData)
	return
}
// 添加菜单
func (nd *MenuDal) InsertMenuOne(menuData *backendModel.BackendMenu) (err error) {
	_,err = dal.DB.InsertOne(menuData)
	return
}
// 更新菜单
func (nd *MenuDal) UpdateMenuOne(menuData *backendModel.BackendMenu) (err error) {
	id := menuData.Id
	_,err = dal.DB.Id(id).Update(menuData)
	return
}